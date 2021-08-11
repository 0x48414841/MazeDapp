package game

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

const (
	UP    = iota //0
	RIGHT        //1
	DOWN         //2
	LEFT         //3
)

type Msg struct {
	Action string
	AllPos []Position
	Maze   [][]MazeData
}

type Position struct { //add more data here
	X, Y int
}

//var upgrader = websocket.Upgrader{} // use default options
var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }} //Prevents CORS error during local testing

func (G *Game) isJoinable() bool {
	return G.HasGameStarted == false && len(G.PlayersPosition) < 2
}

func (G *Game) handleIsJoinable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	if data, err := json.Marshal(IsJoinable{Answer: G.isJoinable(), Id: G.Id, Port: G.Port}); err == nil {
		w.Write(data)
	} else {
		log.Fatal(err)
	}
}

//defining a handler on a struct to access thread-specific data
func (G *Game) handleGameClient(w http.ResponseWriter, r *http.Request) {
	//check if lobby is joinable first. This should've already been done, so this is a redundant/sanity check
	if G.isJoinable() == false {
		return
	}

	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	//send maze to player
	G.sendMaze(c)
	G.initPlayer(c) //create and broadcast new player to all users

	for {
		msg := Msg{}
		if err := c.ReadJSON(&msg); err != nil {
			G.disconnectPlayer(c)
			log.Println("read:", err)
			return
		}

		//apply action to player's state in the maze
		if success := G.applyMove(c, msg.Action); success {
			G.broadcastMsg() // broadcast change to all users
		}
	}
}

func (G *Game) initPlayer(c *websocket.Conn) {
	G.mutex.Lock()
	G.PlayersPosition[c] = Position{X: 0, Y: 0}
	G.mutex.Unlock()

	G.broadcastMsg()
}

func (G *Game) disconnectPlayer(c *websocket.Conn) {
	//disconnecting a player entails giving the player
	//invalid coordinates --> broadcasting the update -->
	//then removing them from the map of players in that order

	G.mutex.Lock()
	//invalid coordinates prevents player from rendering in React app
	G.PlayersPosition[c] = Position{X: -1, Y: -1}
	G.mutex.Unlock()

	G.broadcastMsg()

	G.mutex.Lock()
	defer G.mutex.Unlock()
	delete(G.PlayersPosition, c)
}

func (G *Game) broadcastMsg() {
	G.mutex.Lock()
	defer G.mutex.Unlock()

	data := Msg{Action: "RECV_POS", AllPos: make([]Position, 0, len(G.PlayersPosition))} //capacity is the len of the map to prevent mem allocations within this crit section
	for _, pos := range G.PlayersPosition {
		data.AllPos = append(data.AllPos, pos)
	}

	for conn := range G.PlayersPosition {
		conn.WriteJSON(data)
	}
}

func (G *Game) sendMaze(c *websocket.Conn) {
	m := Msg{Action: "RECV_MAZE", Maze: G.Maze}
	if err := c.WriteJSON(m); err != nil {
		log.Println(err)
	}
}

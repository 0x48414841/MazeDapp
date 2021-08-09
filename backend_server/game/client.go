package game

import (
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
	Action int
	Pos    Position
}

type Position struct {
	X, Y int
}

var upgrader = websocket.Upgrader{} // use default options

//defining a handler on a struct to access thread-specific data
func (G *Game) handleGameClient(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

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
	for conn, pos := range G.PlayersPosition {
		conn.WriteJSON(Msg{Pos: pos})
	}
}

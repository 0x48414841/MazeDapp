package game

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const MAZE_ADDR = ":8080/maze"

var upgrader = websocket.Upgrader{} // use default options

type Game struct {
	Id             string //this is most likely redundant --> remove if necessary
	Maze           []byte
	Players        map[string]*net.Conn
	StartedAt      time.Time
	DataFromClient chan *Action
}

type Action struct {
}

//TODO use protobuf or grpc instead of structs
type MazeData struct {
	Row, Col                                   int
	IsNLinked, IsSLinked, IsELinked, IsWLinked bool
}
type API struct {
	Maze [][]MazeData
}

func createGame(addr, lobbyId string) {
	newGame := &Game{
		StartedAt: time.Now(),
	}

	//get Maze here

	r := mux.NewRouter()
	//path looks something like /game?id=LOBBYID
	r.Path("/game").Queries("id", fmt.Sprintf("{id:%s}", lobbyId)).
		HandlerFunc(newGame.handleGameClient) //TODO check this
	http.Handle("/", r)

	go runGame()
	log.Fatal(http.ListenAndServe(addr, r))
}

//defining a handler on a struct to access thread-specific data
func (G *Game) handleGameClient(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Fatal("read:", err)
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Fatal("write:", err)
		}
	}
}

func runGame() {
	//grab maze from other server
	conn, err := net.Dial("tcp", MAZE_ADDR) //net.http instead
	if err != nil {
		log.Fatal(err)
		return
	}

	//conn.Read()
	_ = conn

}

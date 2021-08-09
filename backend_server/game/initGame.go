package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

const MAZE_ADDR = "http://localhost:8080/maze"

type Game struct {
	Maze            [][]MazeData
	PlayersPosition map[*websocket.Conn]Position
	StartedAt       time.Time
	mutex           *sync.Mutex
}

//TODO use protobuf or grpc instead of POD with JSON encoding
type MazeData struct {
	Row, Col                                   int
	IsNLinked, IsSLinked, IsELinked, IsWLinked bool
}
type API struct {
	Maze [][]MazeData
}

//creates a new game and returns the new id and address of that game
//This function is exposed for other packages and is a wrapper for
//createGame()
func CreateGame() (string, string) {
	id := generateLobbyId()

	newAddress := ":" + randomPort()
	for newAddress == "" {
		newAddress = ":" + randomPort()
	}
	go createGame(newAddress, id)
	return id, newAddress
}

//Inits the Game struct and starts server for clients
func createGame(addr, lobbyId string) {
	newGame := Game{
		StartedAt:       time.Now(),
		Maze:            getMaze(),
		PlayersPosition: make(map[*websocket.Conn]Position),
		mutex:           &sync.Mutex{},
	}

	r := mux.NewRouter()

	//path looks something like /game?id=LOBBYID
	r.Path("/game").Queries("id", fmt.Sprintf("{id:%s}", lobbyId)).
		HandlerFunc(newGame.handleGameClient)

	log.Printf("starting server at http://localhost%s/game?id=%s", addr, lobbyId)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		log.Fatal(err)
	}
}

//grab maze from other server
func getMaze() [][]MazeData {
	resp, err := http.Get(MAZE_ADDR)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	maze := API{}
	if err := json.Unmarshal(body, &maze); err != nil {
		log.Fatal(err)
	}
	return maze.Maze
}

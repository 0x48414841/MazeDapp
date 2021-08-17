package game

import (
	"backend/lobbies"
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
	Id, Port        string
	Maze            [][]MazeData
	PlayersPosition map[*websocket.Conn]*Player
	StartedAt       time.Time
	HasGameStarted  bool
	mutex           *sync.Mutex
	eventChan       chan bool
	verifiedWages   int
}

type Player struct {
	Username string
	Pos      Position
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
		Id:              lobbyId,
		Port:            addr,
		StartedAt:       time.Now(),
		Maze:            getMaze(),
		PlayersPosition: make(map[*websocket.Conn]*Player),
		HasGameStarted:  false,
		mutex:           &sync.Mutex{},
		eventChan:       make(chan bool),
		verifiedWages:   0,
	}
	if len(newGame.Maze) == 0 {
		return
	}
	r := mux.NewRouter()

	//path looks something like /game?id=LOBBYID
	r.Path("/game").Queries("id", fmt.Sprintf("{id:%s}", lobbyId)).
		HandlerFunc(newGame.handleGameClient)
	r.HandleFunc("/isJoinable", newGame.handleIsJoinable)

	//REPLACE WITH lobbies PACKAGE
	//ActiveLobbies = append(ActiveLobbies,
	//	Lobby{LobbyId: lobbyId, Port: addr, EventListener: make(chan bool)},
	//)
	eventListenerChan := make(chan bool)
	lobbies.AddLobby(lobbyId, &lobbies.Lobby{Port: addr, EventListener: eventListenerChan})

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
		log.Println(err)
		return [][]MazeData{}
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	maze := API{}
	if err := json.Unmarshal(body, &maze); err != nil {
		log.Fatal(err)
	}
	return maze.Maze
}

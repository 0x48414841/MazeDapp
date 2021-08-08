package game

import (
	"net"
	"sync"
	"time"
)

//TODO consider pros of csp-style approach

var ActiveLobbies map[string]string

var mutex *sync.Mutex

type Lobby struct {
	Id        string //this is most likely redundant --> remove if necessary
	Maze      []byte
	Players   map[string]*net.Conn
	StartedAt time.Time
}

//returns newly created lobby id
func CreateLobby() (string, string) {
	id := RandStringBytes()
	mutex.Lock()
	defer mutex.Unlock()

	var newAddress string = ""
	for newAddress != "" {
		newAddress = ":" + RandomPort()
	}

	ActiveLobbies[id] = newAddress

	/*ActiveLobbies[id] = &Lobby{
		Id:      id,
		Players: map[string]*net.Conn{},
	} */
	go createGame(newAddress, id)
	return id, newAddress
}

func DeleteLobby(id string) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := ActiveLobbies[id]; ok {
		delete(ActiveLobbies, id)
	}
}

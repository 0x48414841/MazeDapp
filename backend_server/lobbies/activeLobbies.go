package lobbies

import "sync"

var activeLobbies map[string]*Lobby
var mutex sync.Mutex

type Lobby struct {
	Port string
	//The chan will be used to synchronize the events module with a
	//server hosting a game.
	//NOTE: when refactoring to headless servers, each server will have their own chan
	//listening for events
	EventListener chan bool
}

func AddLobby(lobbyId string, lobby *Lobby) {
	mutex.Lock()
	defer mutex.Unlock()
	activeLobbies[lobbyId] = lobby
}

func RemoveLobby(lobbyId string) {
	mutex.Lock()
	defer mutex.Unlock()
	if _, ok := activeLobbies[lobbyId]; ok {
		delete(activeLobbies, lobbyId)
	}
}

func GetLobbyChan(lobbyId string) chan bool {
	mutex.Lock()
	defer mutex.Unlock()
	if lobby, ok := activeLobbies[lobbyId]; ok {
		return lobby.EventListener
	}
	return nil
}

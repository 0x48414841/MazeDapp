package main

import (
	"backend/game"
	"backend/lobbies"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type GameAddr struct {
	Id, Addr string
}

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	lobbyId, gameAddr := game.CreateGame()
	if data, err := json.Marshal(GameAddr{Id: lobbyId, Addr: gameAddr}); err == nil {
		w.Write(data)
	}
}

func JoinGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK) //maybe add e
	//lobbyId, gameAddr := game.JoinGame()
	lobbyId, gameAddr := lobbies.FindJoinableLobby()
	if data, err := json.Marshal(GameAddr{Id: lobbyId, Addr: gameAddr}); err == nil {
		w.Write(data)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createGame", CreateGameHandler)
	r.HandleFunc("/joinGame", JoinGameHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}

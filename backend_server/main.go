package main

import (
	"backend/game"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type NewGame struct {
	Id, Addr string
}

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	lobbyId, gameAddr := game.CreateGame()
	if data, err := json.Marshal(NewGame{Id: lobbyId, Addr: gameAddr}); err == nil {
		w.Write(data)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createGame", CreateGameHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}

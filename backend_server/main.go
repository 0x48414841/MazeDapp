package main

import (
	"backend/game"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var ActiveLobbies map[string]bool

func CreateGameHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	lobbyId, gameAddr := game.CreateLobby()
	fmt.Fprintf(w, lobbyId)
	_ = gameAddr
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/createGame", CreateGameHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8000", r)
}

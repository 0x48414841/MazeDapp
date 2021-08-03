package main

import (
	"log"
	"net/http"
	"server/maze"
)

func main() {
	//for now, only generate 10x10 mazes
	//TODO add quey strings for customizable mazes
	mazeHandler := func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		maze := maze.MakeBinaryTree(10, 10)
		maze.On()
		mazeJSON := maze.Grid.GridToJSON()
		w.Write(mazeJSON)
	}

	http.HandleFunc("/maze", mazeHandler)
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}

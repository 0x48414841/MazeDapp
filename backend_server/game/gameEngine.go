package game

import (
	"github.com/gorilla/websocket"
)

func (G *Game) applyMove(player *websocket.Conn, action string) bool {
	G.mutex.Lock()
	defer G.mutex.Unlock()

	currentPos := G.PlayersPosition[player].Pos

	//applyAction
	newPosition := currentPos

	if action == "UP" && currentPos.X-1 >= 0 &&
		G.Maze[currentPos.X][currentPos.Y].IsNLinked {
		newPosition.X -= 1
	} else if action == "DOWN" && currentPos.X+1 < len(G.Maze) &&
		G.Maze[currentPos.X][currentPos.Y].IsSLinked {
		newPosition.X += 1
	} else if action == "LEFT" && currentPos.Y-1 >= 0 &&
		G.Maze[currentPos.X][currentPos.Y].IsWLinked {
		newPosition.Y -= 1
	} else if action == "RIGHT" && currentPos.Y+1 < len(G.Maze[0]) &&
		G.Maze[currentPos.X][currentPos.Y].IsELinked {
		newPosition.Y += 1
	}

	if currentPos.X == newPosition.X && currentPos.Y == newPosition.Y {
		return false //newPos == oldPos --> no change
	} else {
		//commit change
		G.PlayersPosition[player].Pos = newPosition
		return true
	}
}

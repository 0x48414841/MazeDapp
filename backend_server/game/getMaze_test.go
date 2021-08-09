package game

import (
	"fmt"
	"testing"
)

//this requires the maze at MAZE_ADDR to be running
func TestGetMaze(t *testing.T) {
	maze := getMaze()
	fmt.Println(maze[0])
	fmt.Println(maze[1][0])
	fmt.Println(maze[1][0].IsNLinked, maze[1][0].IsELinked)
	t.Error() // force error
}

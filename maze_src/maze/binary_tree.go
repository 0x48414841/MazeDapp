package maze

import (
	"math/rand"
	"time"
)

type BinaryTree struct {
	Grid *Grid
}

func MakeBinaryTree(row, col int) *BinaryTree {
	return &BinaryTree{Grid: MakeGrid(row, col)}
}

func (this *BinaryTree) On() {
	rand.Seed(time.Now().UnixNano())
	nextCell := this.Grid.EachCell()
	for cell := nextCell(); cell != nil; cell = nextCell() {
		neigbors := []*Cell{}

		if cell.Neighbors[NORTH] != nil {
			neigbors = append(neigbors, cell.Neighbors[NORTH])
		}
		if cell.Neighbors[EAST] != nil {
			neigbors = append(neigbors, cell.Neighbors[EAST])
		}
		if len(neigbors) == 0 {
			continue
		}

		index := rand.Intn(len(neigbors))
		neighbor := neigbors[index]
		cell.Link(neighbor, true)
	}
}

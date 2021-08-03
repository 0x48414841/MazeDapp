package maze

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type Grid struct {
	Row, Col int       //number of rows and cols in maze
	Grid     [][]*Cell //container for all cells in maze
}

func MakeGrid(width, height int) *Grid {
	grid := &Grid{Row: height, Col: width, Grid: nil}
	grid.PrepareGrid()
	grid.ConfigureCells()
	return grid
}

func (this *Grid) PrepareGrid() {
	this.Grid = make([][]*Cell, this.Row, this.Row)
	for i := range this.Grid {
		this.Grid[i] = make([]*Cell, this.Col, this.Col)
		for j := range this.Grid[i] {
			this.Grid[i][j] = MakeCell(i, j)
		}
	}
}

func (this *Grid) ConfigureCells() {
	for row := range this.Grid {
		for col := range this.Grid[row] {
			cell := this.Grid[row][col]
			cell.Neighbors[NORTH] = this.GetCell(row-1, col)
			cell.Neighbors[SOUTH] = this.GetCell(row+1, col)
			cell.Neighbors[WEST] = this.GetCell(row, col-1)
			cell.Neighbors[EAST] = this.GetCell(row, col+1)
		}
	}
}

func (this *Grid) GetRow(row int) []*Cell {
	if row < 0 || row >= len(this.Grid) {
		return nil
	}
	return this.Grid[row]
}

func (this *Grid) GetCell(row, col int) *Cell { //provides bounds checking on Grid
	if row < 0 || row >= len(this.Grid) || col < 0 || col >= len(this.Grid[row]) {
		return nil
	}
	return this.Grid[row][col]
}

func (this *Grid) GetRandomCell() *Cell {
	rand.Seed(time.Now().UnixNano())
	row := rand.Intn(this.Row)
	col := rand.Intn(this.Col)
	return this.GetCell(row, col)
}

func (this *Grid) GetSize() int {
	return this.Row * this.Col
}

//An iterator using a closure!
func (this *Grid) EachRow() func() []*Cell {
	i := -1
	return func() []*Cell {
		i++
		return this.GetRow(i)
	}
}

//Another iterator using a closure!
func (this *Grid) EachCell() func() *Cell {
	//row := this.EachRow()() //the extra '()' automatically invokes iterator and return first row
	nextRow := this.EachRow()
	for row := nextRow(); row != nil; row = nextRow() {
		i, j := 0, -1
		return func() *Cell {
			j++
			if j == len(this.Grid) {
				i, j = i+1, 0
			}
			return this.GetCell(i, j)
		}
	}

	//need a final return value to make Go stop bullying me
	return func() *Cell {
		return nil
	}
}

func (this *Grid) DrawGrid() string {
	output := "+" + strings.Repeat("---+", this.Col) + "\n"

	for _, row := range this.Grid {
		top, bottom := "|", "+"
		for _, cell := range row {
			if cell == nil {
				cell = MakeCell(-1, -1)
			}
			body := "   " // Three (3) spaces
			var eastBoundary, southBoundary string
			if cell.IsLinked(cell.Neighbors[EAST]) {
				eastBoundary = " "
			} else {
				eastBoundary = "|"
			}
			top += body + eastBoundary

			if cell.IsLinked(cell.Neighbors[SOUTH]) {
				southBoundary = "   "
			} else {
				southBoundary = "---"
			}
			corner := "+"
			bottom += southBoundary + corner
		}
		output += top + "\n"
		output += bottom + "\n"
	}
	return output
}

func (this *Grid) GridToJSON() []byte {
	type APIData struct {
		Row, Col                                   int
		IsNLinked, IsSLinked, IsELinked, IsWLinked bool
	}
	type API struct {
		Maze [][]APIData
	}

	response := API{Maze: make([][]APIData, len(this.Grid))}

	for i, row := range this.Grid {
		response.Maze[i] = make([]APIData, len(row))
		for j, cell := range row {
			response.Maze[i][j] = APIData{
				Row: i, Col: j,
				//if cell[dir] is linked, then there is no border in that direction
				IsNLinked: cell.IsLinked(cell.Neighbors[NORTH]),
				IsSLinked: cell.IsLinked(cell.Neighbors[SOUTH]),
				IsELinked: cell.IsLinked(cell.Neighbors[EAST]),
				IsWLinked: cell.IsLinked(cell.Neighbors[WEST])}
		}
	}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return []byte{}
	}
	return b
}

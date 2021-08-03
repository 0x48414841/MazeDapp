package maze

const (
	NORTH = iota //0
	SOUTH        //1
	EAST         //2
	WEST         //3
)

const NEIGHBORS int = 4

type Cell struct {
	Row, Col  int              //where this cell lives in the grid
	Neighbors [NEIGHBORS]*Cell //points to immediate four cells
	Links     map[*Cell]bool   // keeps track of which neighboring cells are
	//linked (joined by a passage) to this cell
}

func MakeCell(row, col int) *Cell {
	return &Cell{
		Row:       row,
		Col:       col,
		Neighbors: [NEIGHBORS]*Cell{},
		Links:     make(map[*Cell]bool),
	}
}

func (this *Cell) Link(c *Cell, linkBoth bool) {
	this.Links[c] = true
	if linkBoth {
		c.Links[this] = true
	}
}

func (this *Cell) Unlink(c *Cell, unlinkBoth bool) {
	if _, ok := this.Links[c]; ok {
		delete(this.Links, c)
	}

	if unlinkBoth {
		if _, ok := c.Links[this]; ok {
			delete(c.Links, this)
		}
	}
}

func (this *Cell) GetLinks() []*Cell {
	keys := make([]*Cell, 0, len(this.Links))
	for k := range this.Links {
		keys = append(keys, k)
	}
	return keys
}

func (this *Cell) IsLinked(c *Cell) bool {
	if _, ok := this.Links[c]; ok {
		return true
	}
	return false
}

func (this *Cell) GetNeighbors() []*Cell {
	neighbors := make([]*Cell, 0)
	for _, n := range this.Neighbors {
		if n != nil {
			neighbors = append(neighbors, n)
		}
	}
	return this.Neighbors[:]
}

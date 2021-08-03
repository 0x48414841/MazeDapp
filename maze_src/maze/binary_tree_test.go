package maze

import (
	"fmt"
	"testing"
)

func TestBT(t *testing.T) {
	BT := MakeBinaryTree(10, 10)
	BT.On()
	fmt.Println(BT.Grid.DrawGrid())
	a := BT.Grid.GridToJSON()
	fmt.Println(a)
	fmt.Println(string(a))
	t.Error() //force Go to show print statements
}

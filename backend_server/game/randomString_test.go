package game

import "testing"

func TestEntropy(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 5e6; /*5 (five) million*/ i++ {
		id := RandStringBytes()
		if _, ok := m[id]; ok == true {
			t.Fatal("Id Colision")
		} else {
			m[id] = true
		}
	}
}

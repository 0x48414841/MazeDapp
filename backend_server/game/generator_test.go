package game

import (
	"fmt"
	"testing"
)

func TestEntropy(t *testing.T) {
	m := make(map[string]bool)
	for i := 0; i < 5e6; /*5 (five) million*/ i++ {
		id := generateLobbyId()
		if _, ok := m[id]; ok == true {
			t.Fatal("Id Colision")
		} else {
			m[id] = true
		}
	}
}

func TestRandomPort(t *testing.T) {
	m := make(map[string]int)
	for i := 0; i < 1000; i++ {
		id := randomPort()
		fmt.Println("id = ", id)
		if id == "" {
			t.Fatal("")
		}
		m[id] += 1
		if m[id] > 10 {
			t.Error("not enough entropy")
		}
	}
}

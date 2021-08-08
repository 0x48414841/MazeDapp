package game

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789"
const LOBBY_ID_SIZE = 25

func RandStringBytes() string {
	b := make([]byte, LOBBY_ID_SIZE)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func RandomPort() string {
	defer func() { //capture panic in case rand.Intn returns 0
		if r := recover(); r != nil {
			log.Println("Recovered!!", r)
		}
	}()
	rand.Seed(time.Now().UTC().UnixNano())
	return strconv.Itoa(9000 + rand.Intn(1000))
}

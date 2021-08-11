package game

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var ActiveLobbies []string

type IsJoinable struct { //this might be superfluous
	Answer   bool
	Port, Id string
}

//loop through all known servers and query each one to determine if server is joinable
//**Why use extra resources when I can simply access the game servers in this process?
//Well, eventually I want to make all of the new servers a separate PROCESS to mimic headless servers
//and making the extra requests now will mean less refactoring later. Cheerio
func JoinGame() (string, string) {
	for _, server := range ActiveLobbies {

		resp, err := http.Get(fmt.Sprintf("http://localhost%s/isJoinable", server))
		if err != nil {
			log.Println(err)
			continue
		}
		body, err := ioutil.ReadAll(resp.Body)

		answer := IsJoinable{}
		if err := json.Unmarshal(body, &answer); err != nil {
			log.Fatal(err)
		}
		resp.Body.Close()
		if answer.Answer {
			return answer.Id, answer.Port
		}
	}
	//this should never execute
	log.Println("NO LOBBY FOUND")
	return "", ""
}

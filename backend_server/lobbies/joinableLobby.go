package lobbies

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IsJoinable struct { //this might be superfluous
	Answer   bool
	Port, Id string
}

//loop through all known servers and query each one to determine if server is joinable
//**Why use extra resources when I can simply access the game servers in this process?
//Well, eventually I want to make all of the new servers a separate PROCESS to mimic headless servers
//and making the extra requests now will mean less refactoring later. Cheerio
func FindJoinableLobby() (string, string) {
	mutex.Lock()
	defer mutex.Unlock()

	for server := range activeLobbies {
		log.Println("sending request to ", fmt.Sprintf("http://localhost%s/isJoinable", server))
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
	//TODO add more logic here
	log.Println("NO LOBBY FOUND")
	return "", ""
}

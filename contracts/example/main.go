package main

import (
	operations "contract/example/operations"
	"fmt"
	"os"
)

func main() {
	//get the argument from user-input
	var arg string = ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	arg = "event"
	switch arg {
	case "deploy":
		operations.RunDeploy()
	case "load":
		operations.RunLoad()
	case "query":
		operations.RunQuery()
	case "write":
		operations.RunWrite()
	case "event":
		operations.RunEvent()
	default:
		fmt.Println("No matching arguments")
	}
}

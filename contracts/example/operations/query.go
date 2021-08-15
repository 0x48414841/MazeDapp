package operations

import (
	"contract/example/sc"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RunQuery() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/642bc71ec51d412290db00d03bb86687")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x893C5Ea4575E5929231E7168F1D72eB072CE39bE") //smart contract address returned after deploying it
	instance, err := sc.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")

	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}

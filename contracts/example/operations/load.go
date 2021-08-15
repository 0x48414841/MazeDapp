package operations

import (
	"fmt"
	"log"

	"contract/example/sc"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RunLoad() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/642bc71ec51d412290db00d03bb86687")
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress("0x893C5Ea4575E5929231E7168F1D72eB072CE39bE") //smart contract address returned after deploying it
	instance, err := sc.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}
	_ = instance

	fmt.Println("contract is loaded")
}

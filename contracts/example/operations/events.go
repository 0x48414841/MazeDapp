package operations

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RunEvent() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/642bc71ec51d412290db00d03bb86687")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("0x893C5Ea4575E5929231E7168F1D72eB072CE39bE")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	//execute SetItem function on the Smart Contract to emit an event that this thread will catch
	go RunWrite()

	select {
	case err := <-sub.Err():
		log.Fatal(err)
	case vLog := <-logs:
		//fmt.Println(vLog) // pointer to event log
		//Data = data emitted
		//TxHash is the TxHash that called the function to emit data
		//Contract Address is ... the address of the contract

		//vLog.Data == [102 111 111 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 98 97 114 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
		// --> two 32 byte-size elements
		fmt.Fprintf(os.Stdout, "Data = %s\nTxHash = %s\nContract Address = %s\n\n\n", vLog.Data, vLog.TxHash, vLog.Address)
	}

	//experiment: If multiple threads are listening for the same event, will they all hear it?
	//A: No :v
}

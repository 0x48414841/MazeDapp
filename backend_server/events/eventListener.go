package events

import (
	"backend/lobbies"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func eventListener() {
	client, err := ethclient.Dial("")
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress("")
	query := ethereum.FilterQuery{
		//listen to events created by the supplied contract(s) in slice below
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Fprintf(os.Stdout, "Data = %s\nTxHash = %s\nContract Address = %s\n\n\n", vLog.Data, vLog.TxHash, vLog.Address)
			serverId := ""
			alertServer(serverId)
		}
	}

}

func alertServer(serverId string) {
	gameServerChan := lobbies.GetLobbyChan(serverId)
	if gameServerChan != nil {
		gameServerChan <- true
	}
}

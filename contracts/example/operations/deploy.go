package operations

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"contract/example/sc"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func RunDeploy() {
	client, err := ethclient.Dial("https://ropsten.infura.io/v3/642bc71ec51d412290db00d03bb86687")
	if err != nil {
		log.Fatal(err)
	}

	//private key is from Ganache
	privateKey, err := crypto.HexToECDSA("REPLACE WITH YOUR PRIVATE KEY")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		panic(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		panic(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := sc.DeployStore(auth, client, input)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // SC address --> 0x893C5Ea4575E5929231E7168F1D72eB072CE39bE
	fmt.Println(tx.Hash().Hex()) // Transaction Hash --> 0x3ec9ba8f8a61e1a784cb678320a28eb781b4d374bbb27024a7f6fc45d4085523

	_ = instance
}

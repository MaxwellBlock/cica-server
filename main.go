package main

import (
	"context"
	"fmt"
	"server/contract"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Connector struct {
	ctx      context.Context
	conn     *ethclient.Client
	contract *contract.Hello
}

func NewConnector() *Connector {
	conn, err := ethclient.Dial("")
	if err != nil {
		panic(err)
	}

	helloContract, err := contract.NewHello(common.HexToAddress(""), conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	return &Connector{
		ctx:      context.Background(),
		conn:     conn,
		contract: helloContract,
	}
}

func main() {
	fmt.Println("Simulate the starting")
}

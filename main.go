package main

import (
	"cica-server/conf"
	"cica-server/contract/hello"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Connector struct {
	ctx           context.Context
	conn          *ethclient.Client
	helloContract *hello.Hello
}

func NewConnector() *Connector {
	conn, err := ethclient.Dial(conf.Config.Local.Provider)
	if err != nil {
		panic(err)
	}

	helloContract, err := hello.NewHello(common.HexToAddress(conf.Config.Remote.Addr.Hello), conn)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected!")

	return &Connector{
		ctx:           context.Background(),
		conn:          conn,
		helloContract: helloContract,
	}
}

func main() {
	fmt.Println("Simulate the starting")

	conf.Init()
	conn := NewConnector()

	go func() {
		for {
			counter, err := conn.helloContract.Counter(nil)
			if err != nil {
				panic(err)
			}
			fmt.Println(counter.String())
			time.Sleep(time.Second * 10)
		}
	}()

	gin.SetMode(gin.DebugMode)

	r := gin.Default()

	// 测试连通性
	r.GET("ping", Ping)

	r.Run(":2022")
}

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"pong": "ok!",
	})
}

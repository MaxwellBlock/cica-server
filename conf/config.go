package conf

import (
	"cica-server/utils"
	"encoding/json"
	"fmt"
	"os"
)

var Config config

type config struct {
	Local  localConfig
	Remote remoteConfig
}

type localConfig struct {
	Provider string `json:"provider"`
}

type remoteConfig struct {
	Addr  addr
	Event event
}

type addr struct {
	Hello string
}

type event struct {
	Hello map[string]string
}

func Init() {

	data, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Config.Local)
	if err != nil {
		panic(err)
	}

	Config.Remote = fetchRemoteConfig()

	fmt.Printf("%+v\n\n", Config)
}

func fetchRemoteConfig() (res remoteConfig) {

	res.Addr.Hello = utils.FetchContractAddress(
		"cica-contract", "develop", "latest", "bsctest", "Hello")

	res.Event.Hello = utils.FetchContractEventHash("cica-contract", "develop", "latest", "bsctest", "Hello")

	return
}

package conf

import (
	"encoding/json"
	"os"
)

var Config config

type config struct {
	Provider  string `json:"provider"`
	HelloAddr string `json:"hello_addr"`
}

func Init() {

	data, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(data, &Config)
	if err != nil {
		panic(err)
	}
}

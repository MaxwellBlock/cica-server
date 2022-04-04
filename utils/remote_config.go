package utils

import (
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"strings"
)

func FetchContractAddress(project, env, version, network, contract string) string {

	url := fmt.Sprintf("http://47.241.117.7:2022/contracts/get?project=%s&env=%s&"+
		"version=%s&key=address&network=%s&contract=%s", project, env, version, network, contract)

	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	return strings.ReplaceAll(string(data), "\"", "")
}

func FetchContractEventHash(project, env, version, network, contract string) map[string]string {

	res := make(map[string]string)

	url := fmt.Sprintf("http://47.241.117.7:2022/contracts/get?project=%s&env=%s&"+
		"version=%s&key=event&network=%s&contract=%s", project, env, version, network, contract)

	rsp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	data, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		panic(err)
	}

	json := gjson.ParseBytes(data)

	for k, v := range json.Map() {
		res[v.Get("name").String()] = k
	}

	return res
}

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type balancerConfig struct {
	NetworkInterface string     `json:"interface"`
	Upstreams        []upstream `json:"upstreams`
}

type upstream struct {
	HttpPath    string       `json:"path"`
	HttpMethods []HttpMethod `json:"methods"`
	Backends    []Backend    `json:"backends"`
	ProxyMethod string       `json:"proxyMethod"`
}

type Backend string

type HttpMethod string

var (
	configFilePath string
	config         balancerConfig
)

func main() {
	configFilePath = "config.json"
	configFile, err := os.Open(configFilePath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer configFile.Close()

	b, err := ioutil.ReadAll(configFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	json.Unmarshal([]byte(b), &config)
	fmt.Println(config.Upstreams[0].Backends[0])
}

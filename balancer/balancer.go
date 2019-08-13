package main

import (
	// "bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func runBalancer() error {
	r := mux.NewRouter()
	r.HandleFunc(config.HttpPaths, balancerHandler)
	// r.HandleFunc(config.HttpPath, handler)

	log.Fatal(http.ListenAndServe(config.NetworkInterface, r))
	return nil
}

func balancerHandler(w http.ResponseWriter, req *http.Request) {

	var randomBackend int
	if config.ProxyMethod == "round-robin" {
		randomBackend = rand.Intn(len(config.Backends))

	}
	url := config.Backends[randomBackend]

	fmt.Println(url)

	proxyReq, err := http.NewRequest(req.Method, url, req.Body)

	client := http.Client{}

	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	w.Write([]byte(respBody))
}

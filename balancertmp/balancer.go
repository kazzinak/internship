package main

import (
	"bytes"
	// "fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
)

func runBalancer() error {
	r := mux.NewRouter()
	r.HandleFunc(config.HttpPath, balancerHandler)
	// r.HandleFunc(config.HttpPath, handler)

	log.Fatal(http.ListenAndServe(config.NetworkInterface, r))
	return nil
}

func balancerHandler(w http.ResponseWriter, req *http.Request) {

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	var randomBackend int
	if config.ProxyMethod == "round-robin" {
		randomBackend = rand.Intn(len(config.Backends))

	}
	url := config.Backends[randomBackend]

	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(body))
	proxyReq.Header = make(http.Header)
	for h, val := range req.Header {
		proxyReq.Header[h] = val
	}

	client := http.Client{}
	resp, err := client.Do(proxyReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// legacy code
	respBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write([]byte(respBody))

}

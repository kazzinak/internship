package main

import (
	"bytes"
	"fmt"
	// "github.com/gorilla/mux"
	"github.com/urfave/cli"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, req *http.Request) {
	// we need to buffer the body if we want to read it here and send it
	// in the request.
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// you can reassign the body if you need to parse it as multipart
	req.Body = ioutil.NopCloser(bytes.NewReader(body))

	// create a new url from the raw RequestURI sent by the client
	url := fmt.Sprintf("%s://%s%s", "http", "ya.rur", req.RequestURI)

	proxyReq, err := http.NewRequest(req.Method, url, bytes.NewReader(body))

	// We may want to filter some headers, otherwise we could just use a shallow copy
	// proxyReq.Header = req.Header
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
	fmt.Sprintf("%s", w)
}

func main() {
	// r := mux.NewRouter()
	// // r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// // 	vars := mux.Vars(r)
	// // 	title := vars["title"]
	// // 	page := vars["page"]

	// // 	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	// // })
	// r.HandleFunc("/", handler)

	// log.Fatal(http.ListenAndServe(":8080", r))
	app := cli.NewApp()
	app.Name = "balancer"
	app.Usage = "reverse proxy"
	app.Action = func(c *cli.Context) error {
		fmt.Println("boom! I say!")
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

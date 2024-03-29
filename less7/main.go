package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type FibNumber struct {
	Current int
	Prev    int
	Next    int
}

func handler(w http.ResponseWriter, r *http.Request) {
	s := strings.Split(r.URL.Path, "/")

	number, err := strconv.Atoi(s[len(s)-1])
	if err != nil {
		panic(err)
	}

	res := &FibNumber{}
	if number > 0 {
		res = &FibNumber{Current: fibonacci(number), Prev: fibonacci(number - 1), Next: fibonacci(number + 1)}
	} else {

		res = &FibNumber{Current: 0, Prev: 0, Next: 0}
	}

	data, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}

	w.Write([]byte(data))
}

func main() {
	http.HandleFunc("/fibonacci/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fibonacci(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

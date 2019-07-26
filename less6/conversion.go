package main

import (
	"fmt"
)

type A struct {
	C string
}

type B struct {
	C string
	K string
}

func main() {
	a := A{"idris"}
	// b := A(a)
	b := B{a, "asd"}
	fmt.Println(a)
}

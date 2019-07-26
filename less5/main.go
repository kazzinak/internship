package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	f, err := os.OpenFile("output.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n := 10
	ch := make(chan int)
	// stdoutCh := make(chan int)
	// fileCh := make(chan int)

	go fibonacci(n, ch)

	for i := 0; i < n; i++ {
		wg.Add(1)
		x := <-ch
		go printStdout(x, wg)
		//go func() {
		// 	// wg.Add(1)
		// 	// stdoutCh <- x
		// }()
		// go func() {
		// 	wg.Add(1)
		// 	fileCh <- x
		// }()
	}
	// go writeToFile(n, fileCh, f, wg)
	// printStdout(n, stdoutCh, wg)
	wg.Wait()
}

func fibonacci(n int, ch chan int) {
	res := 0
	x := 0
	y := 1
	for i := 0; i <= n; i++ {
		ch <- x
		res = x + y
		x = y
		y = res
	}
}

func writeToFile(fileCh <-chan int, f *os.File, wg sync.WaitGroup) {
	f.WriteString(string(<-fileCh))
	defer wg.Done()
}

func printStdout(x int, wg sync.WaitGroup) {
	// for i := 0; i < n; i++ {
	// 	fmt.Println(<-stdoutCh)
	// 	// defer wg.Done()
	// }
	// stdoutCh := make(chan int)
	// stdoutCh <- x
	// close(stdoutCh)
	// fmt.Println(<-stdoutCh)
	fmt.Println(x)
	defer wg.Done()
}

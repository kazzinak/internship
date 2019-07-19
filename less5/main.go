package main

import (
	"fmt"
	"sync"
	"time"
)

var a []int = []int{}

func main() {
	// fmt.Println("hello")
	// ch := make(chan int)
	// ch2 := make(chan int)
	// select {
	// case b := <-ch2:
	// 	fmt.Println("read from", b)
	// case x := <-ch:
	// 	fmt.Println("read from", x)
	// default:
	// 	fmt.Println("default")
	// }
	var mu sync.Mutex
	var muOnce sync.Once
	var wg sync.WaitGroup

	wg.Add(1)
	wg.Done()

	for i := 0; i < 10; i++ {
		go func(k int) {
			defer wg.Done()
			mu.Lock()
			a = append(a, k)
			mu.Unlock()
		}(i)
		muOnce.Do(func() {
			fmt.Println("only once")
		})
	}
	time.Sleep(1 * time.Second)
	fmt.Println(a)
}

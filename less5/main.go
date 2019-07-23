package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	f, err := os.OpenFile("output.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	n := 10

	ch := make(chan int)
	stdoutCh := make(chan int)
	fCh := make(chan int)

	go fibonacci(n, ch)
	// for i := range ch {

	// 	fmt.Println(i)
	// }
	for i := 0; i < n; i++ {
		wg.Add(1)

		x := <-ch
		fmt.Println(x)
		stdoutCh <- x
		fCh <- x
	}
	wg.Wait()
	// for i := range stdoutCh {
	// 	// fmt.Println(<-stdoutCh)
	// 	fmt.Println(i)
	// 	wg.Done()
	// }
	// for res := range fCh {
	// 	f.WriteString(string(res))
	// }
	fmt.Println("hello")
	wg.Wait()

}

func fibonacci(n int, ch chan int) {
	res := 0
	x := 0
	y := 1
	// if n == 0 || n == 1 {
	// 	ch <- n
	// }
	for i := 0; i <= n; i++ {
		ch <- x
		res = x + y
		x = y
		y = res
	}
	close(ch)
}

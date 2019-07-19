package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type letters map[rune]int

func main() {

	var lines []string

	data, err := os.Open("notes.txt")
	check(err)

	file := bufio.NewReader(data)

	for {
		line, _, err := file.ReadLine()
		if err == io.EOF {
			break
		}
		lines = append(lines, string(line))
	}

	fmt.Println(lines)

	fmt.Println(getLines(lines))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func letterCount(s string) letters {
	m := letters{}
	for _, i := range s {
		m[i]++
	}
	return m
}

func getLines(str []string) letters {
	linesCh := make(chan map[rune]int)
	final := letters{}
	count := len(str)

	for _, s := range str {
		go func(s string) {
			linesCh <- letterCount(s)
		}(s)
	}
	for i := 0; i < count; i++ {
		for i, j := range <-linesCh {
			final[i] += j
		}
	}
	return final
}

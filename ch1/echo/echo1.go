package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf("fileName: %s\n", os.Args[0])
	for i, arg := range os.Args[1:] {
		fmt.Printf("os.Args[%d] = %s\n", i, arg)
	}
	stringConcatWithJoin()
	stringConcatWithoutJoin()
}

func stringConcatWithoutJoin() {
	var arguments string

	start := time.Now()
	fmt.Printf("start: %s\n", start)
	for _, arg := range os.Args[1:] {
		arguments += arg + " "
	}
	fmt.Printf("arguments: %s\n", arguments)
	fmt.Println("Time Taken: ", time.Since(start))
}

func stringConcatWithJoin() {
	var arguments string

	start := time.Now()
	fmt.Printf("start: %s\n", start)
	arguments = strings.Join(os.Args[1:], " ")
	fmt.Printf("arguments: %s\n", arguments)
	fmt.Println("Time Taken: ", time.Since(start))
}

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, make(map[string]int))
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, make(map[string]int))
			f.Close()
		}
	}
}

func countLines(file *os.File, counts map[string]int) {
	input := bufio.NewScanner(file)

	for input.Scan() {
		counts[input.Text()]++
	}

	duplicateFound := false
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s\n", n, line)
			duplicateFound = true
		}
	}

	if duplicateFound {
		fmt.Printf(" Duplicates File Name: %s\n", file.Name())
	}
}

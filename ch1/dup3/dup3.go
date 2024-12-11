package main

import (
	"fmt"
	"os"
	"strings"
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
	data, _ := os.ReadFile(file.Name())

	for _, d := range strings.Split(string(data), "\n") {
		counts[d]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d %s\n", n, line)
		}
	}
}

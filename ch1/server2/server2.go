package main

import (
	"fmt"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server1 Error: %v\n", err)
	}
}

func counter(writer http.ResponseWriter, request *http.Request) {
	mu.Lock()
	count++
	_, err := fmt.Fprintf(writer, "Count: %d\n", count)
	if err != nil {
		fmt.Printf("server1 Error: %v\n", err)
	}
	mu.Unlock()
}

func handler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "URL Path:=%q\n", request.URL.Path)
	if err != nil {
		fmt.Printf("server1 Error: %v\n", err)
	}
}

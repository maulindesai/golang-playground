package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("server1 Error: %v\n", err)
	}
}

func handler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "URL Path:=%q\n", request.URL.Path)
	if err != nil {
		fmt.Printf("server1 Error: %v\n", err)
	}
}

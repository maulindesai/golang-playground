package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	channel := make(chan string)
	for _, url := range os.Args[1:] {
		var prefix string
		if !strings.HasPrefix(url, "http://") {
			prefix = "http://"
		}
		url = prefix + url
		go fetch(url, channel)
	}

	for range os.Args[1:] {
		fmt.Println(<-channel)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	response, err := http.Get(url)

	if err != nil {
		ch <- fmt.Sprintf("fetch: %v\n", err)
		return
	}
	if response.StatusCode != 200 {
		ch <- fmt.Sprintf("fetch: %s\n", response.Status)
		return
	}
	nBytes, err := io.Copy(io.Discard, response.Body)
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v\n", err)
		return
	}
	elapsed := time.Since(start).Seconds()

	ch <- fmt.Sprintf("%.2fs %7d %s", elapsed, nBytes, url)
	err = response.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("fetch: %v\n", err)
		return
	}
}

package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		var prefix string
		if !strings.HasPrefix(url, "http://") {
			prefix = "http://"
		}
		url = prefix + url
		fmt.Printf("fetch: %s\n", url)
		fmt.Printf("fetch: %s\n", strings.Replace(url, "http://", "", 1))

		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
		}
		if response.StatusCode != 200 {
			fmt.Printf("fetch: %s\n", response.Status)
		}
		err = response.Write(os.Stdout)
		if err != nil {
			fmt.Printf("fetch: %v\n", err)
		}
		response.Body.Close()
	}
}

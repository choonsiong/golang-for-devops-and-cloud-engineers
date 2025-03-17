package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./http-get <url>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(os.Args[1]); err != nil {
		fmt.Println("invalid url:", err)
		os.Exit(1)
	}

	response, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal("http get error:", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("close body error:", err)
		}
	}(response.Body)

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Response Body:", string(bytes))
}

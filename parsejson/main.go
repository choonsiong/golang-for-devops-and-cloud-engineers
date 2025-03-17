package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func doReq(url string) (Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, errors.New("http get error")
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("close body error:", err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("invalid status code: " + strconv.Itoa(response.StatusCode))
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("io read all error")
	}

	var page Page
	err = json.Unmarshal(bytes, &page)
	if err != nil {
		return nil, errors.New("unmarshal error")
	}

	switch page.Name {
	case "words":
		var w Words
		err = json.Unmarshal(bytes, &w)
		if err != nil {
			return nil, errors.New("unmarshal error")
		}
		fmt.Printf("%#v\n", w)
		fmt.Printf("%v\n", strings.Join(w.Words, ","))
		return w, nil
	case "occurrence":
		var o Occurrence
		err = json.Unmarshal(bytes, &o)
		if err != nil {
			return nil, errors.New("unmarshal error")
		}
		fmt.Printf("%#v\n", o)
		for k, v := range o.Words {
			fmt.Printf("%s: %d\n", k, v)
		}
		return o, nil
	default:
		return nil, nil
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: ./parsejson <url>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(os.Args[1]); err != nil {
		fmt.Println("invalid url:", err)
		os.Exit(1)
	}

	resp, err := doReq(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	if resp == nil {
		fmt.Println("response is nil")
		os.Exit(1)
	}

	fmt.Println("Response:", resp.GetResponse())
}

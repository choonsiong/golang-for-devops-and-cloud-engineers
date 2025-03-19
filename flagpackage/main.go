package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func doReq(url string) (Response, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     "",
			Err:      err.Error(),
		}
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal("close body error:", err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     "",
			Err:      "invalid status code",
		}
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(bytes),
			Err:      err.Error(),
		}
	}

	if !json.Valid(bytes) {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     "",
			Err:      "no valid JSON response",
		}
	}

	var page Page
	err = json.Unmarshal(bytes, &page)
	if err != nil {
		return nil, RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(bytes),
			Err:      err.Error(),
		}
	}

	switch page.Name {
	case "words":
		var w Words
		err = json.Unmarshal(bytes, &w)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(bytes),
				Err:      err.Error(),
			}
		}
		fmt.Printf("%#v\n", w)
		fmt.Printf("%v\n", strings.Join(w.Words, ","))
		return w, nil
	case "occurrence":
		var o Occurrence
		err = json.Unmarshal(bytes, &o)
		if err != nil {
			return nil, RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(bytes),
				Err:      err.Error(),
			}
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
		fmt.Println("usage: ./prog <url>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(os.Args[1]); err != nil {
		fmt.Println("invalid url:", err)
		os.Exit(1)
	}

	resp, err := doReq(os.Args[1])
	if err != nil {
		if newErr, ok := err.(RequestError); ok {
			fmt.Println(newErr.HTTPCode, newErr.Body, newErr.Err)
		}

		log.Fatal(err.Error())
	}

	if resp == nil {
		fmt.Println("response is nil")
		os.Exit(1)
	}

	fmt.Println("Response:", resp.GetResponse())
}

package main

import (
	"fmt"
	"io"
	"log"
	"time"
)

type SlowReader struct {
	content string
	pos     int
}

func (s *SlowReader) Read(p []byte) (n int, err error) {
	if s.pos < len(s.content) {
		log.Printf("Reading '%s' and sleep for a while...", string(s.content[s.pos]))
		time.Sleep(time.Second * 1)
		n = copy(p, s.content[s.pos:s.pos+1])
		s.pos += 1
		return n, nil
	}

	log.Println("We are done reading 完成了，ありがとうございました 💁🏿💁🏿💁🏿...")

	return 0, io.EOF
}

func main() {
	slowReader := SlowReader{
		content: "hello world",
	}

	bytes, err := io.ReadAll(&slowReader)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Bytes: %v\n", string(bytes))
}

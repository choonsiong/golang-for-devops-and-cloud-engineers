package main

import (
	"fmt"
	"io"
	"log"
)

type SlowReader struct {
	content string
	pos     int
}

func (s *SlowReader) Read(p []byte) (n int, err error) {
	if s.pos < len(s.content) {
		n = copy(p, s.content[s.pos:s.pos+1])
		s.pos += 1
		return n, nil
	}

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

package main

import (
	"fmt"
	"strings"
)

type Words struct {
	Page
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func (w Words) GetResponse() string {
	return fmt.Sprintf("%s", strings.Join(w.Words, ","))
}

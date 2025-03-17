package main

import (
	"fmt"
	"strings"
)

type Occurrence struct {
	Page
	Words map[string]int `json:"words"`
}

func (o Occurrence) GetResponse() string {
	var arr []string // create a nil slice
	for k, v := range o.Words {
		arr = append(arr, fmt.Sprintf("%s:%d", k, v))
	}
	return fmt.Sprintf("%s", strings.Join(arr, ","))
}

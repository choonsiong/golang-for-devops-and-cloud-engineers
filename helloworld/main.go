package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	args := os.Args[1:]
	fmt.Println(args)
}

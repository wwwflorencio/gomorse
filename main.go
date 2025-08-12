package main

import (
	"fmt"
	"github.com/wwwflorencio/gomorse/morse"
	"os"
)

func help() {
	fmt.Println("usage example: ./morse \"<code>\"")
}

func validateArgs() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
}

func main() {
	validateArgs()

	input := os.Args[1]

	decoder := morse.NewMorseDecoder()
	fmt.Println(decoder.Decode(input))
}

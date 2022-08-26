package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("s5n (shorten) : shortens passed in argument\n")
		fmt.Printf("usage: s5n <word to shorten>\n")
		fmt.Printf("v0.0.1\n")
		os.Exit(1)
	}

	input := os.Args[1]

	if len(input) > 2 {
		firstChar := input[0:1]
		lastChar := input[len(input)-1:]
		abbrLen := len(input) - 2

		fmt.Printf("%s%d%s\n", firstChar, abbrLen, lastChar)
	} else {
		fmt.Printf("%s\n", input)
	}

	os.Exit(0)
}

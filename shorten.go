package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// piped
		fmt.Printf("we got it from the stream!!!")
		processPipe()
	} else {
		// from terminal

		if len(os.Args) != 2 {
			fmt.Printf("shorten : shortens passed in argument\n")
			fmt.Printf("usage: shorten <word to shorten>\n")
			fmt.Printf("v0.0.3\n")
			os.Exit(1)
		}

		output := shorten(os.Args[1])

		fmt.Printf(output)
	}

	os.Exit(0)
}

func processPipe() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf(shorten(scanner.Text()))
	}
}

/*
 * does the shortening of input returning a shortened string (if it can shorten it)
 */
// TODO : could the implementation/logic be improved?
func shorten(input string) string {
	if len(input) <= 2 {
		return input
	}

	firstChar := input[0:1]
	lastChar := input[len(input)-1:]
	abbrLen := len(input) - 2

	return fmt.Sprintf("%s%d%s\n", firstChar, abbrLen, lastChar)
}

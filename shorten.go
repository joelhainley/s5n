package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// piped
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
		fmt.Printf("%s\n", output)
	}

	os.Exit(0)
}

func processPipe() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Printf("%s\n", processLine(scanner.Text()))
	}
}

// breaks up a line into tokens using spaces
// shortens all of the tokens then reconstructs the line and returns it
func processLine(line string) string {
	words := strings.Fields(line)
	shortened := make([]string, len(words))
	for i, s := range words {
		shortened[i] = shorten(s)
	}

	return strings.Join(shortened, " ")
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

	return fmt.Sprintf("%s%d%s", firstChar, abbrLen, lastChar)
}

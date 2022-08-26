package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) != 2 {
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

/*
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

const DEFAULT_OUTPUT_FORMAT = "%t.%s.%n"

var debugMode bool
var format string

func main() {
	// setup the flags
	flag.BoolVar(&debugMode, "debug", false, "turn on debug mode")
	flag.StringVar(&format, "format", DEFAULT_OUTPUT_FORMAT, "set the output format")
	os.Args
	flag.Parse()

	if debugMode {
		fmt.Printf("debug mode is enabled")
	}

	fi, _ := os.Stdin.Stat()

	// TODO leverage this logic to do the right thing
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		fmt.Println("data is from pipe")
		scanner := bufio.NewScanner(os.Stdin)

		for scanner.Scan() {
			//fileInfo := dumpFileAttributes(scanner.Text(), format)
			//fmt.Printf("> %s\n", fileInfo)
			fmt.Printf("boom boom")
		}
	} else {
		fmt.Println("data is from terminal")

		fmt.Sprintf("%s", "sss")
	}

}
*/

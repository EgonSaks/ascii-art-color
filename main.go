package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("\nError: No arguments provided")
		os.Exit(0)
	}

	err := "\nUsage: go run . [STRING] [OPTION] \n\nEX: go run . --color=<color> \"something\" <letters to be colored> "

	if string(os.Args[1]) == "--color" || strings.HasPrefix(string(os.Args[1]), "-color") {
		fmt.Println(err)
		os.Exit(0)
	} else {
		flag.Parse()

	}
	args := flag.Args()

	banners := "banner/standard.txt"
	str := args[0]

	letter := mapFont(banners)
	sentence := mapInput(str, letter)
	output := printInput(str, sentence, args)

	colorPrint(output)

}

package main

import (
	"ascii-art-custom/art"
	"ascii-art-custom/color"
	"ascii-art-custom/output"
	"ascii-art-custom/reverse"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) <= 1 {
		fmt.Println("Usage: go run . [OPTION]")
		fmt.Println()
		fmt.Println("EX: go run . <text>")
		return
	}

	// flagExist := false

	if strings.HasPrefix(os.Args[1], "--") {
		// flagExist = true
		switch {
		case strings.HasPrefix(os.Args[1], "--reverse") && !(strings.Contains(os.Args[1], "=")):
			fmt.Println("Usage: go run . [OPTION]")
			fmt.Println()
			fmt.Println("EX: go run . --reverse=<fileName>")
			return
		case strings.Contains(os.Args[1], "--color") && !(strings.Contains(os.Args[1], "=")):
			fmt.Println("Usage: go run . [OPTION] [STRING]")
			fmt.Println()
			fmt.Println("EX: go run . --color=<color> <letters to be colored> 'something'")
			return
		case strings.Contains(os.Args[1], "--output") && !(strings.Contains(os.Args[1], "=")):
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
			fmt.Println()
			fmt.Println("EX: go run . --output=<fileName.txt> something standard")
			return
			// Add more cases for different incorrect flags.
		}
	}

	reverseFlag := flag.String("reverse", "", "ascii-art-reverse")
	colorFlag := flag.String("color", "", "ascii-art-color")
	outputFlag := flag.String("output", "", "ascii-art-output")
	// Add more flags for different processes.
	flag.Parse()

	var args []string

	font := "fonts/standard.txt"
	// Handle the fonts.
	switch {
	case os.Args[len(os.Args)-1] == "thinkertoy":
		args = os.Args[1 : len(os.Args)-1]
		font = "fonts/thinkertoy.txt"
	case os.Args[len(os.Args)-1] == "shadow":
		args = os.Args[1 : len(os.Args)-1]
		font = "fonts/shadow.txt"
	case os.Args[len(os.Args)-1] == "standard":
		args = os.Args[1 : len(os.Args)-1]
		font = "fonts/standard.txt"
	default:
		args = os.Args[1:]
	}

	var result string

	switch {
	case *reverseFlag != "":
		reverse.ProcessReverse(args[0], font)
	case *colorFlag != "":
		result = art.ProcessArt(args[1:], font)
		color.ProcessColor(args, result, font)
		/* switch {
		case len(args) == 2:
			result = art.ProcessArt(args[1:], font)
			color.ProcessColor(result, font)
		case len(args) == 3:
			result = art.ProcessArt(args[2:], font)
			color.ProcessColor(result, font)
		} */
	case *outputFlag != "":
		output.ProcessOutput(font)
	default:
		result = art.ProcessArt(args[:], font)
		fmt.Print(result)
	}
}

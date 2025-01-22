package color

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
	purple = "\033[35m"
	cyan   = "\033[36m"
	orange = "\033[38;2;255;165;0m"
	reset  = "\033[0m"
)

func ProcessColor(args []string, result, font string) {

	color := SelectColor(args[0])

	switch {
	case len(args) == 2:
		SimpleColor(result, color)
	case len(args) == 3:
		text := args[2]
		ComplexColor(text, color, font)
	}
}

func SelectColor(arg string) (color string) {
	switch {
	case strings.Contains(arg, "red"):
		return red
	case strings.Contains(arg, "green"):
		return green
	case strings.Contains(arg, "yellow"):
		return yellow
	case strings.Contains(arg, "blue"):
		return blue
	case strings.Contains(arg, "purple"):
		return purple
	case strings.Contains(arg, "cyan"):
		return cyan
	case strings.Contains(arg, "orange"):
		return orange
	default:
		return reset
	}

}

func SimpleColor(result, color string) {
	fmt.Print(string(color))
	fmt.Print(result)
}

func ComplexColor(text, color, font string) {

	// Deal with invalid ascii characters --> Print the error.
	for _, char := range text {
		if char < 32 || char > 126 {
			fmt.Printf("Invalid Ascii character: %s", string(char))
			fmt.Println()
			return
		}
	}

	// Open the correct font text file.
	file, err := os.ReadFile(font)
	if err != nil {
		log.Fatal(err)
		return
	}

	// Split up the font text file into lines.
	filecontents := string(file)
	filecontents = strings.ReplaceAll(filecontents, "\r\n", "\n")
	filelines := strings.Split(filecontents, "\n")

	// Split up the arg text into phrases.
	phrases := strings.Split(text, "\\n")
	letters := os.Args[2]

	// Variable to store the line number for the font text file.
	var a int
	// Variable to store the ascii result
	var result string

	for _, phrase := range phrases {
		positions := PhraseOverlap(phrase, letters)
		// positions := IndexOverlap(phrase, letters)
		switch {
		// If multiple "\n", make next "\n" into single new line instead of 8 new lines
		case phrase == "":
			result = result + "\n"
		default:
			for counter := 0; counter < 8; counter++ {
				// To iterate over the characters in the phrase
				for i, char := range phrase {
					for _, position := range positions {
						if i == position {
							fmt.Print(color)
							break
						} else {
							fmt.Print(reset)
						}
					}
					a = 9*int(char) - 287
					fmt.Print(filelines[a+counter])
				}
				fmt.Println()
			}
		}
	}
}

func IndexOverlap(text, letters string) []int {
	var positions []int
	for index, char := range text {
		for _, letter := range letters {
			if char == letter {
				positions = append(positions, index)
			}
		}
	}
	return positions
}

func PhraseOverlap(mainstring, substring string) []int {
	var startis []int
	endi := 0
	for {
		index := strings.Index(mainstring[endi:], substring)
		if index == -1 {
			break
		}
		startis = append(startis, endi+index)
		endi += index + len(substring)
	}
	var allindexes []int
	for _, starti := range startis {
		for i := starti; i < starti+len(substring); i++ {
			allindexes = append(allindexes, i)
		}
	}
	return allindexes
}

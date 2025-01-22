package art

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ProcessArt(args []string, font string) string {
	var text string
	for i := range args {
		switch {
		case i != len(args)-1:
			text = text + args[i] + " "
		default:
			text = text + args[i]
		}
	}
	text, work := Generate(text, font)
	if !work {
		return ""
	} else {
		return text
	}
}

func Generate(text string, font string) (string, bool) {

	// Deal with invalid ascii characters --> Print the error.
	for _, char := range text {
		if char < 32 || char > 126 {
			fmt.Printf("Invalid Ascii character: %s", string(char))
			fmt.Println()
			return "", false
		}
	}

	// Open the correct font text file.
	file, err := os.ReadFile(font)
	if err != nil {
		log.Fatal(err)
		return "", false
	}

	// Split up the font text file into lines.
	filecontents := string(file)
	filecontents = strings.ReplaceAll(filecontents, "\r\n", "\n")
	filelines := strings.Split(filecontents, "\n")

	// Split up the arg text into phrases.
	phrases := strings.Split(text, "\\n")

	// Variable to store the line number for the font text file.
	var a int
	// Variable to store the ascii result
	var result string

	// To iterate over the phrases (phrases split by typed "\n")
	for _, phrase := range phrases {
		switch {
		// If multiple "\n", make next "\n" into single new line instead of 8 new lines
		case phrase == "":
			result = result + "\n"
		default:
			for counter := 0; counter < 8; counter++ {
				// To iterate over the characters in the phrase
				for _, char := range phrase {
					// To calculate for the line number
					a = 9*int(char) - 287
					result = result + filelines[a+counter]
				}
				result = result + "\n"
			}
		}
		// To iterate over the 8 lines of the ascii font
	}
	return result, true

}

// Everything below here is old and unused. Keeping for reference.

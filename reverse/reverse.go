package reverse

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func ProcessReverse(arg, font string) {
	var path string
	// Is the file name "example..."?
	if strings.Contains(arg, "example") {
		// If it is, look in "examples/"
		path = "examples/"
	} else {
		// If it is not, just look in the root folder.
		path = ""
	}
	// Read the contents of the example.txt file given in the arguments.
	asciiFile, err := os.ReadFile(path + strings.Split(arg, "=")[1])
	if err != nil {
		log.Fatal(err)
	}
	// Format the contents of the example.txt file to read vertically.
	asciiFormat := HorizontaltoVertical(string(asciiFile))

	// Read the contents of the chosen font file.
	fontFile, err := os.ReadFile(font)
	if err != nil {
		log.Fatal(err)
	}
	// Format the contents of the font file.
	fontFormat := HorizontaltoVertical(string(fontFile))

	result := FindAscii(asciiFormat, fontFormat)

	fmt.Println(result)
}

func HorizontaltoVertical(text string) map[int]string {
	// This has to be to make sure the inital formatting of the font files is
	// not an issue.
	text = strings.ReplaceAll(text, "\r\n", "\n")

	mapresult := make(map[int]string)

	var storage string

	lines := strings.Split(text, "\n")

	for lineposition := 0; lineposition < len(lines); lineposition = lineposition + 9 {
		// The font files have an empty first line. We want to skip over it.
		if len(lines[lineposition]) < 1 {
			lineposition++
		}
		if len(lines[lineposition]) > 0 {
			for b := 0; b < len(lines[lineposition]); b++ {
				for a := 0; a <= 7; a++ {
					storage = storage + string(lines[lineposition+a][b])
				}
			}
			// fmt.Println(storage)
			// fmt.Println("POSITION:	", lineposition)
			// fmt.Println("LENGTH:		", len(lines[lineposition]))
			mapresult[(lineposition-1)/9] = storage
			storage = ""
		}
	}
	return mapresult
}

func FindAscii(asciiFormat, fontFormat map[int]string) string {

	var result string

	for i := 0; i < len(fontFormat); i++ {

		remainingAscii, found := strings.CutPrefix(asciiFormat[0], fontFormat[i])
		if found {
			result = result + string(rune(i+32))
			asciiFormat[0] = remainingAscii
			i = -1
		}
	}
	return result
}

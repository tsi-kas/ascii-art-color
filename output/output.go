package output

import (
	"ascii-art-custom/art"
	"os"
	"strings"
)

func ProcessOutput(font string) {
	var args []string
	switch {
	case os.Args[len(os.Args)-1] == "shadow":
		args = os.Args[2 : len(os.Args)-1]
	case os.Args[len(os.Args)-1] == "thinkertoy":
		args = os.Args[2 : len(os.Args)-1]
	default:
		args = os.Args[2:]
	}
	contents := art.ProcessArt(args, font)
	filecreate := strings.Split(os.Args[1], "=")[1]
	CreateFile(filecreate, contents)
}

func CreateFile(filecreate, contents string) {
	file, err := os.Create(filecreate)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(contents)
	if err != nil {
		return
	}
}

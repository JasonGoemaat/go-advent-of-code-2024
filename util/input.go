package util

import (
	"fmt"
	"io"
	"os"
)

var StdinFlag = false
var InputFile = ""

func GetContent() string {
	content := ""
	if StdinFlag {
		bytes, err := io.ReadAll(os.Stdin)
		if err != nil {
			panic(fmt.Sprint("Error reading from stdin:", err))
		}
		content = string(bytes)
	} else if InputFile != "" {
		file, err := os.Open(InputFile)
		if err != nil {
			panic(fmt.Sprintf("Error reading from '%s': %v", InputFile, err))
		}
		defer file.Close()
		bytes, err := io.ReadAll(file)
		content = string(bytes)
	}
	return content
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

//
func main() {
	if len(os.Args) <= 2 {
		fmt.Println("usage: go run escapejson.go input-file-path output-file-path")
		return
	}

	inputFilePath := os.Args[1]
	outputPath := os.Args[2]

	if _, err := os.Stat(inputFilePath); os.IsNotExist(err) {
		fmt.Println("ERROR: file does not exists.")
		return
	}

	jsonStr, err := ioutil.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("ERROR: failed to read file.")
		return
	}

	// Unicode characters to Unicode escape sequences.
	converted := strconv.QuoteToASCII(string(jsonStr))

	// Undo - return code
	re := regexp.MustCompile(`([^\\])\\n`)
	converted = string(re.ReplaceAll([]byte(converted), []byte("$1\n")))

	// Undo - return code in value
	converted = strings.ReplaceAll(converted, `\\n`, `\n`)

	// Undo - double quote
	converted = strings.ReplaceAll(converted, `\"`, `"`)

	// Remove unnecessary double quote
	converted = converted[1 : len(converted)-1]

	err = ioutil.WriteFile(outputPath, []byte(converted), 0644)
	if err != nil {
		fmt.Println("ERROR: failed to output.")
	}
	fmt.Sprintf("done. file=%s\n", outputPath)
}

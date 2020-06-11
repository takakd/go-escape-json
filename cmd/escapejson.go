package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

//
func convertJsonElement(data interface{}) interface{} {
	switch v := data.(type) {
	case map[string]interface{}:
		for ik, iv := range v {
			v[ik] = convertJsonElement(iv)
		}
		data = v
	case []interface{}:
		for ik, iv := range v {
			v[ik] = convertJsonElement(iv)
		}
		data = v
	case string:
		v = fmt.Sprintf("%+q", v)
		// Remove unnecessary double-quote
		data = v[1 : len(v)-1]
	default:
	}
	return data
}

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

	// Convert
	var data interface{}
	json.Unmarshal(jsonStr, &data)
	convertedData := convertJsonElement(data)

	// Pretty
	indentedData, err := json.MarshalIndent(convertedData, "", "  ")
	if err != nil {
		fmt.Println("ERROR: failed to MarshalIndent.")
		return
	}

	// Output file
	err = ioutil.WriteFile(outputPath, indentedData, 0644)
	if err != nil {
		fmt.Println("ERROR: failed to output.")
	}
	fmt.Sprintf("done. file=%s\n", outputPath)
}

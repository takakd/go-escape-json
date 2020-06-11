package main

import (
	"os"
	"path/filepath"
	"testing"
)

func Test_main(t *testing.T) {
	oldArgs := os.Args
	defer func() {
		os.Args = oldArgs
	}()

	wd, _ := os.Getwd()
	filePath, _ := filepath.Abs(wd + "/../testdata/input.json")
	outputPath, _ := filepath.Abs(wd + "/../testdata/output.json")
	os.Args = []string{"", filePath, outputPath}

	main()
}

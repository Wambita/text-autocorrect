package goreloaded

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenFile(t *testing.T) {
	// test for valid file path
	inputFilepath := "sample.txt"
	outputFilepath := "result.txt"
	err := OpenFile(inputFilepath, outputFilepath)
	if err != nil {
		t.Errorf("Failed to open existing file: %v", err)
	}
}

// test for non existent file
func TestNonExistentFile(t *testing.T) {
	inputFilepath := "non-existent.txt"
	outputFilepath := "nosuchfile.txt"
	err := OpenFile(inputFilepath, outputFilepath)
	if err == nil {
		t.Errorf("Expected error for non existent file, but got none")
	}
}

// test for wrong extension
func TestWrongExtension(t *testing.T) {
	inputFilepath := "image.png"
	if filepath.Ext(inputFilepath) == ".txt" {
		t.Errorf("Error: file has .txt expected wrong filepath")
	}
	outputFilepath := "image.png"
	if filepath.Ext(outputFilepath) == ".txt" {
		t.Errorf("Error: file has .txt expected wrong filepath")
	}
}

// test empty file
func TestEmptyfile(t *testing.T) {
	inputFilepath := "empty.txt"
	fileInfo, err := os.Stat(inputFilepath)
	if err != nil {
		t.Errorf("error: %s", err)
	}
	if fileInfo.Size() != 0 {
		t.Errorf("Error: file is not empty. Expected empty file")
	}
}

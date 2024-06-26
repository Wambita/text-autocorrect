package goreloaded

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

// Function to open input file and create output file if one does not exist
func OpenFile(inputFilepath string, outputFilepath string) error {
	//handling possible errors
	fileInfo, err := os.Stat(inputFilepath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist")
		return err

	}
	if err != nil {
		fmt.Printf("Error checking the file: %s\n", err)
		return err
	}
	// check if the file size is less than zero
	if fileInfo.Size() == 0 {
		fmt.Println("File exists but is empty")
		return err
	}

	if filepath.Ext(inputFilepath) != ".txt" {
		fmt.Printf("Wrong file extension. Use .txt files")
		return err
	}

	inputFile, err := os.Open(inputFilepath)
	if err != nil {
		fmt.Println(err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create(outputFilepath)
	if filepath.Ext(outputFilepath) != ".txt" {
		fmt.Printf("Wrong file extension. Use .txt files")
	}

	if err != nil {
		fmt.Println(err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		lines := scanner.Text()
		fmt.Fprintln(outputFile, lines)
		fmt.Println(lines)
	}
	return nil
}

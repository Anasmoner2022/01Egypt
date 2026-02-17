package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Usage: go run . <input_file> <output_file>")
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]
	fmt.Printf("Input: %s, Output: %s\n", inputFile, outputFile)

}

func readFile(filename string) (string, error) {
	// Use os.ReadFile to read entire file

	// Return content as string and error
}

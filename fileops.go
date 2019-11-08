package main

import (
	"bufio"
	"fmt"
	"os"
)

// readLinesFromFile reads a whole file into memory
// and returns a slice of its lines.
func readLinesFromFile(path string) ([]string, error) {

	file, err := os.OpenFile(path, os.O_RDONLY, os.ModeAppend)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Print("File Does Not Exist: ")
		}
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

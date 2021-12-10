package utilities

import (
	"bufio"
	"os"
)

func ReadFile(fileName string) *os.File  {
	fileH, err := os.Open(fileName)
	if err != nil {
		return nil
	} else {
		return fileH
	}
}

func ReadLinesIntoStringArr(fileH *os.File) []string {
	var lines []string
	scanner := bufio.NewScanner(fileH)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
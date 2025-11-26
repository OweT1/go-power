package utils

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

func OpenFile(file_path string) (*os.File, error) {
	return os.Open(file_path)
}

func parseLines(lines [][]string) []Problem {
    // Make a slice of Problems with the exact size we need
	ret := make([]Problem, len(lines))
    
	for i, line := range lines {
		ret[i] = Problem{
			Question: line[0], // Column 1
			Answer:   strings.TrimSpace(line[1]), // Column 2 (remove spaces)
		}
	}
	return ret
}

func ParseCsv(file *os.File) []Problem {
	csvReader := csv.NewReader(file)
	lines, err := csvReader.ReadAll() // This returns a slice of slices: [][]string
	if err != nil {
		fmt.Printf("Failed to parse the provided CSV file.")
		os.Exit(1)
	}

	// Convert CSV lines to Problem struct
	problems := parseLines(lines)
	return problems
}
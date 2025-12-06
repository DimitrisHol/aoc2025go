package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	// testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) int {

	total := 0

	rows := len(data) - 1 // remove one since the last one contains the operations
	columns := len(strings.Fields(data[0]))

	numbers := make([][]int, rows)
	operations := make([]string, columns)

	for i, row := range data {

		// Remove the whitespace and split
		row := strings.Fields(row)

		if i == rows {
			for j, v := range row {
				operations[j] = v
			}
			break
		}

		// Create local row numbers
		rowNumbers := make([]int, columns)
		for j, v := range row {

			intNumber, _ := strconv.Atoi(v)

			rowNumbers[j] = intNumber
		}

		numbers[i] = rowNumbers

	}

	for i := 0; i < columns; i++ {

		localSum := 0
		if operations[i] == "*" {
			localSum = 1
		}

		for j := 0; j < rows; j++ {

			if operations[i] == "*" {
				localSum *= numbers[j][i]
			} else {
				localSum += numbers[j][i]
			}
		}
		total += localSum
	}

	return total
}

func main() {

	data := parseFile("06.txt")
	fmt.Println(part1(data))
}

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

	validRegions := 0

	for index, row := range data {

		if index < 30 {
			// LAMAO
			continue
		} else {

			firstSplit := strings.Split(row, ":")

			dimensions := strings.Split(firstSplit[0], "x")

			width, _ := strconv.Atoi(dimensions[0])
			length, _ := strconv.Atoi(dimensions[1])

			totalArea := width * length

			quantities := strings.Split(firstSplit[1], " ")

			requiredArea := 0

			for i := 0; i < 6; i++ {

				reqShape, _ := strconv.Atoi(quantities[i])
				reqShape *= 9

				requiredArea += reqShape
			}

			if totalArea > requiredArea {
				validRegions += 1
			}
		}

	}

	return validRegions
}

func main() {

	data := parseFile("12.txt")
	fmt.Println(part1(data))
}

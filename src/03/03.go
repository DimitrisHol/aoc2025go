package main

import (
	// "fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"

	// "slices"
	// "strconv"
	"strings"
)

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	// testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(banks []string) {

	sumPart1 := 0
	for _, v := range banks {
		sumPart1 += calculateJolts1(v)
	}

	println("Part 1 :", sumPart1)
}

func calculateJolts1(bank string) int {

	// Conver the string into integers
	numbers := make([]int, len(bank))
	for i, digit := range bank {
		numbers[i] = int(digit - '0')
	}

	for i := 9; i >= 1; i-- {

		// Find the index of the largest number
		maxIndex := slices.Index(numbers, i)

		// Skip if largest number not found or is at the end
		if maxIndex == -1 || maxIndex == len(numbers)-1 {
			// println("maximum index not found or at end ", maxIndex, "out of", len(numbers)-1)
			continue
		}

		// Seek right for the next largest number
		max := 0
		for j := maxIndex + 1; j < len(numbers); j++ {

			if numbers[j] > max {
				max = numbers[j]
			}
		}

		// i concat max
		result, _ := strconv.Atoi(strconv.Itoa(i) + strconv.Itoa(max))
		return result
	}

	println("horrible mistake")
	return 0
}

func main() {

	banks := parseFile("03.txt")
	part1(banks)
}

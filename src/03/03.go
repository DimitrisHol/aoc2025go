package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
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

func solution(banks []string) (int, int64) {

	part1 := 0
	var part2 int64 = 0
	for _, v := range banks {

		part1 += calculateJolts1(v)
		part2 += calculateJolts2(v)
	}

	return part1, part2
}

func calculateJolts1(bank string) int {

	// Convert the string into integers
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

func calculateJolts2(bank string) int64 {
	// Conver the string into integers
	numbers := make([]int, len(bank))
	for i, digit := range bank {
		numbers[i] = int(digit - '0')
	}

	resultNumbers := []int{}
	var window = 12 // How many digits we need to activate

	maxIndex := findFirstLargestNumberVol(numbers, window)
	resultNumbers = append(resultNumbers, numbers[maxIndex])

	var startingIndex = maxIndex + 1 // We start right after the largest number
	var digitsToadd = window - 1     // The first one was already found

	for digitsToadd > 0 {

		// Seek right for the next largest number
		max := 0
		maxIndex = -1

		for j := startingIndex; j < len(numbers)-digitsToadd+1; j++ {

			if numbers[j] > max {
				max = numbers[j]
				maxIndex = j
			}
		}

		digitsToadd -= 1
		startingIndex = maxIndex + 1
		resultNumbers = append(resultNumbers, max)
	}

	resultString := ""

	for _, v := range resultNumbers {
		resultString += strconv.Itoa(v)
	}

	result, _ := strconv.ParseInt(resultString, 10, 64)
	return result
}

func findFirstLargestNumberVol(numbers []int, window int) int {

	// 15 numbers, window = 12 numbers.
	// [0, 1, 2, 3 || 4]
	// 15 - 12 = 3
	// Then since slice range is non-inclusive we add one 3 + 1 = 4
	firstNumberRange := len(numbers) - window + 1

	var maxValue = 0
	var maxIndex = -1

	for i := 0; i < firstNumberRange; i++ {

		if numbers[i] > maxValue {
			maxValue = numbers[i]
			maxIndex = i
		}
	}

	return maxIndex

}

func main() {

	banks := parseFile("03.txt")
	sumPart1, sumPart2 := solution(banks)

	fmt.Println("Part 1 :", sumPart1)
	fmt.Println("Part 2 :", sumPart2)
}

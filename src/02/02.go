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
	return strings.Split(string(data), ",")
}

func part1(ranges []string) (int, int) {

	totalSumPart1 := 0
	totalSumPart2 := 0

	for _, v := range ranges {

		standEndRange := strings.Split(v, "-")

		startRange, _ := strconv.Atoi(standEndRange[0])
		endRange, _ := strconv.Atoi(standEndRange[1])

		for i := startRange; i <= endRange; i++ {

			if invalidIdPart1(strconv.Itoa(i)) {
				totalSumPart1 += i
			}

			if invalidIdPart2(strconv.Itoa(i)) {
				totalSumPart2 += i
			}
		}
	}

	return totalSumPart1, totalSumPart2
}

func invalidIdPart1(id string) bool {

	length := len(id)

	firstPart := id[0 : length/2]
	secondPart := id[(length / 2):length]

	return firstPart == secondPart

}

func invalidIdPart2(id string) bool {

	// [0, 3] = [121 121]    , next [3:6]
	// [0, 2] = [12 12 12]   , next [2:4], [4:6]
	// [0, 1] = [1 2 1 2 1 2], next [1:2], [2:3], [3:4], [4:5], [5:6]

	length := len(id)

	// Search for the first split in x parts that makes it invalid
	for parts := 2; parts <= 10; parts++ {

		// Exit early if you can't divide the id into equal blocks
		if length%parts != 0 || length < parts {
			continue
		}

		// println("Id : ", id, "trying to split into ", parts, "parts.")

		blockSize := length / parts
		firstBlock := id[0:blockSize]

		allBlocksEqual := true

		// Check if all other blocks are equal to the first block
		for i := 1; i < parts; i++ { // loop for parts -1 times (the rest of the blocks)

			start := i * blockSize
			stop := (i * blockSize) + blockSize

			nextBlock := id[start:stop]

			// println("checking first block equals with ", i+1, "block :", firstBlock, "==", nextBlock)
			if firstBlock != nextBlock {
				allBlocksEqual = false
				break
			}
		}

		if allBlocksEqual {
			return true
		}
	}

	return false
}

func main() {

	ranges := parseFile("02.txt")
	fmt.Println(part1(ranges))
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	// "strconv"
	"strings"
)

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	// testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), ",")
}

func part1(ranges []string) int {

	totalSumPart1 := 0
	for _, v := range ranges {

		standEndRange := strings.Split(v, "-")

		startRange, _ := strconv.Atoi(standEndRange[0])
		endRange, _ := strconv.Atoi(standEndRange[1])

		for i := startRange; i <= endRange; i++ {

			if invalidIdPart1(strconv.Itoa(i)) {
				totalSumPart1 += i
			}
		}
	}

	return totalSumPart1
}

func invalidIdPart1(id string) bool {

	length := len(id)

	firstPart := id[0 : length/2]
	secondPart := id[(length / 2):length]

	return firstPart == secondPart

}

func main() {

	ranges := parseFile("02.txt")
	fmt.Println(part1(ranges))
}

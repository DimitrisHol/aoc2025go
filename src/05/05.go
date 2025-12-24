package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
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

	freshIngredients := 0

	ranges := make([][]int, 0)

	var firstPart bool = true

	for _, v := range data {

		if v == "" {
			firstPart = false
			continue
		}

		// Parse the ranges
		if firstPart {

			rangeData := strings.Split(v, "-")

			localRange := make([]int, 2)
			localRange[0], _ = strconv.Atoi(rangeData[0])
			localRange[1], _ = strconv.Atoi(rangeData[1])

			ranges = append(ranges, localRange)
		} else { // These are the ids to check

			id, _ := strconv.Atoi(v)
			for _, r := range ranges {

				if id >= r[0] && id <= r[1] {

					freshIngredients += 1
					break
				}
			}
		}

	}

	return freshIngredients
}

func part2(data []string) int {

	ranges := make([][]int, 0)

	for _, v := range data {

		if v == "" {
			// Part 2 we only need the ranges
			break
		}

		// Parse the ranges
		rangeData := strings.Split(v, "-")

		localRange := make([]int, 2)
		localRange[0], _ = strconv.Atoi(rangeData[0])
		localRange[1], _ = strconv.Atoi(rangeData[1])

		// Add them all, then we sort them
		ranges = append(ranges, localRange)
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	mergedRanges := [][]int{}

	// Starting range
	currentStart := ranges[0][0]
	currentStop := ranges[0][1]

	// Loop through the ranges.
	for i := 1; i < len(ranges); i++ {

		nextStart := ranges[i][0]
		nextStop := ranges[i][1]

		if currentStop >= nextStart { // Overlap

			currentStop = max(currentStop, nextStop)
		} else {

			// No overlap, save range and move to the next one
			mergedRanges = append(mergedRanges, []int{currentStart, currentStop})
			currentStart = nextStart
			currentStop = nextStop
		}

		if i == len(ranges)-1 {
			mergedRanges = append(mergedRanges, []int{currentStart, currentStop})
		}
	}

	freshIngredients := 0
	for i := 0; i < len(mergedRanges); i++ {
		freshIngredients += mergedRanges[i][1] - mergedRanges[i][0] + 1
	}

	return freshIngredients
}

func main() {

	data := parseFile("05.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

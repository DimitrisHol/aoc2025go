package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	// testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(rows []string) int {

	var numberOfRows = len(rows)
	var numberOfCols = len(rows[0])
	grid := makeGrid(rows)

	var requiredAccessSpots = 3
	var validPaperRolls = 0

	for i := 1; i <= numberOfRows; i++ {
		for j := 1; j <= numberOfCols; j++ {

			if grid[i][j] == '@' {
				// We are at a paper roll, check all corners
				if paperRollCanBeAccessed(grid, requiredAccessSpots, i, j, false) {
					validPaperRolls += 1
				}
			}

		}
	}

	return validPaperRolls
}

func part2(rows []string) int {

	var numberOfRows = len(rows)
	var numberOfCols = len(rows[0])
	grid := makeGrid(rows)

	var requiredAccessSpots = 3
	var validPaperRolls = 0

	for i := 0; i < 100; i++ { // Or until no more :)

		for i := 1; i <= numberOfRows; i++ {
			for j := 1; j <= numberOfCols; j++ {

				if grid[i][j] == '@' {
					// We are at a paper roll, check all corners
					if paperRollCanBeAccessed(grid, requiredAccessSpots, i, j, true) {
						validPaperRolls += 1
					}
				}

			}
		}
	}

	// Print
	for i := range grid {
		fmt.Println(string(grid[i]))
	}

	return validPaperRolls

}

func makeGrid(data []string) [][]rune {

	var numberOfRows = len(data)
	var numberOfCols = len(data[0])

	// Make a larger grid 10x10 - > 12x12 with dots
	grid := make([][]rune, numberOfRows+2)
	for r := range grid {
		grid[r] = make([]rune, numberOfCols+2)
		for c := range grid[r] {
			grid[r][c] = 'W'
		}

	}

	// Replace the inner circle with the actual values
	for i := 0; i < numberOfRows; i++ {
		row := []rune(data[i])
		for j := 0; j < numberOfCols; j++ {
			grid[i+1][j+1] = row[j]
		}
	}

	return grid

}

func paperRollCanBeAccessed(grid [][]rune, minimumAccessSpots int, i int, j int, remove bool) bool {

	var availableSpots = 0

	neighbourIndexes := [][]int{

		{-1, -1},
		{-1, 0},
		{-1, 1},

		{0, -1},
		{0, 1},

		{1, -1},
		{1, 0},
		{1, 1},
	}

	for _, coords := range neighbourIndexes {

		newI := coords[0] + i
		newJ := coords[1] + j

		if grid[newI][newJ] == '@' {
			availableSpots += 1
		}
	}

	if remove && availableSpots <= minimumAccessSpots {
		grid[i][j] = '.'
	}

	return availableSpots <= minimumAccessSpots

}

func main() {

	banks := parseFile("04.txt")
	part1 := part1(banks)
	part2 := part2(banks)

	fmt.Println(part1, part2)
}

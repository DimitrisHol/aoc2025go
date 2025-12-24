package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type tile struct {
	x int
	y int
}

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	// testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) int {

	numberOfRows := len(data)
	numberOfCols := len(data[0])

	// Convert the []string into [][]rune to access
	grid := make([][]rune, numberOfRows)
	for i := 0; i < numberOfRows; i++ {

		grid[i] = make([]rune, numberOfCols) // Create the new slice
		row := []rune(data[i])               // Convert row -> []rune

		for j := 0; j < numberOfCols; j++ {
			grid[i][j] = row[j]
		}
	}

	// Step 1 : Find the S position
	startY := 0
	for i := 0; i < numberOfCols; i++ {
		if grid[0][i] == 'S' {
			startY = i
			break
		}
	}

	// Store the places we already gone
	visitedNodes := map[tile]bool{}

	var traverse func(tile)
	traverse = func(startTile tile) {

		if (startTile.x) == numberOfRows-1 {
			return
		}

		nextTile := tile{x: startTile.x + 1, y: startTile.y}

		// Normal tile
		switch grid[nextTile.x][nextTile.y] {
		case '.':
			grid[nextTile.x][nextTile.y] = '|' // DEBUG
			traverse(nextTile)

			// Splitter
		case '^':
			visitedNodes[startTile] = true

			leftTile := tile{x: nextTile.x, y: nextTile.y - 1}
			rightTile := tile{x: nextTile.x, y: nextTile.y + 1}

			if !visitedNodes[leftTile] {
				grid[leftTile.x][leftTile.y] = '|'
				traverse(leftTile)
			}

			if !visitedNodes[rightTile] {
				grid[rightTile.x][rightTile.y] = '|'
				traverse(rightTile)
			}

		}

	}

	traverse(tile{x: 0, y: startY})

	return len(visitedNodes)
}

func part2(data []string) int {

	numberOfRows := len(data)
	numberOfCols := len(data[0])

	// Convert the []string into [][]rune to access
	grid := make([][]rune, numberOfRows)
	for i := 0; i < numberOfRows; i++ {

		grid[i] = make([]rune, numberOfCols) // Create the new slice
		row := []rune(data[i])               // Convert row -> []rune

		for j := 0; j < numberOfCols; j++ {
			grid[i][j] = row[j]
		}
	}

	// Step 1 : Find the S position
	startY := 0
	for i := 0; i < numberOfCols; i++ {
		if grid[0][i] == 'S' {
			startY = i
			break
		}
	}

	cache := map[tile]int{}

	var traverse func(tile) int
	traverse = func(startTile tile) int {

		if (startTile.x) == numberOfRows-1 {
			return 1
		}

		nextTile := tile{x: startTile.x + 1, y: startTile.y}

		// Normal tile
		switch grid[nextTile.x][nextTile.y] {
		case '.':
			return traverse(nextTile)

		case '^':

			leftTile := tile{x: nextTile.x, y: nextTile.y - 1}
			leftTileCount := 0

			if count, exists := cache[leftTile]; exists {
				leftTileCount = count
			} else {
				leftTileCount = traverse(leftTile)
				cache[leftTile] = leftTileCount
			}

			rightTile := tile{x: nextTile.x, y: nextTile.y + 1}
			rightTileCount := 0

			if count, exists := cache[rightTile]; exists {
				rightTileCount = count
			} else {
				rightTileCount = traverse(rightTile)
				cache[rightTile] = rightTileCount
			}

			return leftTileCount + rightTileCount
		}
		return 0
	}

	return traverse(tile{x: 0, y: startY})
}

func main() {

	data := parseFile("07.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

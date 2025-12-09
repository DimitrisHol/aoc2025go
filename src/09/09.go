package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type tile struct {
	y int
	x int
}

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	// testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) int {

	tiles := make([]tile, len(data))

	for index, v := range data {

		coords := strings.Split(v, ",")

		y, _ := strconv.Atoi(coords[0])
		x, _ := strconv.Atoi(coords[1])

		tile := tile{x: x, y: y}
		tiles[index] = tile
	}

	maxArea := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {

			area := calculateArea(tiles[i], tiles[j])
			if area > maxArea {
				maxArea = area
			}
		}
	}

	return maxArea
}

func main() {

	data := parseFile("09.txt")
	fmt.Println(part1(data))
}

func calculateArea(p1 tile, p2 tile) int {
	return (absDiffInt(p1.x, p2.x) + 1) * (absDiffInt(p1.y, p2.y) + 1)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

package main

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	// testPath := filepath.Join(currentPath, "..", "input", "test", filename)
	testPath := filepath.Join(currentPath, "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func calculateRotation(rotation string) int {
	switch rotation {
	case "R":
		return 0
	case "L":
		return 1
	default:
		print("HELP", rotation)
		os.Exit(1)
		return -1
	}
}

func main() {

	rotations := parseFile("01.txt")

	var position int = 50
	timesAtZero := 0

	for _, entry := range rotations {

		rot := calculateRotation(entry[0:1])
		distance, _ := strconv.Atoi(entry[1:])

		// 0 -> Right
		if rot == 0 {
			position = (position + distance) % 100
		} else {
			newPosition := position - (distance % 100)
			if newPosition < 0 {
				position = 100 + newPosition
			} else {
				position = newPosition
			}
		}

		if position == 0 {
			timesAtZero += 1
		}
	}
	print(timesAtZero)

}

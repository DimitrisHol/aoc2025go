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

func part1(rotations []string) int {

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

	return timesAtZero
}

func part2(rotations []string) int {

	var position int = 50
	timesPassedZero := 0

	for _, entry := range rotations {

		rot := calculateRotation(entry[0:1])
		distance, _ := strconv.Atoi(entry[1:])

		// 0 -> Right
		if rot == 0 {

			// 40 + 60  = 100 -> 1 time on top
			// 40 + 65  = 105 -> 1 time across
			// 40 + 160 = 200 -> 1 time across, 1 time on top
			// 40 + 165 = 205 -> 2 time across
			// 40 + 165 = 205 -> 2 time across, 1 time on top

			timesPassedZero += (position + distance) / 100
			position = (position + distance) % 100

		} else {
			// 40 - 40   =    0      -> 1 time on top
			// 40 - 60   = - 20 (80) -> 1 time across
			// 40 - 140  = -120      -> 1 time across, 1 time on top
			// 40 - 160  = -220      -> 2 time across
			// 40 - 240  =    0      -> 2 time across, 1 time on top
			//  0 -  50  =   50      -> 1 time across

			// Step 1 : Do the subtraction normally 40 - 60 = -20
			// Step 2 : while result < 0 -> + 100, count how many times you did that.
			// Edge case 1 : result = 0 -> simply add 1
			// Edge case 2 : we start from 0, subtract one :(

			// Exit early
			if position-distance == 0 {
				timesPassedZero += 1
				position = 0
				// short-circuit, simply adding one
				continue
			}

			var timesPassed = 0
			if position == 0 {
				timesPassed -= 1
			}

			result := position - distance
			for result < 0 {
				result += 100
				timesPassed += 1
			}

			timesPassedZero += timesPassed
			if result == 0 {
				timesPassedZero += 1
			}
			position = result

		}
	}

	return timesPassedZero
}

func main() {

	rotations := parseFile("01.txt")
	fmt.Println(part1(rotations))
	fmt.Println(part2(rotations))
}

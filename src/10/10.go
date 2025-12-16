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
	testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	// testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) int {

	totalNumberPresses := 0

	for _, v := range data {

		factory := strings.Split(v, " ")

		indicatorLight := factory[0][1 : len(factory[0])-1] // e.g. .##. (this is the goal)
		wiringSchematics := factory[1 : len(factory)-1]     // e.g. [(3), (1,3), (2)] etc.
		// joltageRequirements := factory[len(factory)-1:][0]  // e.g. {3,5,4,7}

		// fmt.Println(factory, indicatorLight, wiringSchematics, joltageRequirements)
		fmt.Println(factory, indicatorLight, wiringSchematics)

		totalNumberPresses += calculatePresses(indicatorLight, wiringSchematics)

	}

	return totalNumberPresses
}

func calculatePresses(goal string, permutations []string) int {

	// Indicator light -> . = off, # = on

	// Start position with all off
	current := make([]string, len(goal))
	for i := 0; i < len(goal); i++ {
		current[i] = "."
	}

	goalState := strings.Split(goal, "")

	// priorityQueue := make([]string, 0)

	for strings.Join(current, "") != goal {

		for _, perm := range permutations {
			currentCopy := make([]string, len(current))
			copy(currentCopy, current)

			// Apply the changes
			change := strings.Split(strings.Trim(perm, "()"), ",")

			for _, t := range change {

				index, _ := strconv.Atoi(t)
				toggleLight(currentCopy, index)
			}

			score := calculateScore(goalState, currentCopy)
			fmt.Println(current, currentCopy, change, goalState, score)
		}

		os.Exit(1)
	}

	return 0
}

func toggleLight(state []string, position int) {

	if state[position] == "." {
		state[position] = "#"
	} else {
		state[position] = "."
	}

}

func calculateScore(goal []string, currentState []string) int {

	score := 0

	for i := 0; i < len(goal); i++ {
		if goal[i] == currentState[i] {
			score++
		}
	}

	return score
}

func main() {

	data := parseFile("10.txt")
	fmt.Println(part1(data))
}

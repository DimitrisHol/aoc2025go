package main

import (
	"fmt"
	"math"
	"math/bits"
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
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) int {

	totalNumberPresses := 0

	for _, v := range data {

		factory := strings.Split(v, " ")

		indicatorLight := factory[0][1 : len(factory[0])-1] // e.g. .##. (this is the goal)
		wiringSchematics := factory[1 : len(factory)-1]     // e.g. [(3), (1,3), (2)] etc.
		joltageRequirements := factory[len(factory)-1:][0]  // e.g. {3,5,4,7}

		fmt.Println(indicatorLight, wiringSchematics, joltageRequirements)

		result := calculatePressesVol2(indicatorLight, wiringSchematics)
		// fmt.Printf("result: %v\n", result)

		totalNumberPresses += result

	}

	return totalNumberPresses
}

func calculatePressesVol2(goal string, permutations []string) int {

	// Step 1 : Convert the goal to bits
	goalBinary := goalStringToBinary(goal)
	var maxBinarySize = len(goal) // Get the string length, since the binary value can be cut-off
	// fmt.Printf("goal %v, goalB: %b, size %v\n", goal, goalBinary, len(goal))

	// Step 2 : Convert the permutations to bits
	permutationsBinary := make([]int, len(permutations))
	for i := 0; i < len(permutations); i++ {
		permutationsBinary[i] = permutationToBinary(permutations[i], maxBinarySize)
	}

	// Step 3 : Which combinations of "buttons" lead to the goal. And find the shortest one
	numberOfPermutations := math.Pow(2, float64(len(permutations))) - 1

	minEnables := int(numberOfPermutations)
	// Step 4 : Loop through all possible combinations of the permutations to see if they lead to the goal
	for i := 0; i <= int(numberOfPermutations); i++ {

		// The index in binary represents which permutations to enable
		permutationResult := calculatePermutationResult(permutationsBinary, i)
		numberOfEnables := bits.OnesCount(uint(i))

		if permutationResult == goalBinary {
			minEnables = min(minEnables, numberOfEnables)
		}
	}

	return minEnables
}

func goalStringToBinary(s string) int {

	result := 0

	for i := 0; i < len(s); i++ {
		result <<= 1 // shift left (effectively raise to power of 2)
		if s[i] == '#' {
			result |= 1 // append 1 when the rune is #, which is 1.
		}
	}

	return result
}

func permutationToBinary(perm string, size int) int {

	result := 0
	change := strings.Split(strings.Trim(perm, "()"), ",")

	for i := 0; i < len(change); i++ {
		index, _ := strconv.Atoi(change[i])
		result |= (1 << (size - index - 1))
	}

	return result
}

func calculatePermutationResult(permutations []int, index int) int {

	result := 0

	for i := 0; i < len(permutations); i++ {

		// Check if the permutation should be included : binary is magic
		if (index>>i)&1 == 1 {
			result ^= permutations[i]
		}
	}

	return result
}

func main() {

	data := parseFile("10.txt")
	fmt.Println(part1(data))
}

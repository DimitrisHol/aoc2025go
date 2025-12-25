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
	testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	// testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func solve(data []string) (int, int) {

	totalNumberPresses := 0
	part2 := 0

	for _, v := range data {

		factory := strings.Split(v, " ")

		// indicatorLight := factory[0][1 : len(factory[0])-1] // e.g. .##. (this is the goal)
		wiringSchematics := factory[1 : len(factory)-1]    // e.g. [(3), (1,3), (2)] etc.
		joltageRequirements := factory[len(factory)-1:][0] // e.g. {3,5,4,7}

		fmt.Println(wiringSchematics, joltageRequirements)

		// result := calculatePressesVol2(indicatorLight, wiringSchematics)
		// part2 += BFS(wiringSchematics, joltageRequirements)    // TAKES TOO LONG
		// part2 += solver(wiringSchematics, joltageRequirements) // THERE IS NO SOLVER AVAILABLE IN GO
		part2 += dfs(wiringSchematics, joltageRequirements)

		// totalNumberPresses += result

	}

	return totalNumberPresses, part2
}

func dfs(buttons []string, goals string) int {

	return 0
}

func solver(buttons []string, goals string) int {

	/*

		INPUT : (3) (1,3) (2) (2,3) (0,2) (0,1) ->  {3,5,4,7}

		We want to find the number of presses for each button (a,b,c,d,e,f are the number of presses)
		a * (3) + b * (1,3) + c * (2) + d * (2,3) + e * (0,2) + f * (0,1) = {3,5,4,7}

		MATRIX :

		- Each row is the goal
		- Each column is what the button affects (0/1) : each button press increments the counter by 1

		[0 0 0 0 1 1 = 3]
		[0 1 0 1 0 1 = 5]
		[0 0 1 1 1 0 = 4]
		[1 1 0 0 0 0 = 7]
	*/

	// Split the goal into strings
	stringGoal := strings.Split(strings.Trim(goals, "{}"), ",")

	numCounters := len(stringGoal) // Rows
	numButtons := len(buttons)     // Columns

	// Convert the goals into float64 array
	goalFloats := make([]float64, numCounters)
	for i := 0; i < len(stringGoal); i++ {
		n, _ := strconv.Atoi(stringGoal[i])
		goalFloats[i] = float64(n)
	}

	// Fill The matrix data :
	matrixData := make([]float64, numCounters*numButtons)

	for i := 0; i < len(buttons); i++ {

		buttonPress := strings.Split(strings.Trim(buttons[i], "()"), ",")

		for j := 0; j < len(buttonPress); j++ {
			index, _ := strconv.Atoi(buttonPress[j])

			// Now we finally need to map it to the correct position
			// For example : (1,3) needs
			// 2nd column (it's the second button) : i
			// 1 -> 2nd row, 2nd column -> flatIndex = (numberOfRows * rowIndex) + columnIndex
			// 3 -> 4th row, 2nd column

			flatIndex := (index * numButtons) + i
			matrixData[flatIndex] = 1.0
		}
	}

	// A := mat.NewDense(numCounters, numButtons, matrixData)
	// B := mat.NewVecDense(numCounters, goalFloats)

	// there is NO SOLVER THAT WORKS FFS

	return 0
}

type ButtonsState struct {
	currentState []int
	presses      int
}

func BFS(buttons []string, goals string) int {

	// Step 1 : Parse the goal into an integer array
	stringGoal := strings.Split(strings.Trim(goals, "{}"), ",")
	intGoal := make([]int, len(stringGoal))

	for i := 0; i < len(stringGoal); i++ {
		n, _ := strconv.Atoi(stringGoal[i])
		intGoal[i] = n
	}

	// Set the starting state of the buttons all counters from 0
	state := make([]int, len(stringGoal))
	for i := 0; i < len(stringGoal); i++ {
		state[i] = 0
	}

	// Now the actual calculation : Press the buttons until we reach the goal
	buttonPresses := 0 // How many buttons have been pressed

	// test := map[ButtonsState]bool{}
	queue := []ButtonsState{{currentState: state, presses: 0}}

	visitedStates := map[string]bool{}

	for len(queue) > 0 {

		fmt.Printf("len(queue): %v\n", len(queue))

		// Step 1 : Deque :
		state := queue[0]
		queue = queue[1:]

		if reachedGoal(state.currentState, intGoal) {
			return state.presses
		}

		for i := 0; i < len(buttons); i++ {

			newState := pressButton(buttons[i], state.currentState)
			newButtonState := ButtonsState{newState, state.presses + 1}

			stateKey := stateToString(newState)

			if validState(newState, intGoal) && !visitedStates[stateKey] {
				queue = append(queue, newButtonState)
				visitedStates[stateKey] = true
			}
		}

	}

	return buttonPresses
}

func stateToString(state []int) string {

	return fmt.Sprintf("%v", state)
}

func reachedGoal(state []int, goal []int) bool {

	for i := 0; i < len(state); i++ {
		if state[i] != goal[i] {
			return false
		}
	}
	return true
}

func validState(state []int, goal []int) bool {

	for i := 0; i < len(state); i++ {
		if state[i] > goal[i] { // Not overshooting the goal
			return false
		}
	}
	return true
}

func pressButton(button string, state []int) []int {

	stateCopy := make([]int, len(state))
	copy(stateCopy, state)

	change := strings.Split(strings.Trim(button, "()"), ",")
	for i := 0; i < len(change); i++ {
		index, _ := strconv.Atoi(change[i])
		stateCopy[index]++
	}
	return stateCopy
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
	fmt.Println(solve(data))
}

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

func part1(data []string) int {

	var adjacencyList = map[string][]string{}

	for i := 0; i < len(data); i++ {

		nodeAndChildren := strings.Split(data[i], ":")
		adjacencyList[nodeAndChildren[0]] = strings.Fields(nodeAndChildren[1])

	}

	start := "svr"
	goal := "out"

	stack := [][]string{{start}}

	counter := 0

	for len(stack) > 0 {

		// Get last object, update the stack
		path := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		// Get the last node of the path
		lastNode := path[len(path)-1]
		if lastNode == goal {
			fmt.Printf("path: %v\n", path)
			counter += 1
		}

		for _, nextNode := range adjacencyList[lastNode] {

			if pathContainsNode(path, nextNode) {
				continue
			}

			newPath := append([]string{}, path...)
			newPath = append(newPath, nextNode)

			stack = append(stack, newPath)
		}

	}

	return counter
}

func push(stack []string, value string) []string {
	return append(stack, value)
}

func pop(stack []string) (string, []string) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

func pathContainsNode(path []string, node string) bool {

	for i := 0; i < len(path); i++ {
		if path[i] == node {
			return true
		}
	}
	return false
}

func part1rec(data []string) int {

	var adjacencyList = map[string][]string{}

	for i := 0; i < len(data); i++ {

		nodeAndChildren := strings.Split(data[i], ":")
		adjacencyList[nodeAndChildren[0]] = strings.Fields(nodeAndChildren[1])
	}

	start := "you"
	goal := "out"

	return depthFirstSearchAllSolutions(adjacencyList, start, goal)
}

func part2rec(data []string) int {

	var adjacencyList = map[string][]string{}

	for i := 0; i < len(data); i++ {

		nodeAndChildren := strings.Split(data[i], ":")
		adjacencyList[nodeAndChildren[0]] = strings.Fields(nodeAndChildren[1])
	}

	start := "svr"
	goal := "out"

	return depthFirstSearch(adjacencyList, start, goal)
}

func depthFirstSearchAllSolutions(graph map[string][]string, start string, goal string) int {

	countToGoal := 0

	var dfs func(string) // CLOSURE!
	dfs = func(currentNode string) {

		if currentNode == goal {
			countToGoal++
		}

		for _, nextNode := range graph[currentNode] {
			dfs(nextNode)
		}
	}

	dfs(start)
	return countToGoal
}

func depthFirstSearch(graph map[string][]string, start string, goal string) int {

	reqFFT := "fft"
	reqDAC := "dac"

	memo := map[string]int{}

	var dfs func(string, bool, bool) int // CLOSURE!
	dfs = func(currentNode string, fft bool, dac bool) int {

		// Memoization ! (caching)
		cacheKey := fmt.Sprintf("%s-%t-%t", currentNode, fft, dac)
		if count, exists := memo[cacheKey]; exists {
			return count
		}

		if currentNode == reqFFT {
			fft = true
		}

		if currentNode == reqDAC {
			dac = true
		}

		// We reached the goal !
		if currentNode == goal {
			if fft && dac {
				return 1
			}
			return 0
		}

		count := 0
		// Sum up all the results from all recursive calls to children
		for _, nextNode := range graph[currentNode] {
			count += dfs(nextNode, fft, dac)
		}
		memo[cacheKey] = count

		return count
	}

	return dfs(start, false, false)
}

func main() {

	data := parseFile("11.txt")
	// data2 := parseFile("11p2.txt")

	// PROD
	data2 := parseFile("11.txt")

	fmt.Println(part1rec(data))
	fmt.Println(part2rec(data2))
}

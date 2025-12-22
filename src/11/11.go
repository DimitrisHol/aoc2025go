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

	start := "you"
	goal := "out"

	stack := [][]string{{start}}

	counter := 0

	// visitedNodes := map[string]bool{}
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

func main() {

	data := parseFile("11.txt")
	fmt.Println(part1(data))
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

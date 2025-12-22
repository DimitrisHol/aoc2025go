package main

import (
	"fmt"
	"os"
	"path/filepath"
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

	var adjacencyList = map[string][]string{}

	for i := 0; i < len(data); i++ {

		nodeAndChildren := strings.Split(data[i], ":")
		adjacencyList[nodeAndChildren[0]] = strings.Fields(nodeAndChildren[1])

	}

	// Start from "A", with goal : "E"
	stack := []string{"A"}
	goal := "E"

	visitedNodes := map[string]bool{}
	for len(stack) > 0 {

		// Get last object, update the stack
		top, newStack := pop(stack)
		stack = newStack

		// Check if we reached the destination
		if goal == top {
			fmt.Println("Goal reached", top)
			break
		}

		fmt.Println("Top :", top, "newstack", newStack)
		if visitedNodes[top] {
			continue
		}
		visitedNodes[top] = true

		for _, nextNode := range adjacencyList[top] {
			fmt.Printf("nextNode: %v\n", nextNode)
			stack = push(stack, nextNode)
		}

	}

	return 0
}

func main() {

	data := parseFile("112.txt")
	fmt.Println(part1(data))
}

func push(stack []string, value string) []string {
	return append(stack, value)
}

func pop(stack []string) (string, []string) {
	return stack[len(stack)-1], stack[:len(stack)-1]
}

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	testPath := filepath.Join(currentPath, filename)
	// testPath := filepath.Join(currentPath, "..", "..", "input", file`name)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) int {

	var adjacencyList = map[string][]string{}

	for i := 0; i < len(data); i++ {

		nodeAndChildren := strings.Split(data[i], ":")
		adjacencyList[nodeAndChildren[0]] = strings.Fields(nodeAndChildren[1])

	}

	start := "A"
	goal := "E"

	depthFirstSearch(adjacencyList, start, goal)

	return 0
}

func depthFirstSearch(graph map[string][]string, start string, goal string) bool {

	// Inner function has access to the visitedNodes map!
	visitedNodes := map[string]bool{}

	var dfs func(string) bool // CLOSURE!
	dfs = func(currentNode string) bool {

		visitedNodes[currentNode] = true
		fmt.Println("Currently in node ", currentNode)

		if currentNode == goal {
			fmt.Println("Found it !")
			return true
		}

		for _, nextNode := range graph[currentNode] {

			fmt.Printf("nextNode: %v\n", nextNode)

			if !visitedNodes[nextNode] {
				if dfs(nextNode) {
					return true
				}
			}
		}

		return false
	}

	return dfs(start)
}

func main() {

	data := parseFile("graphExample.txt")
	fmt.Println(part1(data))
}

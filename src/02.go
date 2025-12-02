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
	return strings.Split(string(data), ",")
}


func part1(ranges []string) int {

	for _, v := range v {
		
	}

	return 0
	
}

func part2(ranges []string) int {

	return 0
}

func main() {

	ranges := parseFile("02.txt")
	fmt.Println(part1(ranges))
	fmt.Println(part2(ranges))
}


12 12 13

123123
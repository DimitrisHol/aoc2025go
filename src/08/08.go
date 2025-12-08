package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type position struct {
	x int
	y int
	z int
}

type distance struct {
	pos1     position
	pos2     position
	distance float64
}

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	// testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) {

	positions := make([]position, len(data))

	for index, v := range data {

		coords := strings.Split(v, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		pos := position{x: x, y: y, z: z}
		positions[index] = pos

	}

	// for 20 elements, we need 190 distances n * (n-1) / 2
	distancesLength := len(positions) * (len(positions) - 1) / 2
	distances := make([]distance, distancesLength)

	distanceIndex := 0

	for i := 0; i < len(positions); i++ {

		for j := i + 1; j < len(positions); j++ {

			pos1 := positions[i]
			pos2 := positions[j]
			d := calculateDistance(positions[i], positions[j])

			dist := distance{pos1: pos1, pos2: pos2, distance: d}
			distances[distanceIndex] = dist
			distanceIndex += 1
		}
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].distance < distances[j].distance
	})

	fmt.Println(len(distances))

}

func main() {

	data := parseFile("08.txt")
	part1(data)
}

func calculateDistance(box1 position, box2 position) float64 {

	dx := box1.x - box2.x
	dy := box1.y - box2.y
	dz := box1.z - box2.z

	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

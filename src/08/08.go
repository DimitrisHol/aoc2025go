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
	index int
	x     int
	y     int
	z     int
}

type distance struct {
	pos1     position
	pos2     position
	distance float64
}

type DisjoinSet struct {
	parent   map[int]int
	elements map[int]position
}

// Create a new set (a new graph)
func (ds DisjoinSet) makeSet(pos position) DisjoinSet {

	if _, exists := ds.parent[pos.index]; !exists {
		ds.parent[pos.index] = pos.index // the first element is its own parent
		ds.elements[pos.index] = pos
	}
	return ds
}

// Find the representative of the index (of the position)
func (ds DisjoinSet) find(index int) int {

	// We've reached the base case (element is it's own parent)
	if ds.parent[index] == index {
		return index
	}

	// TODO : OPTIMIZATION you can flatten the tree as you go
	// E -> D -> C -> B -> A
	// so the next time you find(D) go directly to A instead of moving one step

	return ds.find(ds.parent[index])
}

func (ds DisjoinSet) Union(x int, y int) DisjoinSet {

	rootX := ds.find(x)
	rootY := ds.find(y)

	if rootX != rootY {
		ds.parent[rootX] = rootY
	}
	return ds
}

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	// testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) int {

	positions := make([]position, len(data))

	for index, v := range data {

		coords := strings.Split(v, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		pos := position{index: index, x: x, y: y, z: z}
		positions[index] = pos

	}

	// for 20 elements, we need 190 distances n * (n-1) / 2
	// for 1000 elements we need 499.500 distances
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

	ds := DisjoinSet{parent: make(map[int]int), elements: make(map[int]position)}

	// First create sets with all the individual elements
	for i := 0; i < len(positions); i++ {
		ds.makeSet(positions[i])
	}

	// Union the first 1000 closest boxes
	for i := 0; i < 1000; i++ { // 10 for the example

		pos1 := distances[i].pos1
		pos2 := distances[i].pos2

		ds.Union(pos1.index, pos2.index)
	}

	circuitSizes := make(map[int]int)

	// Loop through the positions, and find the representative for each position (recursively)
	// and simply count how many times each representative is present
	// representative = group leader, which means they are part of that group !
	for i := 0; i < len(positions); i++ {
		representative := ds.find(i)
		circuitSizes[representative] += 1
	}

	// Transform the map into list, and sort it. We only need the first 3 elements
	var sizes []int
	for _, size := range circuitSizes {
		sizes = append(sizes, size)
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	product := sizes[0] * sizes[1] * sizes[2]
	return product
}

func part2(data []string) int {

	positions := make([]position, len(data))

	for index, v := range data {

		coords := strings.Split(v, ",")

		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		z, _ := strconv.Atoi(coords[2])

		pos := position{index: index, x: x, y: y, z: z}
		positions[index] = pos

	}

	// for 20 elements, we need 190 distances n * (n-1) / 2
	// for 1000 elements we need 499.500 distances
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

	ds := DisjoinSet{parent: make(map[int]int), elements: make(map[int]position)}

	// First create sets with all the individual elements
	for i := 0; i < len(positions); i++ {
		ds.makeSet(positions[i])
	}

	// Union until all representatives are the same
	// which means that we're all under the same group
	for i := 0; i < len(distances); i++ {

		pos1 := distances[i].pos1
		pos2 := distances[i].pos2

		ds.Union(pos1.index, pos2.index)

		representative := ds.find(0)
		var allRepresentativesSame bool = true

		for j := 1; j < len(positions); j++ {

			if representative != ds.find(j) {
				allRepresentativesSame = false
				break
			}
		}

		if allRepresentativesSame {
			return pos1.x * pos2.x
		}
	}

	fmt.Println("Something went wrong")
	return 0
}

func main() {

	data := parseFile("08.txt")
	fmt.Println(part1(data))
	fmt.Println(part2(data))
}

func calculateDistance(box1 position, box2 position) float64 {

	dx := box1.x - box2.x
	dy := box1.y - box2.y
	dz := box1.z - box2.z

	return math.Sqrt(float64(dx*dx + dy*dy + dz*dz))
}

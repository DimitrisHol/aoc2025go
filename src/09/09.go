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

type tile struct {
	x int
	y int
}

type rectangle struct {
	a    tile
	b    tile
	area int
}

func parseFile(filename string) []string {
	currentPath, _ := os.Getwd()
	// testPath := filepath.Join(currentPath, "..", "..", "input", "test", filename)
	testPath := filepath.Join(currentPath, "..", "..", "input", filename)

	data, _ := os.ReadFile(testPath)
	return strings.Split(string(data), "\r\n")
}

func part1(data []string) (rectangle, rectangle) {

	minX := math.MaxInt
	minY := math.MaxInt

	maxX := math.MinInt
	maxY := math.MinInt

	tiles := make([]tile, len(data))

	for index, v := range data {

		coords := strings.Split(v, ",")

		x, _ := strconv.Atoi(coords[1])
		y, _ := strconv.Atoi(coords[0])

		minX = min(minX, x)
		minY = min(minY, y)

		maxX = max(maxX, x)
		maxY = max(maxY, y)

		tile := tile{x: x, y: y}
		tiles[index] = tile
	}

	// fmt.Println(minX, minY, maxX, maxY)

	// Part 1 :
	candidates := make([]rectangle, len(data)*(len(data)-1)/2)

	counter := 0
	for i := 0; i < len(tiles); i++ {
		for j := i + 1; j < len(tiles); j++ {

			area := calculateArea(tiles[i], tiles[j])
			candidates[counter] = rectangle{a: tiles[i], b: tiles[j], area: area}
			counter++
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].area > candidates[j].area
	})

	maxAreaPart1 := candidates[0]

	// Fill in the green tiles :

	edgeTiles := make(map[tile]bool)
	for i := 0; i < len(tiles); i++ {

		currentX := tiles[i].x
		currentY := tiles[i].y

		// loop arround from last to first
		nextIndex := (i + 1) % len(tiles)

		nextX := tiles[nextIndex].x
		nextY := tiles[nextIndex].y

		// Fill in horizontally
		if currentX == nextX {

			min := min(currentY, nextY)
			max := max(currentY, nextY)

			for y := min; y <= max; y++ {

				edgeTiles[tile{x: currentX, y: y}] = true
			}

			// Fill in vertically
		} else if currentY == nextY {
			min := min(currentX, nextX)
			max := max(currentX, nextX)

			for x := min; x <= max; x++ {
				edgeTiles[tile{x: x, y: currentY}] = true
			}
		} else {
			fmt.Println("Something went wrong during greens")
			os.Exit(1)
		}
	}

	// Loop through all the opposite side red tile rectangles, sorted by max area
	totalCandidates := len(candidates)
	for index, rectangle := range candidates {

		if index%10 == 0 {
			fmt.Printf("Checked %d/%d candidates (%.1f%%)...\n", index, totalCandidates, float64(index)*100/float64(totalCandidates))
		}

		chosenOne := true

		// Loop through all the rectangle edge tiles, to check if they are inside the polygon
		for _, edgeTile := range getRectangleEdges(rectangle) {

			// Easily check if part of the polygon border
			if edgeTiles[edgeTile] {
				continue
			}

			if !isInsidePolygon(edgeTile, tiles) {
				chosenOne = false
				break
			}
		}

		if chosenOne {
			return maxAreaPart1, rectangle
		}

	}

	return maxAreaPart1, rectangle{a: tile{x: 666, y: 666}}
}

func isInsidePolygon(point tile, redTiles []tile) bool {

	// RAY TRACING DEEZ NUTS
	// Count how many times it crosses the polygon edges
	// Even number = outside the polygon
	// Odd number  = inside the polygon

	crossings := 0

	n := len(redTiles)

	for i := 0; i < n; i++ {
		v1 := redTiles[i]
		v2 := redTiles[(i+1)%n]

		// Check if this edge crosses our horizontal ray going right
		// The edge must straddle our point's x coordinate
		if (v1.x <= point.x && v2.x > point.x) || (v2.x <= point.x && v1.x > point.x) {
			// Calculate where the edge intersects our ray's x level
			// Linear interpolation to find y coordinate
			t := float64(point.x-v1.x) / float64(v2.x-v1.x)
			yIntersect := float64(v1.y) + t*float64(v2.y-v1.y)

			// Check if intersection is to the right of our point
			if yIntersect > float64(point.y) {
				crossings++
			}
		}
	}

	return crossings%2 == 1
}

func main() {

	data := parseFile("09.txt")
	fmt.Println(part1(data))
}

func calculateArea(p1 tile, p2 tile) int {
	return (absDiffInt(p1.x, p2.x) + 1) * (absDiffInt(p1.y, p2.y) + 1)
}

func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func getRectangleEdges(r rectangle) []tile {

	minX := min(r.a.x, r.b.x)
	maxX := max(r.a.x, r.b.x)

	minY := min(r.a.y, r.b.y)
	maxY := max(r.a.y, r.b.y)

	edges := []tile{}

	// Top Edge : minX { minY - maxY}
	for i := minY; i <= maxY; i++ {
		edges = append(edges, tile{x: minX, y: i})
	}

	// Bottom Edge : maxX {minY - maxY}
	for i := minY; i <= maxY; i++ {
		edges = append(edges, tile{x: maxX, y: i})
	}

	// Left edge : minY (minX - maxX)
	for i := minX; i <= maxX; i++ {
		edges = append(edges, tile{x: i, y: minY})
	}

	// Right Edge : maxY (minX - maxX)
	for i := minX; i <= maxX; i++ {
		edges = append(edges, tile{x: i, y: maxY})
	}

	return edges
}

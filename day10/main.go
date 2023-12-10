package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	filename := "input.txt"

	start := time.Now()
	part1Result := part1(filename)
	fmt.Printf("Part 1 answer: %d in %v\n", part1Result, time.Since(start))

	start = time.Now()
	part2Result := part2(filename)
	fmt.Printf("Part 2 answer: %d in %v\n", part2Result, time.Since(start))
}

type Coordinate struct {
	X, Y int
}

var directions = map[string][]Coordinate{
	"|": {{0, 1}, {0, -1}},
	"-": {{1, 0}, {-1, 0}},
	"L": {{0, -1}, {1, 0}},
	"J": {{-1, 0}, {0, -1}},
	"7": {{0, 1}, {-1, 0}},
	"F": {{1, 0}, {0, 1}},
	"S": {{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
}

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var diagram [][]string

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		diagram = append(diagram, splitLine)
	}

	startY, startX := findStart(diagram)
	maxDistance, _ := findLoopLengthAndCoordinates(diagram, startX, startY)
	return (maxDistance / 2)
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var diagram [][]string

	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, "")
		diagram = append(diagram, splitLine)
	}

	startY, startX := findStart(diagram)
	_, loopPath := findLoopLengthAndCoordinates(diagram, startX, startY)

	// https://en.wikipedia.org/wiki/Shoelace_formula
	polygonArea := 0
	for i := 0; i < len(loopPath); i++ {
		cur := loopPath[i]
		next := loopPath[(i+1)%len(loopPath)]

		polygonArea += cur.X*next.Y - cur.Y*next.X
	}

	if polygonArea < 0 {
		polygonArea = -polygonArea
	}
	polygonArea /= 2

	// https://en.wikipedia.org/wiki/Pick%27s_theorem
	return polygonArea - len(loopPath)/2 + 1
}

func findStart(diagram [][]string) (int, int) {
	for i, row := range diagram {
		for j, value := range row {
			if value == "S" {
				return i, j
			}
		}
	}

	return -1, -1
}

func findLoopLengthAndCoordinates(diagram [][]string, startX, startY int) (int, []Coordinate) {
	visited := make(map[Coordinate]bool)
	path := []Coordinate{}
	var traverse func(x, y int, prev Coordinate, distance int) int
	traverse = func(x, y int, prev Coordinate, distance int) int {
		currentCoord := Coordinate{x, y}

		if x == startX && y == startY && visited[currentCoord] {
			return distance
		}

		if _, seen := visited[currentCoord]; !seen {
			visited[currentCoord] = true
			path = append(path, currentCoord)
			for _, dir := range directions[diagram[y][x]] {
				nextX, nextY := x+dir.X, y+dir.Y
				if nextX != prev.X || nextY != prev.Y {
					if _, ok := directions[diagram[nextY][nextX]]; ok {
						if steps := traverse(nextX, nextY, currentCoord, distance+1); steps != -1 {
							return steps
						}
					}
				}
			}
		}
		return -1
	}

	loopLength := traverse(startX, startY, Coordinate{-1, -1}, 0)
	return loopLength, path
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
	"slices"
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

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	var matrix [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		split := strings.Split(line, "")

		matrix = append(matrix, split)

	}

	rows := len(matrix)
	cols := len(matrix[0])

	for i := range matrix {
		j := 0
		for j < cols {
			if isDigit(matrix[i][j]) {
				symbolFound := false
				for x := max(0, i-1); x <= min(rows-1, i+1) && !symbolFound; x++ {
					for y := max(0, j-1); y <= min(len(matrix[i])-1, j+1) && !symbolFound; y++ {
						if (x != i || y != j) && isSymbol(matrix[x][y]) {
							symbolFound = true
						}
					}
				}

				if symbolFound {
					number, _, right := findCompleteNumber(matrix, i, j)
					sum += number
					j = right
				}
			}
			j++
		}
	}

	return sum
}

func isDigit(character string) bool {
	if _, err := strconv.Atoi(character); err == nil {
		return true
	} else {
		return false
	}
}

func isSymbol(character string) bool {
	return !isDigit(character) && character != "."
}

func findCompleteNumber(schematic [][]string, i, j int) (int, int, int) {
	number := schematic[i][j]
	left := j - 1
	right := j + 1
	cols := len(schematic[i])

	for left >= 0 {
		if _, err := strconv.Atoi(schematic[i][left]); err == nil {
			number = schematic[i][left] + number
			left--
		} else {
			break
		}
	}

	for right < cols {
		if _, err := strconv.Atoi(schematic[i][right]); err == nil {
			number += schematic[i][right]
			right++
		} else {
			break
		}
	}

	completeNumber, err := strconv.Atoi(number)
	if err != nil {
		log.Fatal(err)
	}
	return completeNumber, left + 1, right - 1
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	var matrix [][]string

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		split := strings.Split(line, "")

		matrix = append(matrix, split)

	}

	rows := len(matrix)
	cols := len(matrix[0])

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == "*" {
				adjacentNumbers := []int{}

				for x := max(0, i-1); x <= min(rows-1, i+1); x++ {
					for y := max(0, j-1); y <= min(cols-1, j+1); y++ {
						if (x != i || y != j) && isDigit(matrix[x][y]) {
							number, _, _ := findCompleteNumber(matrix, x, y)
							if !slices.Contains(adjacentNumbers, number) {
								adjacentNumbers = append(adjacentNumbers, number)
							}
						}
					}
				}

				if len(adjacentNumbers) == 2 {
					sum += adjacentNumbers[0] * adjacentNumbers[1]
				}

			}
		}
	}

	return sum
}

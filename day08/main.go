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

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan()

	nodes := make(map[string][2]string)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		node := parts[0]
		connections := strings.Trim(parts[1], "()")
		connectionParts := strings.Split(connections, ", ")
		nodes[node] = [2]string{connectionParts[0], connectionParts[1]}
	}

	currentNode := "AAA"
	steps := 0
	for currentNode != "ZZZ" {
		direction := instructions[steps%len(instructions)]
		if direction == 'L' {
			currentNode = nodes[currentNode][0]
		} else {
			currentNode = nodes[currentNode][1]
		}
		steps++
	}

	return steps
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	instructions := scanner.Text()

	scanner.Scan()

	nodes := make(map[string][2]string)
	var startingNodes []string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " = ")
		node := parts[0]
		connections := strings.Trim(parts[1], "()")
		connectionParts := strings.Split(connections, ", ")
		nodes[node] = [2]string{connectionParts[0], connectionParts[1]}

		if strings.HasSuffix(node, "A") {
			startingNodes = append(startingNodes, node)
		}
	}

	var stepsSlice []int
	for _, startNode := range startingNodes {
		steps := traverse(nodes, startNode, instructions)
		stepsSlice = append(stepsSlice, steps)
	}
	return lcmSlice(stepsSlice)
}

func traverse(nodes map[string][2]string, startNode, instructions string) int {
	currentNode := startNode
	steps := 0
	for !strings.HasSuffix(currentNode, "Z") {
		direction := instructions[steps%len(instructions)]
		if direction == 'L' {
			currentNode = nodes[currentNode][0]
		} else {
			currentNode = nodes[currentNode][1]
		}
		steps++
	}
	return steps
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func lcmSlice(nums []int) int {
	result := nums[0]
	for _, num := range nums[1:] {
		result = lcm(result, num)
	}
	return result
}

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	filename := "input.txt"

	start := time.Now()
	part1Result := part1(filename)
	fmt.Printf("Part 1 answer: %d in %v\n", part1Result, time.Since(start))

	start = time.Now()
	part2Result := part2(filename, 10)
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

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		splitGame := strings.SplitN(line, ":", 2)
		numbersPart := strings.TrimSpace(splitGame[1])

		splitMyAndWinningNumbers := strings.Split(numbersPart, "|")
		myNumbers := convertToInt(strings.TrimSpace(splitMyAndWinningNumbers[0]))
		winningNumbers := convertToInt(strings.TrimSpace(splitMyAndWinningNumbers[1]))

		myNumbersMap := make(map[int]bool)
		for _, v := range myNumbers {
			myNumbersMap[v] = true
		}

		sameNumbers := 0
		for _, v := range winningNumbers {
			if myNumbersMap[v] {
				sameNumbers++
			}
		}
		if sameNumbers == 0 {
			continue
		}
		sum += max(0, 1<<(sameNumbers-1))

	}

	return sum
}

func convertToInt(string string) []int {
	var slice []int
	for _, numString := range strings.Fields(string) {
		num, err := strconv.Atoi(numString)
		if err != nil {
			log.Fatal(err)
		}
		slice = append(slice, num)
	}

	return slice
}

func part2(filename string, splitIndex int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var cards [][]int

	for scanner.Scan() {
		line := scanner.Text()
		if err != nil {
			log.Fatal(err)
		}

		splitGame := strings.SplitN(line, ":", 2)
		numbersPart := strings.TrimSpace(splitGame[1])

		splitMyAndWinningNumbers := strings.Split(numbersPart, "|")
		myNumbers := convertToInt(strings.TrimSpace(splitMyAndWinningNumbers[0]))
		winningNumbers := convertToInt(strings.TrimSpace(splitMyAndWinningNumbers[1]))

		cards = append(cards, append(myNumbers, winningNumbers...))
	}

	cardCopies := make([]int, len(cards))
	for i := range cardCopies {
		cardCopies[i] = 1
	}

	for i := 0; i < len(cards); i++ {
		myNumbersMap := make(map[int]bool)
		for _, v := range cards[i][:splitIndex] {
			myNumbersMap[v] = true
		}

		sameNumbers := 0
		for _, v := range cards[i][splitIndex:] {
			if myNumbersMap[v] {
				sameNumbers++
			}
		}

		for j := 1; j <= sameNumbers && i+j < len(cards); j++ {
			cardCopies[i+j] += cardCopies[i]
		}
	}

	totalCards := 0
	for _, copies := range cardCopies {
		totalCards += copies
	}

	return totalCards

}

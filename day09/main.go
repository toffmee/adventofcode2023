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

	var histories [][]int
	
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		history := stringToInt(splitLine)
		histories = append(histories, history)
	}

	return sumExtrapolatedValues(histories)
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var histories [][]int
	
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, " ")
		history := stringToInt(splitLine)
		slices.Reverse(history)
		histories = append(histories, history)
	}

	return sumExtrapolatedValues(histories)
}

func stringToInt(strings []string) []int {
	intSlice := make([]int, 0)
	for _, str := range strings {
		num, err := strconv.Atoi(str)
		if err != nil {
			log.Fatal(err)
		}

		intSlice = append(intSlice, num)
	}
	
	return intSlice
}

func calculateDifferences(sequence []int) []int {
    differences := make([]int, len(sequence)-1)
    for i := 0; i < len(sequence)-1; i++ {
        differences[i] = sequence[i+1] - sequence[i]
    }
    return differences
}

func predictNextValue(sequence []int) int {
    var sequences [][]int
    sequences = append(sequences, sequence)

    for {
        lastSequence := sequences[len(sequences)-1]
        differences := calculateDifferences(lastSequence)
        sequences = append(sequences, differences)
        if allZeros(differences) {
            break
        }
    }

    for i := len(sequences) - 2; i >= 0; i-- {
        lastValue := sequences[i][len(sequences[i])-1]
        nextDiff := sequences[i+1][len(sequences[i+1])-1]
        sequences[i] = append(sequences[i], lastValue+nextDiff)
    }

    return sequences[0][len(sequences[0])-1]
}

func allZeros(sequence []int) bool {
    for _, v := range sequence {
        if v != 0 {
            return false
        }
    }
    return true
}

func sumExtrapolatedValues(sequences [][]int) int {
    sum := 0
    for _, sequence := range sequences {
        sum += predictNextValue(sequence)
    }
    return sum
}

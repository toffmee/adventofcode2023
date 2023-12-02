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
	part1Result := part1(filename, 12, 13, 14)
	fmt.Printf("Part 1 answer: %d in %v\n", part1Result, time.Since(start))

	start = time.Now()
	part2Result := part2(filename)
	fmt.Printf("Part 2 answer: %d in %v\n", part2Result, time.Since(start))
}

func part1(filename string, bagReds int, bagGreens int, bagBlues int) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		gameParts := strings.Split(line, ": ")
		gameNumberStr := strings.TrimPrefix(gameParts[0], "Game ")
		gameNumber, err := strconv.Atoi(gameNumberStr)
		if err != nil {
			log.Fatal(err)
		}

		setsData := strings.Split(gameParts[1], "; ")

		gamePossible := true

		for _, set := range setsData {
			reds, greens, blues := 0, 0, 0

			colors := strings.Split(set, ", ")

			for _, color := range colors {
				colorParts := strings.Split(color, " ")

				count, err := strconv.Atoi(colorParts[0])
				if err != nil {
					log.Fatal(err)
					continue
				}

				switch colorParts[1] {
				case "red":
					reds += count
				case "green":
					greens += count
				case "blue":
					blues += count
				}
			}

			if reds > bagReds || greens > bagGreens || blues > bagBlues {
				gamePossible = false
			}

		}

		if gamePossible {
			sum += gameNumber
		}

	}

	return sum
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		gameParts := strings.Split(line, ": ")
		if err != nil {
			log.Fatal(err)
		}

		setsData := strings.Split(gameParts[1], "; ")

		gameMinReds, gameMinGreens, gameMinBlues := 0, 0, 0

		for _, set := range setsData {

			colors := strings.Split(set, ", ")

			for _, color := range colors {
				colorParts := strings.Split(color, " ")

				count, err := strconv.Atoi(colorParts[0])
				if err != nil {
					log.Fatal(err)
					continue
				}

				switch colorParts[1] {
				case "red":
					if count > gameMinReds {
						gameMinReds = count
					}
				case "green":
					if count > gameMinGreens {
						gameMinGreens = count
					}
				case "blue":
					if count > gameMinBlues {
						gameMinBlues = count
					}
				}
			}

		}

		sum += gameMinReds * gameMinGreens * gameMinBlues

	}

	return sum

}

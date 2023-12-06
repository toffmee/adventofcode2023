package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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
	part2Result := part2(filename)
	fmt.Printf("Part 2 answer: %d in %v\n", part2Result, time.Since(start))
}

type Race struct {
	time, distance int
}

type RaceSlice []Race

func (r *RaceSlice) Insert(time, distance int) {
	*r = append(*r, Race{time, distance})
}

func calculateWaysToWin(time, record int) int {
	a, b, c := -1, time, -record
	discriminant := b*b - 4*a*c

	sqrtDisc := math.Sqrt(float64(discriminant))
	root1 := (-float64(b) + sqrtDisc) / (2 * float64(a))
	root2 := (-float64(b) - sqrtDisc) / (2 * float64(a))

	start := int(math.Ceil(root1))
	end := int(math.Floor(root2))

	ways := 0
	for i := start; i <= end; i++ {
		if i*(time-i) > record {
			ways++
		}
	}

	return ways
}

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	races := parseDocument(lines)
	totalWays := 1
	for _, race := range races {
		ways := calculateWaysToWin(race.time, race.distance)
		totalWays *= ways
	}
	return totalWays

}

func parseDocument(lines []string) RaceSlice {
	timeLine := strings.TrimPrefix(lines[0], "Time:")
	distanceLine := strings.TrimPrefix(lines[1], "Distance:")
	timeStrings := strings.Fields(timeLine)
	distanceStrings := strings.Fields(distanceLine)
	var races RaceSlice
	for i := 0; i < len(timeStrings); i++ {
		time, err := strconv.Atoi(timeStrings[i])
		if err != nil {
			log.Fatalln(err)
		}
		distance, err := strconv.Atoi(distanceStrings[i])
		if err != nil {
			log.Fatalln(err)
		}
		races.Insert(time, distance)
	}

	return races
}

func parseDocumentKerning(lines []string) Race {
	timeLine := strings.ReplaceAll(strings.TrimPrefix(lines[0], "Time:"), " ", "")
	distanceLine := strings.ReplaceAll(strings.TrimPrefix(lines[1], "Distance:"), " ", "")

	time, err := strconv.Atoi(timeLine)
	if err != nil {
		log.Fatalln(err)
	}
	distance, err := strconv.Atoi(distanceLine)
	if err != nil {
		log.Fatalln(err)
	}
	return Race{time: time, distance: distance}

}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	race := parseDocumentKerning(lines)
	ways := calculateWaysToWin(race.time, race.distance)
	return ways
}

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

type Range struct {
	destStart, sourceStart, rangeLength int
}

type RangeSlice []Range

func (r *RangeSlice) Insert(destStart, sourceStart, rangeLength int) {
	*r = append(*r, Range{destStart, sourceStart, rangeLength})
}

func (r *RangeSlice) Find(s int) int {
	for _, rg := range *r {
		if s >= rg.sourceStart && s < rg.sourceStart+rg.rangeLength {
			return rg.destStart + (s - rg.sourceStart)
		}
	}
	return s
}

func part1(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seeds := parseSeeds(scanner.Text())

	mappings := make([]RangeSlice, 7)
	for i := range mappings {
		scanner.Scan()
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}

			if !strings.Contains(line, "map:") {
				parseMapping(line, &mappings[i])
			}
		}
	}

	lowestLocation := math.MaxInt
	for _, seed := range seeds {
		location := seed
		for _, mapping := range mappings {
			location = mapping.Find(location)
		}
		if location < lowestLocation {
			lowestLocation = location
		}
	}

	return lowestLocation
}

func parseSeeds(seedsLine string) []int {
	parts := strings.Fields(seedsLine)[1:]
	seeds := make([]int, len(parts))
	for i, part := range parts {
		seeds[i], _ = strconv.Atoi(part)
	}
	return seeds
}

func parseMapping(line string, r *RangeSlice) {
	parts := strings.Fields(line)
	destStart, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	sourceStart, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	rangeLength, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal(err)
	}
	r.Insert(destStart, sourceStart, rangeLength)
}

func part2(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	seedRanges := parseSeeds(scanner.Text())

	mappings := make([]RangeSlice, 7)
	for i := range mappings {
		scanner.Scan()
		for scanner.Scan() {
			line := scanner.Text()
			if line == "" {
				break
			}
			if !strings.Contains(line, "map:") {
				parseMapping(line, &mappings[i])
			}
		}
	}

	lowestLocation := math.MaxInt
	for i := 0; i < len(seedRanges); i += 2 {
		start, length := seedRanges[i], seedRanges[i+1]
		for seed := start; seed < start+length; seed++ {
			location := seed
			for _, mapping := range mappings {
				location = mapping.Find(location)
			}
			if location < lowestLocation {
				lowestLocation = location
			}
		}
	}

	return lowestLocation
}

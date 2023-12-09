package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

type HandAndBid struct {
	hand     string
	handType int
	bid      int
}

func part1(filename string) int {
	hands := readInput(filename)
	return calculateWinnings(hands, false)
}

func part2(filename string) int {
	hands := readInput(filename)
	return calculateWinnings(hands, true)
}

func readInput(filename string) []HandAndBid {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hands []HandAndBid
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		hand := parts[0]
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, HandAndBid{hand, 0, bid})
	}

	return hands
}

func getHandType(hand string, part2 bool) int {
	cards := map[byte]int{}
	highestCount := 0
	for _, card := range hand {
		cards[byte(card)]++
		highestCount = max(highestCount, cards[byte(card)])
	}

	jokers := cards['J']
	switch len(cards) {
	case 1:
		return 6
	case 2:
		if part2 && jokers > 0 {
			return 6
		}
		if highestCount == 4 {
			return 5
		}
		return 4
	case 3:
		if highestCount == 3 {
			if part2 && jokers > 0 {
				return 5
			}
			return 3
		}
		if part2 && jokers > 1 {
			return 5
		} else if part2 && jokers > 0 {
			return 4
		}
		return 2
	case 4:
		if part2 && jokers > 0 {
			return 3
		}
		return 1
	case 5:
		if part2 && jokers > 0 {
			return 1
		}
		return 0
	}
	return 0
}

func calculateWinnings(hands []HandAndBid, part2 bool) int {
	for i := range hands {
		hands[i].handType = getHandType(hands[i].hand, part2)
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].handType == hands[j].handType {
			return lessHandType(hands[i].hand, hands[j].hand, part2)
		}
		return hands[i].handType < hands[j].handType
	})

	winnings := 0
	for i, entry := range hands {
		winnings += (i + 1) * entry.bid
	}
	return winnings
}

func lessHandType(hand1, hand2 string, part2 bool) bool {
	nums := map[byte]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}
	if part2 {
		nums['J'] = 1
	}
	for i := range hand1 {
		if hand1[i] != hand2[i] {
			return nums[hand1[i]] < nums[hand2[i]]
		}
	}
	return false
}

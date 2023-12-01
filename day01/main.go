package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
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
    digitRegex := regexp.MustCompile(`\d`)
    sum := 0

    for scanner.Scan() {
        line := scanner.Text()
        digits := digitRegex.FindAllString(line, -1)
        value, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
        if err != nil {
            log.Fatal(err)
        }
        sum += value
    }

    return sum
}


func findDigits(line string) []string {
    
    words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
    var matches []string

    i := 0
    for i < len(line) {
        
        digitMatched := false
        if unicode.IsDigit(rune(line[i])) {
            matches = append(matches, string(line[i]))
            i++
            digitMatched = true
            continue
        }

        wordMatched := false
        for _, word := range words {
            if strings.HasPrefix(line[i:], word) {
                matches = append(matches, word)
                if len(word) > 1 {
                    i += len(word) - 1 
                }
                wordMatched = true
                break
            }
        }

        if !wordMatched || digitMatched {
            i++
        }
    }

    return matches
}

func part2(filename string) int {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    sum := 0
    numberWords := map[string]string{
        "one": "1", "two": "2", "three": "3", "four": "4", 
        "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9",
    }

    for scanner.Scan() {
        line := scanner.Text()
        digits := findDigits(line)

        first, last := digits[0], digits[len(digits)-1]

        if val, ok := numberWords[first]; ok {
            first = val
        }
        
        if val, ok := numberWords[last]; ok {
            last = val
        }

        value, err := strconv.Atoi(first + last)
        if err != nil {
            log.Fatal(err)
        }
        sum += value
    }
    return sum
}

package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	expected := 13
	result := part1("test1.txt")
	if result != expected {
		t.Errorf("part1() = %d; want %d", result, expected)
	}
}

func TestPart2(t *testing.T) {
	expected := 30
	result := part2("test2.txt", 5)
	if result != expected {
		t.Errorf("part2() = %d; want %d", result, expected)
	}
}

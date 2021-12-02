package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/josiemessa/aoc2021/Go/utils"
)

const (
	forward string = "forward"
	down    string = "down"
	up      string = "up"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day2")
	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}

func Part1(lines []string) int {
	var horizontalVals []int
	var depthVals []int

	for i, line := range lines {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			log.Fatalf("Expected two parts to string on line %d but found %d (%q)\n", i, len(split), line)
		}

		val, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Second part of string should be int but was %q\n", split[1])
		}

		switch split[0] {
		case forward:
			horizontalVals = append(horizontalVals, val)
		case up:
			depthVals = append(depthVals, val*-1)
		case down:
			depthVals = append(depthVals, val)
		}
	}

	horizontal := 0
	for _, val := range horizontalVals {
		horizontal += val
	}

	depth := 0
	for _, val := range depthVals {
		depth += val
	}
	return horizontal * depth
}

func Part2(lines []string) int {
	var (
		horizontalPos int
		depthPos      int
		aim           int
	)

	for i, line := range lines {
		split := strings.Split(line, " ")
		if len(split) != 2 {
			log.Fatalf("Expected two parts to string on line %d but found %d (%q)\n", i, len(split), line)
		}

		val, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatalf("Second part of string should be int but was %q\n", split[1])
		}

		switch split[0] {
		case forward:
			horizontalPos += val
			depthPos += val*aim
		case up:
			aim -= val
		case down:
			aim += val
		}
	}
	return horizontalPos*depthPos
}

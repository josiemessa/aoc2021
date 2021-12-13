package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day7")
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer f.Close()
	fmt.Println(f.Name())

	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
	}

	//fmt.Println("Part 1: ", Part1(line))
	fmt.Println("Part 2: ", Part2(line))
}

func Part1(line string) int {
	// find the minimum total cost of fuel
	// parse input for crab positions
	input := strings.Split(line, ",")

	// calculate the median
	var crabs []int
	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not parse %q into int", s)
		}
		crabs = append(crabs, i)
	}

	sort.IntSlice(crabs).Sort()
	position := crabs[len(crabs)/2]

	var totalFuel int
	for _, crabPos := range crabs {
		totalFuel += int(math.Abs(float64(crabPos - position)))
	}
	return totalFuel
}

func Part2(line string) int {
	input := strings.Split(line, ",")

	// calculate the mean
	var crabs []int
	var cumulative int
	for _, s := range input {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not parse %q into int", s)
		}
		cumulative += i
		crabs = append(crabs, i)
	}
	position := cumulative/len(crabs)
	fmt.Println("position:", position)

	var totalFuel int
	for _, crabPos := range crabs {
		distance := int(math.Abs(float64(crabPos - position)))
		var cost int
		for i := 1; i <= distance; i++ {
			cost += i
		}
		totalFuel += cost
	}
	return totalFuel
}

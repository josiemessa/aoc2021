package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day3")
	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}

func Part1(lines []string) int {
	gammaRes, epsilonRes := ParseInput(lines)
	gammaInt, err := strconv.ParseInt(gammaRes, 2, 64)
	if err != nil {
		log.Fatalf("could not convert %q to int", gammaRes)
	}
	epsilonInt, err1 := strconv.ParseInt(epsilonRes, 2, 64)
	if err1 != nil {
		log.Fatalf("could not convert %q to int", gammaRes)
	}
	return int(gammaInt * epsilonInt)
}

func Part2(lines []string) int {
	return OxyGenRating(lines) * Co2Rating(lines)
}

func OxyGenRating(lines []string) int {
	// iterate over values to start filtering out
	// once we have only 1 value left we're done
	for i := 0; ; i++ {
		var newValues []string
		//fmt.Printf("%#v\n", lines)
		commonBit := FindCommonBit(lines, i)
		// on each iteration, look through the remaining candidate values
		for _, value := range lines {
			// if the bit in this position is the most common bit, keep it
			if value[i] == commonBit {
				newValues = append(newValues, value)
			}
		}

		//fmt.Printf("%#v\n", newValues)
		lines = newValues
		if len(lines) == 1 {
			break
		}
	}
	genRating, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		log.Fatalf("Could not parse final oxygen generator value %q", lines[0])
	}
	return int(genRating)
}

func Co2Rating(lines []string) int {
	// iterate over values to start filtering out
	// once we have only 1 value left we're done
	for i := 0; ; i++ {
		var newValues []string
		//fmt.Printf("%#v\n", lines)
		commonBit := FindCommonBit(lines, i)
		// on each iteration, look through the remaining candidate values
		for _, value := range lines {
			// if the bit in this position is the most common bit, keep it
			if value[i] != commonBit {
				newValues = append(newValues, value)
			}
		}

		//fmt.Printf("%#v\n", newValues)
		lines = newValues
		if len(lines) == 1 {
			break
		}
	}
	rating, err := strconv.ParseInt(lines[0], 2, 64)
	if err != nil {
		log.Fatalf("Could not parse final CO2 scrubber value %q", lines[0])
	}
	return int(rating)
}

func FindCommonBit(lines []string, pos int) uint8 {
	var ones int
	for _, line := range lines {
		// bit is a rune so it's the int ASCII representation of the character,
		// 48 => 0, 49 => 1
		ones += int(line[pos] - 48)
	}
	fmt.Printf("Found %d ones\n", ones)
	half := float64(len(lines)) / 2
	if float64(ones) >= half {
		return '1'
	}

	return '0'
}

func ParseInput(lines []string) (gammaRes, epsilonRes string) {
	var gamma []int
	// hacky way to initialise an array given that the input lengths differ between the test and the puzzle input
	for range lines[0] {
		gamma = append(gamma, 0)
	}
	for _, line := range lines {
		//fmt.Printf("%#v\n", line)
		for i, bit := range line {
			// bit is a rune so it's the int ASCII representation of the character,
			// 48 => 0, 49 => 1
			gamma[i] += int(bit - 48)
		}
	}
	for _, bitCount := range gamma {
		// are there more 1s than 0s?
		if bitCount >= len(lines)/2 {
			gammaRes += "1"
			epsilonRes += "0"
		} else {
			gammaRes += "0"
			epsilonRes += "1"
		}
	}
	return
}

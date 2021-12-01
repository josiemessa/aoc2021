package main

import (
	"fmt"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	input := utils.SliceAtoi(utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day1"))
	fmt.Println("Part 1: ", Part1(input))
	fmt.Println("Part 2: ", Part2(input))
}

func Part1(lines []int) (result int) {
	for i, line := range lines {
		if i == 0 {
			continue
		}
		if line > lines[i-1] {
			result++
		}
	}
	return
}

func Part2(lines []int) (result int) {
	for i := range lines {
		if i < 3 {
			continue
		}

		var sumA, sumB int
		for j := 0; j < 3; j++ {
			// window ending on i
			sumA += lines[i-j]
			// window ending on i-1 (previous window)
			sumB += lines[i-j-1]
		}
		if sumA > sumB {
			result++
		}
	}
	return
}
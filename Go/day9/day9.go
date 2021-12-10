package main

import (
	"fmt"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day9")
	heightmap := ParseInput(lines)
	fmt.Println("Part 1:", Part1(heightmap))
}

func Part1(heightmap [][]int) (result int) {
	var lowPoints []int
	for i, row := range heightmap {
		for j, value := range row {
			isLowPoint := true
			// check adjacent values, which will be i-1, i+1, j-1 and j+1
			if j > 0 {
				isLowPoint = isLowPoint && (value < row[j-1])
			}
			if j < len(row) - 1 {
				isLowPoint = isLowPoint && (value < row[j+1])
			}
			if i > 0 {
				isLowPoint = isLowPoint && (value < heightmap[i-1][j])
			}
			if i < len(heightmap) - 1 {
				isLowPoint = isLowPoint && (value < heightmap[i+1][j])
			}
			if isLowPoint {
				lowPoints = append(lowPoints, value + 1)
			}
		}
	}

	fmt.Println(lowPoints)
	for _, point := range lowPoints {
		result += point
	}
	return
}

func Part2(heightmap [][]int) (result int) {
	// need a way of storing the basins
	pointToBasin := make(map[Point]int)
	var basin int
	for i, row := range heightmap {
		for j, value := range row {
			isLowPoint := true
			// check adjacent values, which will be i-1, i+1, j-1 and j+1
			if j > 0 {
				isLowPoint = isLowPoint && (value < row[j-1])
			}
			if j < len(row) - 1 {
				isLowPoint = isLowPoint && (value < row[j+1])
			}
			if i > 0 {
				isLowPoint = isLowPoint && (value < heightmap[i-1][j])
			}
			if i < len(heightmap) - 1 {
				isLowPoint = isLowPoint && (value < heightmap[i+1][j])
			}
			if isLowPoint {
				basin++
				point := Point{ I: i, J: j }
				pointToBasin[point] = basin
			}
		}
	}
	return
}

type Point struct {
	I int
	J int
}

// {
//    { a_1, a_2, a_3},
//    { b_1, b_2, b_3 },
//}
func ParseInput(lines []string) [][]int {
	heightmap := make([][]int, len(lines))
	for i, line := range lines {
		row := make([]int, len(line))
		for j, digit := range line {
			row[j] = int(digit - 48)
		}
		heightmap[i] = row
	}
	return heightmap
}

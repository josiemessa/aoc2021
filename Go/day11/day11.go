package main

import (
	"fmt"
	"math"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day11")
	grid := make([][]int, 10)
	for i, line := range lines {
		row := make([]int, 10)
		for j, x := range line {
			row[j] = int(x - 48)
		}
		grid[i] = row
	}

	Part1(grid)

	//fmt.Println("Part1:", Part1Result)
	fmt.Println("Part2:", Part2(grid))
}

var Part1Result int

func Part1(grid [][]int) {
	for i := 0; i < 100; i++ {
		Stage1(grid)
		Stage2(grid)
		Stage3(grid)
	}
}

func Part2(grid [][]int) int {
	i := 0
	for  {
		i++
		Part2Result = 0
		Stage1(grid)
		Stage2(grid)
		Stage3(grid)
		if Part2Result == 100 {
			return i
		}
	}
}

// Stage 1: Increment everything by 1
func Stage1(grid [][]int) {
	for i, row := range grid {
		for j := range row {
			grid[i][j]++
		}
	}
}

func Stage2(grid [][]int) {
	for i, row := range grid {
		for j := range row {
			CheckForFlash(i,j,grid)
		}
	}
}

var Part2Result int

func Stage3(grid [][]int) {
	for i, row := range grid {
		for j := range row {
			if grid[i][j] < 0 {
				grid[i][j] = 0
				Part2Result++
			}
		}
	}
}

func CheckForFlash(i, j int, grid [][]int) {
	if grid[i][j] > 9 {
		grid[i][j] = math.MinInt8
		Part1Result++
		IncrementAdjacent(i,j, grid)
	}
}

func IncrementAdjacent(i, j int, grid [][]int) {
	if i-1 >= 0 {
		grid[i-1][j]++
		CheckForFlash(i-1, j, grid)
		if j-1 >= 0 {
			grid[i-1][j-1]++
			CheckForFlash(i-1, j-1, grid)
		}
		if j+1 < len(grid[i]) {
			grid[i-1][j+1]++
			CheckForFlash(i-1, j+1, grid)
		}
	}
	if j-1 >= 0 {
		grid[i][j-1]++
		CheckForFlash(i, j-1, grid)
		if i+1 < len(grid) {
			grid[i+1][j-1]++
			CheckForFlash(i+1, j-1, grid)
		}
	}
	if i+1 < len(grid) {
		grid[i+1][j]++
		CheckForFlash(i+1, j, grid)
		if j+1<len(grid) {
			grid[i+1][j+1]++
			CheckForFlash(i+1, j+1, grid)
		}
	}
	if j+1 < len(grid) {
		grid[i][j+1]++
		CheckForFlash(i, j+1,grid)
	}

}

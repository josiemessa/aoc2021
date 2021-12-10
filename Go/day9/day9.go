package main

import (
	"fmt"
	"sort"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day9")
	heightmap := ParseInput(lines)
	fmt.Println("Part 1:", Part1(heightmap))
	fmt.Println("Part 2:", Part2(heightmap))
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
			if j < len(row)-1 {
				isLowPoint = isLowPoint && (value < row[j+1])
			}
			if i > 0 {
				isLowPoint = isLowPoint && (value < heightmap[i-1][j])
			}
			if i < len(heightmap)-1 {
				isLowPoint = isLowPoint && (value < heightmap[i+1][j])
			}
			if isLowPoint {
				lowPoints = append(lowPoints, value+1)
			}
		}
	}

	fmt.Println(lowPoints)
	for _, point := range lowPoints {
		result += point
	}
	return
}

type HeightMap struct {
	HeightMap [][]int
	VisitedPoints map[Point]struct{}
}

func Part2(heightmap [][]int) int {
	h := &HeightMap{
		heightmap,
		make(map[Point]struct{}),
	}
	result := 1
	size := 0
	var sizes []int
	for i, row := range heightmap {
		for j := range row {
			p := Point{i, j}
			size += h.TraverseBasin(p)
			// size will be 0 if we've already seen this point or it's a 9
			// otherwise we'll have the full traversal of the basin
			if size != 0 {
				fmt.Println("Found basin of size", size)
				sizes = append(sizes, size)
				size = 0
			}
		}
	}
	sort.IntSlice(sizes).Sort()
	for i := 0; i < 3; i++ {
		result *= sizes[len(sizes)-1-i]
	}
	return result
}

func (h *HeightMap) TraverseBasin(p Point) int {
	if _, ok := h.VisitedPoints[p]; ok {
		return 0
	}
	if h.HeightMap[p.I][p.J] == 9 {
		return 0
	}
	// visit this point
	h.VisitedPoints[p] = struct{}{}

	size := 1

	if p.J > 0 {
		size += h.TraverseBasin(Point{I: p.I, J: p.J-1})
	}
	if p.J < len(h.HeightMap[p.I])-1 {
		size += h.TraverseBasin(Point{I: p.I, J: p.J+1})
	}
	if p.I > 0 {
		size += h.TraverseBasin(Point{I: p.I-1, J: p.J})
	}
	if p.I < len(h.HeightMap)-1 {
		size += h.TraverseBasin(Point{I: p.I+1, J: p.J})
	}

	return size
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

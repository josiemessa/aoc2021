package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day5")
	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}

//TODO: Parse input to find *horizontal* and *vertical* lines only
// we need to iterate over points in the grid to see how many lines intersect that point
// count how many points where 2 or more lines overlap
// We may need to track how many intersecting lines there are for part 2

// ideas:
// Can we sort the points? Points don't necessarily have it such that x1 > x2 though.
// We don't actually care about direction so we could just check both parts of the tuple
// order by the distance from (0,0), x taking priority

// could expand points into a line
// throw away diagonals

func Part1(input []string) (result int) {
	grid := make(map[Point]int)
	for _, entry := range input {
		// Parse input entry for coordinates
		points := strings.Split(entry, " -> ")
		line := Line{
			Start: ParsePoint(points[0]),
			End:   ParsePoint(points[1]),
		}

		// Discard any line definitions that are not horizontal or vertical
		if line.IsDiagonal() {
			continue
		}

		// Ensure all lines are increasing (so we don't have to worry about direction)
		// and grab horizontal/vertical distinction
		line.OrderPoints()

		// Compute the coordinates of the line
		addToGrid := func(p Point) {
			if _, ok := grid[p]; ok {
				grid[p]++
			} else {
				grid[p] = 1
			}
		}

		if line.Vertical {
			for y := line.Start.Y; y <= line.End.Y; y++ {
				addToGrid(Point{X: line.Start.X, Y: y})
			}
		} else {
			for x := line.Start.X; x <= line.End.X; x++ {
				addToGrid(Point{X: x, Y: line.Start.Y})
			}
		}
	}

	for _, count := range grid {
		if count > 1 {
			result++
		}
	}
	return
}

func Part2(input []string) (result int) {
	grid := make(map[Point]int)
	for _, entry := range input {
		// Parse input entry for coordinates
		points := strings.Split(entry, " -> ")
		line := Line{
			Start: ParsePoint(points[0]),
			End:   ParsePoint(points[1]),
		}

		addToGrid := func(p Point) {
			if _, ok := grid[p]; ok {
				grid[p]++
			} else {
				grid[p] = 1
			}
		}
		
		line.GenerateLineFromPoints()
		//fmt.Printf("%v -> %v: %v\n", line.Start, line.End, line.Line)
		for _, point := range line.Line {
			addToGrid(point)
		}
	}

	for _, count := range grid {
		if count > 1 {
			result++
		}
	}
	return
}

type Line struct {
	Start    Point
	End      Point
	Line     []Point
	Vertical bool
}

func (l *Line) IsDiagonal() bool {
	return l.Start.X != l.End.X && l.Start.Y != l.End.Y
}

func (l *Line) OrderPoints() {
	if l.Start.X == l.End.X {
		l.Vertical = true

		if l.Start.Y > l.End.Y {
			old := l.Start
			l.Start = l.End
			l.End = old
		}
	} else if l.Start.X > l.End.X {
		old := l.Start
		l.Start = l.End
		l.End = old
	} else if l.Start.Y > l.End.Y {
		old := l.Start
		l.Start = l.End
		l.End = old
	}
}

func (l *Line) GenerateLineFromPoints() {
	var start, end int

	if l.Start.X == l.End.X {
		// line is vertical
		if l.Start.Y <= l.End.Y {
			start = l.Start.Y
			end = l.End.Y
		} else {
			start = l.End.Y
			end = l.Start.Y
		}
		for y := start; y <= end; y++ {
			l.Line = append(l.Line, Point{X: l.Start.X, Y: y})
		}
		return
	} else if l.Start.Y == l.End.Y {
		// line is horizontal {
		if l.Start.X <= l.End.X {
			start = l.Start.X
			end = l.End.X
		} else {
			start = l.End.X
			end = l.Start.X
		}
		for x := start; x <= end; x++ {
			l.Line = append(l.Line, Point{X: x, Y: l.Start.Y})
		}
		return
	}
	// line is diagonal
	// ignore direction and ensure that x is always increasing
	if l.Start.X > l.End.X {
		old := l.Start
		l.Start = l.End
		l.End = old
	}
	distance := l.End.X - l.Start.X
	var direction = 1
	if l.Start.Y > l.End.Y {
		direction = -1
	}
	for i := 0; i <= distance; i++ {
		l.Line = append(l.Line, Point{X: l.Start.X+i, Y:l.Start.Y+(i*direction)})
	}
}

type Point struct {
	X int
	Y int
}

// ParsePoint takes a string in the format "x,y" and parses into a point
func ParsePoint(input string) Point {
	xy := strings.Split(input, ",")
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		log.Fatalln("Could not parse x coordinate for input", input)
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		log.Fatalln("Could not parse y coordinate for input", input)
	}
	return Point{
		X: x,
		Y: y,
	}
}

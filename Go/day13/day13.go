package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	input := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day13")
	coords, folds := ParseInput(input, false)
	//fmt.Println("Part 1:", Part1(coords, folds[0]))
	points := Part2(coords, folds)
	Format(points)
}

func ParseInput(input []string, part1 bool) (coords map[Point]struct{}, folds []Fold) {
	coords = make(map[Point]struct{})
	var foldInstructions []string
	for i, line := range input {
		if line == "" {
			foldInstructions = input[i+1:]
			break
		}
		xy := strings.Split(line, ",")
		x, err := strconv.Atoi(xy[0])
		if err != nil {
			log.Fatalln("Could not parse line", line)
		}
		p := Point{X: x}
		y, err := strconv.Atoi(xy[1])
		if err != nil {
			log.Fatalln("could not parse line", line)
		}
		p.Y = y
		coords[p] = struct{}{}
	}

	for _, line := range foldInstructions {
		line = strings.Trim(line, "fold along ")
		if line[0] == 'x' {
			line = strings.Trim(line, "x=")
			x, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalln("Could not parse fold instruction for x line", line)
			}
			folds = append(folds, Fold{
				Axis:  X,
				Value: x,
			})
		} else {
			line = strings.Trim(line, "y=")
			y, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalln("Could not parse fold instruction for y line", line)
			}
			folds = append(folds, Fold{
				Axis:  Y,
				Value: y,
			})
		}
		if part1 {
			break
		}
	}
	return
}

type Point struct {
	X int
	Y int
}

type Axis uint8

const (
	X Axis = iota
	Y
)

type Fold struct {
	Axis  Axis
	Value int
}

func Part1(coords map[Point]struct{}, fold Fold) (count int) {
	folded := make(map[Point]struct{})
	for point := range coords {
		if fold.Axis == X {
			folded[FoldLeft(point, fold.Value)] = struct{}{}
		} else {
			folded[FoldUp(point, fold.Value)] = struct{}{}
		}
	}
	return len(folded)
}

func Part2(coords map[Point]struct{}, folds []Fold) map[Point]struct{} {
	for _, fold := range folds {
		folded := make(map[Point]struct{})
		for point := range coords {
			if fold.Axis == X {
				folded[FoldLeft(point, fold.Value)] = struct{}{}
			} else {
				folded[FoldUp(point, fold.Value)] = struct{}{}
			}
		}
		coords = folded
	}
	return coords
}

func Format(coords map[Point]struct{}) {
	var maxX int
	var maxY int
	for point:= range coords {
		if point.Y > maxY {
			maxY = point.Y
		}
		if point.X > maxX {
			maxX = point.X
		}
	}
	grid := make([][]rune, maxY+5)
	for i:= range grid {
		grid[i] = make([]rune, maxX+5)
		for j := range grid[i] {
			if _, ok := coords[Point{X:j, Y:i}]; ok {
				grid[i][j] = '#'
			} else {
				grid[i][j] = ' '
			}
		}
	}
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func FoldLeft(point Point, xValue int) Point {
	if point.X <= xValue {
		return point
	}
	distance := point.X - xValue
	return Point {
		X: xValue - distance,
		Y: point.Y,
	}
}

func FoldUp(point Point, yValue int) Point {
	if point.Y <= yValue {
		return point
	}
	distance := point.Y - yValue
	return Point {
		X: point.X,
		Y: yValue - distance,
	}
}

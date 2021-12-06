package main

import (
	"fmt"
	"strings"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day4")
	fmt.Printf("Part 1: %d\n", Part1(lines))
	fmt.Printf("Part 2: %d\n", Part2(lines))
}

func Part1(lines []string) int {
	drawn, boards := ParseInput(lines)

	for _, number := range drawn {
		fmt.Printf("\n-- Drawing %d --\n", number)
		for i, board := range boards {
			board.PlayBingoRound(number)
			fmt.Printf(" on board %d\n", i)
			if board.Bingo {
				// calculate score
				fmt.Println("> unmarked numbers: ", board.SumUnmarked())
				fmt.Println("> drawin number: ", number)
				return board.SumUnmarked() * number
			}
		}
	}

	return -1

	// other notes:
	//   - Numbers range from 0-99
	//   - boards are 5x5
	//   - ASSUMING a number can only appear once on a board
}

func Part2(lines []string) int {
	drawn, boards := ParseInput(lines)
	var winningBoards []int
	var lastScore int
	for _, number := range drawn {

		fmt.Printf("\n-- Drawing %d --\n", number)
		for i, board := range boards {
			if Contains(winningBoards, i) {
				continue
			}
			board.PlayBingoRound(number)
			fmt.Printf(" on board %d\n", i)

			if board.Bingo {
				// calculate score
				fmt.Println("> unmarked numbers: ", board.SumUnmarked())
				fmt.Println("> drawing number: ", number)
				fmt.Println("> board score:", board.SumUnmarked()*number)
				lastScore = board.SumUnmarked()*number

				winningBoards = append(winningBoards, i)
			}
		}
	}

	return lastScore

	// other notes:
	//   - Numbers range from 0-99
	//   - boards are 5x5
	//   - ASSUMING a number can only appear once on a board
}

func ParseInput(lines []string) ([]int, []Board) {
	fmt.Println("> parsing drawn numbers")
	drawnNumLine := strings.Split(lines[0], ",")
	drawnNumbers := utils.SliceAtoi(drawnNumLine)

	// remove first two lines as we've already parsed them
	lines = lines[2:]

	// parse the rest of the input as boards
	fmt.Println("> parsing board")
	var boards []Board
	b := Board{
		Columns: make([][]int, 5),
	}
	for _, line := range lines {
		if line == "" {
			boards = append(boards, b)
			// rows are appended but we need to index into the columns so make them in advance
			b = Board{
				Columns: make([][]int, 5),
			}
			continue
		}

		// Put numbers into their rows and columns in the bingo board
		row := utils.SliceAtoi(strings.Split(line, " "))
		b.Rows = append(b.Rows, row)
		for i, r := range row {
			b.Columns[i] = append(b.Columns[i], r)
		}
	}
	boards = append(boards, b)

	return drawnNumbers, boards
}

func RemoveIntFromSlice(s []int, value int) []int {
	for i, e := range s {
		if e == value {
			s = RemoveIndexFromSlice(s, i)
		}
	}
	return s
}

func RemoveIndexFromSlice(s []int, index int) []int {
	if index == len(s)-1 {
		s = s[:index]
	} else {
		s = append(s[:index], s[index+1:]...)
	}
	return s
}

func Contains(s []int, val int) bool {
	for _, e := range s {
		if val == e {
			return true
		}
	}
	return false
}

package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func TestParseInput(t *testing.T) {
	lines := strings.Split(input, "\n")
	drawn, boards := ParseInput(lines)

	require.Len(t, drawn, 27)
	require.Equal(t, 7, drawn[0])
	require.Equal(t, 1, drawn[26])

	require.Len(t, boards, 3, "incorrect number of boards parsed")
	for _, board := range boards {
		require.Len(t, board.Rows, 5)
		require.Len(t, board.Columns, 5)
	}

	require.Contains(t, boards[0].Rows, []int{21, 9, 14, 16, 7})
	require.Contains(t, boards[0].Columns, []int{11, 4, 16, 18, 15})
}

func TestBoard_PlayBingoRound(t *testing.T) {
	board := Board{
		Rows: [][]int{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7}, // 7
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		Columns: [][]int{
			{22, 8, 21, 6, 1},
			{13, 2, 9, 10, 12},
			{17, 23, 14, 3, 20},
			{11, 4, 16, 18, 15},
			{0, 24, 7, 5, 19}, // 7
		},
		Bingo: false,
	}

	drawn := 7
	board.PlayBingoRound(drawn)
	require.Len(t, board.Rows[2], 4)
	require.NotContains(t, board.Rows[2], 7)

	require.Len(t, board.Columns[4], 4)
	require.NotContains(t, board.Columns[4], 7)
	require.False(t, board.Bingo)

	for i, row := range board.Rows {
		if i == 2 {
			continue
		}
		require.Len(t, row, 5)
	}

	for i, row := range board.Columns {
		if i == 4 {
			continue
		}
		require.Len(t, row, 5)
	}

}

func TestPart1(t *testing.T) {
	require.Equal(t, 4512, Part1(strings.Split(input, "\n")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 1924, Part2(strings.Split(input, "\n")))
}

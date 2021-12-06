package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

func TestPart1(t *testing.T) {
	require.Equal(t, 5, Part1(strings.Split(input, "\n")))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 12, Part2(strings.Split(input, "\n")))
}

func TestParsePoint(t *testing.T) {
	require.Equal(t, Point{X: 8, Y: 0}, ParsePoint("8,0"))
}

func TestLine_IsDiagonal(t *testing.T) {
	l := Line{
		Start: Point{8, 0},
		End:   Point{0, 8},
	}
	require.True(t, l.IsDiagonal())
}

func TestLine_OrderPoints(t *testing.T) {
	l := Line{
		Start: Point{9, 4},
		End:   Point{3, 4},
	}
	l.OrderPoints()
	require.False(t, l.Vertical)
	require.Equal(t, Point{3,4}, l.Start)
	require.Equal(t, Point{9,4}, l.End)

	//2,2 -> 2,1
	l2 := Line{
		Start: Point{2,2},
		End: Point{2,1},
	}

	l2.OrderPoints()
	require.True(t, l2.Vertical)
	require.Equal(t, Point{2,1}, l2.Start)
	require.Equal(t, Point{2,2}, l2.End)


}

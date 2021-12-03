package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

func TestOxyGenRating(t *testing.T) {
	lines := strings.Split(input, "\n")
	require.Equal(t, 23, OxyGenRating(lines))
}

func TestCo2Rating(t *testing.T) {
	lines := strings.Split(input, "\n")
	require.Equal(t, 10, Co2Rating(lines))
}

func TestPart2(t *testing.T) {
	lines := strings.Split(input, "\n")
	require.Equal(t, 230, lines)
}
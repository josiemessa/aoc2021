package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `2199943210
3987894921
9856789892
8767896789
9899965678`
func TestPart1(t *testing.T) {
	lines := ParseInput(strings.Split(input, "\n"))
	require.Equal(t, 15, Part1(lines))
}

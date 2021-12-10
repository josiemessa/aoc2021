package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPart2(t *testing.T) {
	lines := []string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"}
	require.Equal(t, 5353, Part2(lines))
}

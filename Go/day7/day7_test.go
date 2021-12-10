package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var line = `16,1,2,0,4,2,7,1,2,14`

func TestPart1(t *testing.T) {
	require.Equal(t, 37, Part1(line))
}

func TestPart2(t *testing.T) {
	require.Equal(t, 168, Part2(line))
}
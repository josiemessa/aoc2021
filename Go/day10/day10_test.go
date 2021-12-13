package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

func TestPart1(t *testing.T) {
	p1, p2 := Solve(strings.Split(input, "\n"))
	t.Run("part 1", func(t *testing.T) {
		require.Equal(t, 26397, p1)
	})
	t.Run("part 2", func(t *testing.T) {
		require.Equal(t, 288957, p2)
	})
}

package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var input = `start-A
start-b
A-c
A-b
b-d
A-end
b-end`

var largerInput = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

var largestInput = `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`

func TestVisitCaveP2(t *testing.T) {
	t.Run("smaller example", func(tt *testing.T) {
		caveSystem := ParseInput(strings.Split(input, "\n"))
		require.Equal(tt, 36, Part2(caveSystem))
	})
	t.Run("larger example", func(tt *testing.T){
		caveSystem := ParseInput(strings.Split(largerInput, "\n"))
		require.Equal(tt, 103, Part2(caveSystem))
	})
	t.Run("largest example", func(tt *testing.T){
		caveSystem := ParseInput(strings.Split(largestInput, "\n"))
		require.Equal(tt, 3509, Part2(caveSystem))
	})
}

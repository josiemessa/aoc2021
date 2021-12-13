package main

import (
	"fmt"
	"log"
	"sort"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day10")
	p1, p2 := Solve(lines)
	fmt.Println("Part 1:", p1)
	fmt.Println("Part 2:", p2)
}

var p1Scores = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var p2Scores = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

var charPairs = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

//
//func IsClosingChar(char rune) bool {
//	return char == ')' || char == ']' || char == '}' || char == '>'
//}

func IsOpeningChar(char rune) bool {
	return char == '(' || char == '[' || char == '{' || char == '<'
}

// ParseChunk parses the line for chunks, and returns the length of the line parsed,
// and whether the chunk was valid
func Solve(lines []string) (int, int) {
	p1Result := 0
	var lineScores []int
	for _, line := range lines {
		var stack []int32
		for _, char := range line {
			if IsOpeningChar(char) {
				stack = append(stack, char)
				continue
			}
			openingChar := stack[len(stack)-1]
			closingChar, ok := charPairs[openingChar]
			if !ok {
				// probably means opening char wasn't set so we should look at line
				log.Fatalf("Could not find closing char for opening char %q in line %q\n", openingChar, line)
			}
			if char != closingChar {
				//fmt.Printf("Expected %q but found %q at position %d in line %q\n", charPairs[openingChar], char, i, line)
				p1Result += p1Scores[char]
				stack = []int32{}
				break
			}
			// closing char is the correct one, so remove from the stack
			stack = stack[:len(stack)-1]
		}
		if len(stack) != 0 {
			// incomplete line
			lineScore := 0
			l := len(stack)
			for i := 0; i < l; i++ {
				var missingChar = charPairs[stack[len(stack)-1]]
				fmt.Printf("%s", string(missingChar))
				lineScore *= 5
				lineScore += p2Scores[missingChar]
				stack = stack[:len(stack)-1]
			}
			fmt.Printf("\n")
			lineScores = append(lineScores, lineScore)
		}
	}
	sort.IntSlice(lineScores).Sort()
	fmt.Println(len(lineScores), lineScores)
	return p1Result, lineScores[len(lineScores)/2]
}

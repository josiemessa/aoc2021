package main

import (
	"fmt"
	"strings"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day14")
	template := lines[0]
	var insertionRules []*InsertionRule
	for _, s := range lines[2:] {
		rule := strings.Split(s, " -> ")
		insertionRules = append(insertionRules, &InsertionRule{
			PairPrefix: rule[0][0],
			PairSuffix: rule[0][1],
			Insertion:  rule[1][0],
		})
	}
	//Part1(template, insertionRules)
	Part2(template, insertionRules)
}

type Pair struct {
	Prefix uint8
	Suffix uint8
}

func Part2(template string, insertionRules []*InsertionRule) {
	// try to reuse computed insertion iterations
	// so you only need to calculate how it grows once,
	// then just append results onto the end
	rules := make(map[Pair]string)
	for _, rule := range insertionRules {
		rules[Pair{Prefix: rule.PairPrefix, Suffix: rule.PairSuffix}] = string(rule.Insertion)
	}

	iterations := make(map[Pair][]string)
	for i := range template {
		if i == 0 {
			continue
		}
		p := Pair{Prefix: template[i-1], Suffix: template[i]}
		if s, ok := iterations[p]; !ok {
			// calculate the 40 iterations for this pair
			for i := 0; i < 40; i++ {

			}
		}

	}
}

func GetIterations(p Pair, i int, insertResults map[Pair][]string) []string {
	if s, ok := insertResults[p]; ok {
		if len(s) >= i {
			return s[:i]
		} else {
			for j := range s[len(s)-1] {
				
			}
		}
	}
}

func Part1(template string, insertionRules []*InsertionRule) {
	for i := 0; i < 10; i++ {
		var newTemplate string
		for j := range template {
			if j == 0 {
				continue
			}
			newTemplate += string(template[j-1])
			for _, rule := range insertionRules {
				if rule.PairPrefix == template[j-1] && rule.PairSuffix == template[j] {
					newTemplate += string(rule.Insertion)
				}
			}
		}
		newTemplate += string(template[len(template)-1])
		template = newTemplate
		if i < 3 {
			fmt.Println(template)
		}
	}

	freq := make(map[rune]int)

	var max rune = 'N'
	var min rune = 'N'
	for _, letter := range template {
		if _, ok := freq[letter]; !ok {
			freq[letter] = 1
		} else {
			freq[letter]++
		}
		if freq[max] < freq[letter] {
			max = letter
		}
		if freq[min] > freq[letter] {
			min = letter
		}
	}

	fmt.Println("Part 1:", freq[max]-freq[min])
}

type InsertionRule struct {
	PairPrefix uint8
	PairSuffix uint8
	Insertion  uint8
}

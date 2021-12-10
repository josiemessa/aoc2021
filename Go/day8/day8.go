package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day8")
	fmt.Println("Part 1:", Part1(lines))
	fmt.Println("Part 2:", Part2(lines))
}

var display = map[string]string{
	"abcefg":  "0",
	"cf":      "1",
	"acdeg":   "2",
	"acdfg":   "3",
	"bcdf":    "4",
	"abdfg":   "5",
	"abdefg":  "6",
	"acf":     "7",
	"abcdefg": "8",
	"abcdfg":  "9",
}

func Part1(lines []string) (count int) {
	for _, line := range lines {
		split := strings.Split(line, " | ")
		//hints := strings.Split(split[0], " ")
		digits := strings.Split(split[1], " ")

		// 1: 2 segments
		// 4: 4 segments
		// 7: 3 segments
		// 8: 7 segments
		//var encoding []string
		for _, digit := range digits {
			switch len(digit) {
			case 2, 3, 4, 7:
				count++
			default:
				continue
			}
		}
	}
	return
}

func Part2(lines []string) (result int) {
	for _, line := range lines {
		split := strings.Split(line, " | ")
		hints := strings.Split(split[0], " ")
		digits := strings.Split(split[1], " ")

		encoding := make([]string, 10)
		frequency := make(map[string]int) // count number of times we see a-g
		for _, hint := range hints {
			switch len(hint) {
			case 2:
				encoding[1] = hint
			case 3:
				encoding[7] = hint
			case 4:
				encoding[4] = hint
			case 7:
				encoding[8] = hint
			}

			for _, letter := range hint {
				if _, ok := frequency[string(letter)]; !ok {
					frequency[string(letter)] = 1
				} else {
					frequency[string(letter)]++
				}
			}
		}

		// work out wire connectors
		segment := make(map[string]string)

		// segment a is the {encoding of 7} - {encoding of 1}
		segment["a"] = Complement(encoding[1], encoding[7])

		// 8: a b c d e f g
		// 9: a b c d   f g
		// 0: a b c   e f g
		// 6: a b   d e f g
		// 5: a b   d   f g
		// 2: a   c d e   g
		// 3: a   c d   f g
		// 4:   b c d   f
		// 7: a   c     f
		// 1:     c     f

		// a: 8/10, b: *6/10*, c: 8/10, d: 7/10, e: *4/10*, f: *9/10*, g: 7/10
		for k, v := range frequency {
			switch v {
			case 4:
				segment["e"] = k
			case 6:
				segment["b"] = k
			case 9:
				segment["f"] = k
			}
		}
		// we now know a, b, e and f
		// work out c from 1
		segment["c"] = Complement(segment["f"], encoding[1])

		// can work out d from encoding of 4 (bcdf)
		segment["d"] = Complement(segment["b"]+segment["c"]+segment["f"], encoding[4])

		// we now know a, b, c, d, e and f, can work out g from 8
		segment["g"] = Complement(segment["a"]+segment["b"]+segment["c"]+segment["d"]+segment["e"]+segment["f"], encoding[8])

		fmt.Println(segment)

		var decodedDigits string
		for _, digit := range digits {
			var decoded string
			for _, letter := range digit {
				k := GetKey(segment, string(letter))
				if k == "" {
					log.Fatalln("could not find key for letter", string(letter), "in digit", digit, "in line\n", line)
				}
				decoded += k
			}
			var sorted []int
			for _, d := range decoded {
				sorted = append(sorted, int(d))
			}
			sort.IntSlice(sorted).Sort()
			decoded = ""
			for _, d := range sorted {
				decoded += string(uint8(d))
			}

			if val, ok := display[decoded]; !ok {
				log.Fatalln("decoded digit", digit, "to display", decoded, "but could not find value. From line\n", line)
			} else {
				decodedDigits += val
			}
		}
		i, err := strconv.Atoi(decodedDigits)
		if err != nil {
			log.Fatalln("decoded display to", decodedDigits,  "but could not parse int. From line\n\"", line)
		}
		result += i
	}
	return
}

// Complement calculates the complement of set A with B.
// i.e. all elements in B that are not in A
func Complement(a string, b string) (complement string) {
	for _, s := range b {
		if !Contains(a, s) {
			complement += string(s)
		}
	}
	return
}

// Contains returns true if s is contained in a, and false otherwise
func Contains(a string, s rune) (found bool) {
	for _, x := range a {
		if x == s {
			found = true
		}
	}
	return found
}

func GetKey(m map[string]string, value string) string {
	for k, v := range m {
		if v == value {
			return k
		}
	}
	return ""
}

/// 0:      1:      2:      3:      4:
// aaaa    ....    aaaa    aaaa    ....
//b    c  .    c  .    c  .    c  b    c
//b    c  .    c  .    c  .    c  b    c
// ....    ....    dddd    dddd    dddd
//e    f  .    f  e    .  .    f  .    f
//e    f  .    f  e    .  .    f  .    f
// gggg    ....    gggg    gggg    ....
//
//  5:      6:      7:      8:      9:
// aaaa    aaaa    aaaa    aaaa    aaaa
//b    .  b    .  .    c  b    c  b    c
//b    .  b    .  .    c  b    c  b    c
// dddd    dddd    ....    dddd    dddd
//.    f  e    f  .    f  e    f  .    f
//.    f  e    f  .    f  e    f  .    f
// gggg    gggg    ....    gggg    gggg

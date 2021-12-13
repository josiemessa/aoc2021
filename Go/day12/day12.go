package main

import (
	"fmt"
	"strings"

	"github.com/josiemessa/aoc2021/Go/utils"
)

func main() {
	lines := utils.ReadFileAsLines("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day12")
	caveSystem := ParseInput(lines)
	//fmt.Println("Part 1:", Part1(caveSystem))
	fmt.Println("Part 2:", Part2(caveSystem))
}

func ParseInput(lines []string) map[string][]string {
	caveSystem := make(map[string][]string)
	for _, line := range lines {
		split := strings.Split(line, "-")
		caveSystem[split[0]] = append(caveSystem[split[0]], split[1])
		caveSystem[split[1]] = append(caveSystem[split[1]], split[0])
	}
	for cave, connected := range caveSystem {
		fmt.Printf("%s: %v\n", cave, connected)
	}
	return caveSystem
}

const (
	start string = "start"
	end   string = "end"
	// in an ASCII chart, the upper case characters appear before lowercase characters
	// so if any character is <= this rune, it is upper, otherwise lower case.
	upperCase uint8 = 'Z'
)

func Part1(caveSystem map[string][]string) int {
	// Need to start at start and end at end
	Prune(caveSystem)
	routes := VisitCave(caveSystem, start, []string{})
	return len(routes)
}

func Part2(caveSystem map[string][]string) int {
	routes := VisitCaveP2(caveSystem, start, []string{}, true)
	//for _, route := range routes {
	//	fmt.Println(route)
	//}
	return len(routes)
}

func VisitCave(caveSystem map[string][]string, cave string, prevRoute []string) (routes [][]string) {
	if IsSmallCave(cave) && utils.StringSliceContains(prevRoute, cave) {
		return nil
	}
	currentRoute := append(prevRoute, cave)
	if cave == end {
		return [][]string{currentRoute}
	}

	for _, connected := range caveSystem[cave] {
		if subRoute := VisitCave(caveSystem, connected, currentRoute); subRoute != nil {
			routes = append(routes, subRoute...)
		}
	}
	return routes
}

func VisitCaveP2(caveSystem map[string][]string, cave string, prevRoute []string, canVisitTwice bool) [][]string {
	if IsSmallCave(cave) && utils.StringSliceContains(prevRoute, cave) {
		if canVisitTwice && cave != start && cave != end {
			// visit this cave twice
			canVisitTwice = false
		} else {
			// we've already visited a cave twice, so exit
			return nil
		}
	}
	currentRoute := make([]string, len(prevRoute), len(prevRoute)+1)
	copy(currentRoute, prevRoute)
	currentRoute = append(currentRoute, cave)
	if cave == end {
		return [][]string{currentRoute}
	}

	var routes [][]string
	for _, connected := range caveSystem[cave] {
		if subRoute := VisitCaveP2(caveSystem, connected, currentRoute, canVisitTwice); subRoute != nil {
			routes = append(routes, subRoute...)
		}
	}
	return routes
}

// Prune removes "dead ends" - i.e. small caves that only lead to one
// other small cave. like "d" in the example, but not "c" (this leads to a large cave)
func Prune(caveSystem map[string][]string) {
	for cave, connections := range caveSystem {
		if cave == start || cave == end {
			continue
		}
		// if a small cave only has one connection to one other small cave, then we will not visit it
		if IsSmallCave(cave) && len(connections) == 1 && IsSmallCave(connections[0]) {
			fmt.Printf("Deleting cave %q\n", cave)
			delete(caveSystem, cave)
			// re-prune as we've altered the caveSystem and don't want to keep
			// iterating over this version of it
			Prune(caveSystem)
			return
		}
	}
}

func IsSmallCave(cave string) bool {
	return cave[0] > upperCase
}

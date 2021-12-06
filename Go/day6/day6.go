package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("C:\\dev\\src\\github.com\\josiemessa\\aoc2021\\Inputs\\Day6")
	if err != nil {
		log.Fatal("Could not open file", err)
	}
	defer f.Close()
	fmt.Println(f.Name())

	scanner := bufio.NewScanner(f)
	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
	}
	fmt.Println("Part 1: ", Part1(line, 80))
	fmt.Println("Part 2: ", Part2(line, 256))
}

func Part1(input string, days int) int {
	init := strings.Split(input, ",")
	var fish []int
	for _, s := range init {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not parse %q into int", s)
		}
		fish = append(fish, i)
	}

	var newFish []int
	for i := 0; i < days; i++ {
		for j := range fish {
			fish[j]--
			if fish[j] == -1 {
				fish[j] = 6
				newFish = append(newFish, 8)
			}
		}
		fish = append(fish, newFish...)
		newFish = []int{}
		//fmt.Printf("%#v\n", fish)
	}

	return len(fish)
}

func Part2(input string, days int) (totalSpawn uint64) {
	init := strings.Split(input, ",")
	// map of timer -> # of fish with that timer
	fish := make([]int, 9)
	for _, s := range init {
		i, err := strconv.Atoi(s)
		if err != nil {
			log.Fatalf("Could not parse %q into int", s)
		}
		fish[i]++
	}

	for day := 0; day < days; day++ {
		newFish := make([]int, 9)
		for timer, count := range fish {
			if timer == 0 {
				newFish[6] = count
				newFish[8] = count
				continue
			}
			newFish[timer-1] += count
		}
		fish = newFish
	}

	for _, val := range fish {
		totalSpawn += uint64(val)
	}
	return
}

// CalculateSpawn calculates how many fish will be spawned by 1 fish,
// which has an initial condition of init over days number of days
// this does not recursively calculate this
func CalculateSpawn(init int, days int) int {
	if days <= init {
		return 0
	}
	return int(math.Floor(float64(days-1+(7-init)) / 7))
}

func CalculateTotalSpawn(init int, days int) (totalSpawn uint64) {
	initSpawn := CalculateSpawn(init, days)
	totalSpawn = uint64(initSpawn)
	for i := 0; i < initSpawn; i++ {
		// calculate which day the fish was spawned on
		// birthday is when it ticks over to 6 again,
		// so we need to count down init days, then tick over
		// one extra
		// then we add how many iterations on
		birthday := init + 1 + i*7

		// firstSpawn is the first date this fish can produce spawn
		// which is 9 days after its birthday
		firstSpawn := birthday + 9

		// See if this first spawn happens or if we've run out of days
		if days-firstSpawn < 0 {
			continue
		}

		// now iterate again!
		totalSpawn += CalculateTotalSpawn(0, days-firstSpawn-1)
	}
	return totalSpawn
}

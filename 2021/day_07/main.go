package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	l := getLine("input.txt")
	input := splitAndToI(l)
	co := cheapestOutcomeOne(input)

	fmt.Println("cheapest outcome", co)

	fuel := calcFuel(co, input)
	fmt.Println("PART ONE: fuel", fuel)

	pos, fuel := cheapestOutcomeTwo(input)
	fmt.Println("PART Two: pos, fuel", pos, fuel)
}

// median
func cheapestOutcomeOne(input []int) int {
	sort.Ints(input)

	median := input[len(input)/2]

	return median
}

// mean
func cheapestOutcomeTwo(input []int) (int, int) {
	sort.Ints(input)
	sum := 0
	for _, n := range input {
		sum += n
	}

	mean := sum / len(input)
	fuel := calcFuelTwo(mean, input)

	// check above
	for {
		fuel2 := calcFuelTwo(mean+1, input)
		if fuel2 < fuel {
			fuel = fuel2
			mean += 1
		} else {
			break
		}
	}

	// check below
	for {
		fuel2 := calcFuelTwo(mean-1, input)
		if fuel2 < fuel {
			fuel = fuel2
			mean -= 1
		} else {
			break
		}
	}

	return mean, fuel
}

func calcFuel(pos int, crabs []int) (fuel int) {
	for _, crab := range crabs {
		fuel += abs(crab - pos)
	}
	return fuel
}

func calcFuelTwo(pos int, crabs []int) (fuel int) {
	for _, crab := range crabs {
		fuel += gaussSum(abs(crab - pos))
	}
	return fuel
}

func gaussSum(n int) int {
	return (n * (n + 1) / 2)
}

//helpers
func getLine(path string) string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	s.Scan()
	return s.Text()
}

func splitAndToI(line string) []int {
	s := strings.Split(line, ",")
	var p []int
	for _, inst := range s {
		p = append(p, toI(inst))
	}

	return p
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

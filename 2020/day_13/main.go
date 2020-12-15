package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"

	lines := getLines(file)

	ts := toI(lines[0])
	busses := parseBus(lines[1])
	nextBus := findNextBus(ts, busses)

	fmt.Printf("Part One: nextBus: %d, Wait time: %d. Answer: %d\n", nextBus, nextBus-(ts%nextBus), nextBus*(nextBus-(ts%nextBus)))

	bus2 := paseBusWithBlanks(lines[1])

	fmt.Println("Part 2: ", sequence(bus2))
}

func findNextBus(ts int, busses []int) int {
	// |n| n -  ( 939 % n) + 939

	min := 0
	minIdx := 0
	for i, b := range busses {
		ns := b - (ts % b) + ts
		if min == 0 || ns < min {
			min = ns
			minIdx = i
		}
	}

	return busses[minIdx]
}

func sequence(busses []int) int {
	timestamp, lcf := busses[0], busses[0]
	for i := 1; i < int(len(busses)); i++ {
		if busses[i] == -1 {
			continue
		}
		for (timestamp+i)%busses[i] != 0 {
			timestamp += lcf
		}

		lcf *= busses[i]
	}
	return timestamp
}

func parseBus(s string) []int {
	bb := []int{}
	b := strings.Split(s, ",")
	for _, x := range b {
		if x != "x" {
			bb = append(bb, toI(x))
		}
	}
	return bb
}

func paseBusWithBlanks(s string) []int {
	bb := []int{}
	b := strings.Split(s, ",")
	for _, x := range b {
		if x != "x" {
			bb = append(bb, toI(x))
		} else {
			bb = append(bb, -1)
		}
	}
	return bb
}

// boiler plate
func getLines(path string) []string {
	file, err := os.Open(path)

	var lines []string

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

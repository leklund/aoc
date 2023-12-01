package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Pair struct {
	l int
	r int
}

var lineRe = regexp.MustCompile(`^(\d+)-(\d+),(\d+)-(\d+)$`)

func main() {
	file := "input.txt"
	lines := getLines(file)

	fmt.Println("Part One: ", countContained(lines))
	fmt.Println("Part Two: ", countOverlap(lines))
}

func countContained(lines []string) int {
	c := 0

	for _, line := range lines {
		if contained(makePairs(line)) {
			c++
		}
	}
	return c
}

func countOverlap(lines []string) int {
	c := 0

	for _, line := range lines {
		if overlap(makePairs(line)) {
			c++
		}
	}
	return c
}

func makePairs(line string) (Pair, Pair) {
	m := lineRe.FindStringSubmatch(line)

	return Pair{toI(m[1]), toI(m[2])}, Pair{toI(m[3]), toI(m[4])}
}

func contained(p1, p2 Pair) bool {
	if (p1.l <= p2.l && p1.r >= p2.r) || (p2.l <= p1.l && p2.r >= p1.r) {
		return true
	}
	return false
}

func overlap(p1, p2 Pair) bool {
	if (p2.l >= p1.l && p2.l <= p1.r) || (p2.r >= p1.l && p2.r <= p1.r) {
		return true
	} else if (p1.l >= p2.l && p1.l <= p2.r) || (p1.r >= p2.l && p1.r <= p2.r) {
		return true
	}
	return false
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

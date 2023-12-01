package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file := "input.txt"

	lines := getLines(file)

	e := elves(lines)
	sort.Ints(e)

	p1 := e[len(e)-1]

	fmt.Println("PART ONE: ", p1)

	p2 := topThree(e)

	fmt.Println("PART ONE: ", p2)

}

func elves(input []string) []int {
	out := []int{}
	i := 0

	for _, entry := range input {
		if entry == "" {
			i++
			continue
		}
		if len(out) == i {
			out = append(out, toI(entry))
		} else {
			out[i] += toI(entry)
		}
	}
	return out
}

func topThree(e []int) int {
	sum := 0
	for i := len(e) - 1; i >= len(e)-3; i-- {
		sum += e[i]
	}
	return sum
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

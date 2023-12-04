package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var bag = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
}
var gr = regexp.MustCompile(`(\d+) (red|blue|green)(?:,|;|$)`)

func main() {
	file := "/Users/leklund/projects/aoc/2023/02/input.txt"
	games := getLines(file)

	fmt.Println(run(games))

}

func run(games []string) (int, int) {
	var v, p int
	for i, g := range games {
		if validGame(g, bag) {
			v += i + 1
		}

		p += power(g)
	}
	return v, p
}

func validGame(g string, b map[string]int) bool {
	match := gr.FindAllStringSubmatch(g, -1)

	for _, m := range match {
		if toI(m[1]) > b[m[2]] {
			return false
		}
	}
	return true
}

func power(g string) int {
	match := gr.FindAllStringSubmatch(g, -1)

	max := make(map[string]int)

	for _, m := range match {
		if toI(m[1]) > max[m[2]] {
			max[m[2]] = toI(m[1])
		}
	}

	return max["red"] * max["green"] * max["blue"]
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type P struct {
	x int
	y int
}

var r = regexp.MustCompile(`(\d+|[^.])`)
var digit = regexp.MustCompile(`\d+`)

func main() {
	file := "/Users/leklund/projects/aoc/2023/03/input.txt"
	lines := getLines(file)

	symbols, numbers, nm := analyzeSchematic(lines)

	fmt.Println("ONE: ", sumParts(symbols, numbers))

	gs := gears(symbols, numbers, nm)

	r := calcRatios(gs)

	fmt.Println("TWO ", r)

}

func sumParts(symbols map[P]string, numbers map[P]string) int {
	var sum int

	for loc, num := range numbers {
		if isPartNumber(num, loc, symbols) {
			sum += toI(num)
		}
	}
	return sum
}

func isPartNumber(num string, loc P, symbolMap map[P]string) bool {
	x := loc.x - 1
	xx := loc.x + len(num)
	y := loc.y - 1
	yy := y + 2

	for i := x; i <= xx; i++ {
		for j := y; j <= yy; j++ {
			if _, ok := symbolMap[P{i, j}]; ok {
				return true
			}
		}
	}
	return false
}

func calcRatios(gears [][]int) int {
	var r int
	for _, g := range gears {
		r += g[0] * g[1]
	}
	return r
}

func gears(symbols map[P]string, numbers map[P]string, numberMap map[P]P) [][]int {
	var g [][]int

	for loc, sym := range symbols {
		if sym != "*" {
			continue
		}
		parts := make(map[P]bool)
		for i := loc.x - 1; i <= loc.x+1; i++ {
			for j := loc.y - 1; j <= loc.y+1; j++ {
				if p, ok := numberMap[P{i, j}]; ok {
					parts[p] = true
				}
			}
		}
		if len(parts) == 2 {
			var pns []int
			for p := range parts {
				v := toI(numbers[p])
				pns = append(pns, v)
			}
			g = append(g, pns)
		}
	}
	return g
}

func analyzeSchematic(lines []string) (symbols map[P]string, numbers map[P]string, numberMap map[P]P) {
	symbols = make(map[P]string)
	numbers = make(map[P]string)
	numberMap = make(map[P]P)

	for y, l := range lines {
		mm := r.FindAllString(l, -1)
		xx := r.FindAllStringIndex(l, -1)

		for i, m := range mm {
			x := xx[i][0]
			p := P{x, y}
			if digit.MatchString(m) {
				numbers[p] = m
				for z := x; z < x+len(m); z++ {
					numberMap[P{z, y}] = p
				}

			} else {
				symbols[p] = m
			}
		}
	}
	return symbols, numbers, numberMap
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

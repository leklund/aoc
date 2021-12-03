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
	gamma, epsilon := PartOne(lines)

	fmt.Printf("PART ONE: gamma: %d, epsilon: %d\n", gamma, epsilon)
	fmt.Println("Power consumption:", gamma*epsilon)

	o2 := oxygenScrubber(lines)
	co2 := co2Scrubber(lines)

	fmt.Printf("Part Two: o2 %d, co2 %d, life support rating: %d\n", o2, co2, o2*co2)
}

func PartOne(input []string) (int, int) {
	g, e := mostCommon(input)

	gamma, _ := strconv.ParseUint(strings.Join(g, ""), 2, 32)
	epsilon, _ := strconv.ParseUint(strings.Join(e, ""), 2, 32)
	return int(gamma), int(epsilon)
}

func mostCommon(lines []string) ([]string, []string) {
	c := make([]int, len(lines[0]))

	for _, l := range lines {
		for i, n := range l {
			if n == '1' {
				c[i]++
			} else {
				c[i]--
			}
		}
	}

	g := make([]string, len(lines[0]))
	e := make([]string, len(lines[0]))
	for i, x := range c {
		if x > 0 {
			g[i] = "1"
			e[i] = "0"
		} else {
			g[i] = "0"
			e[i] = "1"
		}
	}
	return g, e
}

func oxygenScrubber(lines []string) int {
	o := make([]string, len(lines))
	for i := 0; i < len(lines[0]); i++ {
		if i == 0 {
			copy(o, lines)
		}

		o = filterMost(o, i)

		if len(o) == 1 {
			break

		}
	}
	x, _ := strconv.ParseUint(o[0], 2, 32)
	return int(x)
}

func co2Scrubber(lines []string) int {
	o := make([]string, len(lines))
	for i := 0; i < len(lines[0]); i++ {
		if i == 0 {
			copy(o, lines)
		}

		o = filterLeast(o, i)

		if len(o) == 1 {
			break

		}
	}
	x, _ := strconv.ParseUint(o[0], 2, 32)
	return int(x)
}

func filterMost(lines []string, i int) []string {
	var o []string

	g, _ := mostCommonByIndex(lines, i)

	for _, line := range lines {
		if line[i] == g {
			o = append(o, line)
		}
	}
	return o
}

func filterLeast(lines []string, i int) []string {
	var o []string

	_, e := mostCommonByIndex(lines, i)

	for _, line := range lines {
		if line[i] == e {
			o = append(o, line)
		}
	}
	return o
}

func mostCommonByIndex(lines []string, i int) (byte, byte) {
	var c int
	for _, l := range lines {
		if l[i] == '1' {
			c++
		} else {
			c--
		}
	}

	if c >= 0 {
		return '1', '0'

	} else {
		return '0', '1'
	}
}

// fancy binary mask BS
func getRates(g string) (int, int) {
	gamma, _ := strconv.ParseUint(g, 2, 32)

	mask := (1 << len(g)) - 1
	epsilon := int(gamma) ^ mask
	return int(gamma), int(epsilon)
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

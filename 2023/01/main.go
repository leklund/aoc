package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	file := "/Users/leklund/projects/aoc/2023/01/input.txt"
	lines := getLines(file)

	fmt.Println("Part One: ", One(lines))
	fmt.Println("Part Two: ", Two(lines))
}

func One(lines []string) int {
	digit := regexp.MustCompile(`\d`)
	var x int

	for _, line := range lines {
		m := digit.FindAllString(line, -1)
		x += toI(m[0] + m[len(m)-1])
	}
	return x
}

func Two(lines []string) int {
	var x int
	digits := []string{`\d`, "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	base := `(` + strings.Join(digits, "|") + `)`
	first := regexp.MustCompile(base)
	last := regexp.MustCompile(`.*` + base)

	for _, line := range lines {
		ma := first.FindString(line)
		mb := last.FindStringSubmatch(line)[1]

		a := slices.Index(digits, ma)
		if a > 0 {
			x += a * 10
		} else {
			x += toI(ma) * 10
		}
		b := slices.Index(digits, mb)
		if b > 0 {
			x += b
		} else {
			x += toI(mb)
		}
	}
	return x
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

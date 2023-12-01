package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
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
	dm := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	first := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine|\d)`)
	last := regexp.MustCompile(`.*(one|two|three|four|five|six|seven|eight|nine|\d)`)

	for _, line := range lines {
		ma := first.FindString(line)
		mb := last.FindStringSubmatch(line)[1]

		var a, b string
		a, ok := dm[ma]
		if !ok {
			a = ma
		}

		b, ok = dm[mb]
		if !ok {
			b = mb
		}

		x += toI(a + b)
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

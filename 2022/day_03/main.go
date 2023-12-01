package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file := "input.txt"
	lines := getLines(file)
	p := 0
	for _, line := range lines {
		p += priority(line)
	}
	fmt.Println("PART ONE: ", p)
	fmt.Println("PART TWO: ", partTwo(lines))
}

func half(s string) (string, string) {
	return s[0 : len(s)/2], s[len(s)/2:]
}

func common(s string) string {
	l, r := half(s)
	return intersect(strings.Split(l, ""), strings.Split(r, ""))
}

func priority(s string) int {
	return alpha[common(s)]
}

func partTwo(lines []string) int {
	p := 0

	for _, group := range groups(lines) {
		c := trisect(strings.Split(group[0], ""), strings.Split(group[1], ""), strings.Split(group[2], ""))
		p += alpha[c]
	}

	return p
}

func groups(lines []string) [][]string {
	var grouped [][]string
	var group []string

	for i, line := range lines {

		if i == 0 || i%3 != 0 {
			group = append(group, line)
		} else {
			grouped = append(grouped, group)
			group = []string{line}
		}
	}
	grouped = append(grouped, group)

	return grouped
}

func intersect[T comparable](a []T, b []T) T {
	var s T
	hash := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			s = v
			return s
		}
	}
	return s
}

func trisect[T comparable](a []T, b []T, c []T) T {
	var s T
	hash := make(map[T]struct{})
	hash2 := make(map[T]struct{})

	for _, v := range a {
		hash[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := hash[v]; ok {
			hash2[v] = struct{}{}
		}
	}

	for _, v := range c {
		if _, ok := hash2[v]; ok {
			s = v
			return s
		}
	}

	return s
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

var alpha = map[string]int{"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
	"A": 27,
	"B": 28,
	"C": 29,
	"D": 30,
	"E": 31,
	"F": 32,
	"G": 33,
	"H": 34,
	"I": 35,
	"J": 36,
	"K": 37,
	"L": 38,
	"M": 39,
	"N": 40,
	"O": 41,
	"P": 42,
	"Q": 43,
	"R": 44,
	"S": 45,
	"T": 46,
	"U": 47,
	"V": 48,
	"W": 49,
	"X": 50,
	"Y": 51,
	"Z": 52,
}

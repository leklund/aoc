package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"
	lines := getString(file)

	pw := ValidPasswords(lines)
	fmt.Printf("PART ONE: %d valid passwords\n", len(pw))

	pw2 := ValidPasswordsTwo(lines)
	fmt.Printf("PART TWO: %d valid passwords\n", len(pw2))
}

func ValidPasswords(lines string) []string {
	var valid []string

	reg := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

	matches := reg.FindAllStringSubmatch(lines, -1)
	for _, m := range matches {
		min, max, char, pw := toI(m[1]), toI(m[2]), m[3], m[4]

		count := strings.Count(pw, char)
		if count >= min && count <= max {
			valid = append(valid, pw)
		}
	}

	return valid
}

func ValidPasswordsTwo(lines string) []string {
	var valid []string

	reg := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

	matches := reg.FindAllSubmatch([]byte(lines), -1)
	for _, m := range matches {
		i, j, char, pw := toI(string(m[1]))-1, toI(string(m[2]))-1, m[3][0], m[4]

		// XOR
		if (pw[i] == char || pw[j] == char) && !(pw[i] == char && pw[j] == char) {
			valid = append(valid, string(pw))
		}
	}

	return valid
}

// boiler plate
func getString(path string) string {
	data, _ := ioutil.ReadFile(path)

	return string(data)
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

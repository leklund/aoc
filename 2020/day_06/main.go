package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	file := "input.txt"
	groups := getGroups(file)

	count1, count2 := 0, 0

	for _, g := range groups {
		count1 += getUnion(g)
		count2 += getIntersection(g)
	}

	fmt.Println("PART ONE: ", count1)
	fmt.Println("PART TWO: ", count2)
}

func getUnion(grp string) int {
	grp = strings.ReplaceAll(grp, "\n", "")

	m := make(map[rune]interface{})

	for _, char := range grp {
		m[char] = true
	}

	return len(m)
}

func getIntersection(grp string) int {
	peeps := strings.Split(grp, "\n")

	m := make(map[rune]int)

	for _, peep := range peeps {
		for _, char := range peep {
			m[char]++
		}
	}

	count := 0

	for _, x := range m {
		if x == len(peeps) {
			count++
		}
	}

	return count
}

// boiler plate
func getString(path string) string {
	data, _ := ioutil.ReadFile(path)

	return string(data)
}

func getGroups(path string) []string {
	file := getString(path)

	return strings.Split(file, "\n\n")
}

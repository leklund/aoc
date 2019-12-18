package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Path has both a slice path and a map
// this way I can iterate over one path in order an check for intersections in the second path in the map
type Path struct {
	arr  []string
	hash map[string]int
}

func main() {
	m := getLines("input.txt")

	orbitalMap := makeMap(m)

	orbitCount := ancestorCount(orbitalMap)
	fmt.Println("-- Part ONE (171213)---")

	fmt.Println(orbitCount)

	// part two

	youPath := makePath("YOU", orbitalMap)
	sanPath := makePath("SAN", orbitalMap)

	steps := findTransfers(youPath, sanPath)

	fmt.Println("-- Part TWO ---")
	fmt.Println(steps)
}

func makeMap(inputMap []string) map[string]string {
	orbitalMap := make(map[string]string)
	for _, pair := range inputMap {
		planets := strings.Split(pair, ")")

		orbitalMap[planets[1]] = planets[0]
	}

	return orbitalMap
}

func ancestorCount(orbitalMap map[string]string) int {
	aCounts := make(map[string]int)
	var count int
	for child, parent := range orbitalMap {
		if aCount, ok := aCounts[parent]; ok {
			count += aCount + 1
		} else {
			count += getCounts(child, orbitalMap, aCounts)

		}
	}
	return count
}

func getCounts(child string, orbitalMap map[string]string, aCounts map[string]int) int {
	if parent, ok := orbitalMap[child]; ok {
		if aCount, ok := aCounts[parent]; ok {
			return aCount + 1
		}
		return getCounts(parent, orbitalMap, aCounts) + 1
	}
	// COM has no orbits
	return 0
}

func makePath(start string, orbitalMap map[string]string) Path {
	path := Path{[]string{}, make(map[string]int)}
	for parent, i := orbitalMap[start], 0; parent != "COM"; i++ {
		path.arr = append(path.arr, parent)
		path.hash[parent] = i
		start = parent
		parent = orbitalMap[start]
	}
	return path
}

// returns numer of transfers to get from path1[0] to path2[0]
func findTransfers(path1, path2 Path) int {
	for step, planet := range path1.arr {
		if step2, ok := path2.hash[planet]; ok {
			return step + step2
		}
	}
	return 0
}

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

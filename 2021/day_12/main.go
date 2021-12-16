package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("DAY 12!")

	caveMap := parse(getLines("input.txt"))
	v := make(map[string]int)
	paths := FindPaths(START, caveMap, v, true)

	fmt.Println("ONE: ", len(paths))

	v2 := make(map[string]int)
	paths2 := FindPaths(START, caveMap, v2, false)
	fmt.Println("ONE: ", len(paths2))

}

const (
	START = "start"
	END   = "end"
	DASH  = "-"
)

func FindPaths(start string, caveMap map[string][]string, visitCount map[string]int, SmolVisitTwice bool) [][]string {
	var paths [][]string

	branches := caveMap[start]

	for _, cave := range branches {
		if cave == END {
			paths = append(paths, []string{END})
			continue
		}

		has2 := SmolVisitTwice
		if strings.ToLower(cave) == cave && visitCount[cave] >= 1 {
			if !has2 {
				has2 = true
			} else {
				continue
			}
		}

		vc2 := make(map[string]int)
		for c, x := range visitCount {
			vc2[c] = x
		}
		vc2[cave]++

		res := FindPaths(cave, caveMap, vc2, has2)

		for _, path := range res {
			newPath := []string{cave}
			newPath = append(newPath, path...)
			paths = append(paths, newPath)
		}
	}

	return paths
}

func parse(paths []string) map[string][]string {
	nodes := make(map[string][]string)

	for _, path := range paths {
		caves := strings.Split(path, DASH)
		a, b := caves[0], caves[1]

		if a == END {
			nodes[b] = append(nodes[b], END)
		} else if b == END {
			nodes[a] = append(nodes[a], END)
		} else if a == START {
			nodes[START] = append(nodes[START], b)
		} else if b == START {
			nodes[START] = append(nodes[START], a)
		} else {
			nodes[a] = append(nodes[a], b)
			nodes[b] = append(nodes[b], a)
		}
	}
	return nodes
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

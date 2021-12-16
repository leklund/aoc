package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func main() {
	lines := getLines("input.txt")

	var hmap [][]int
	for _, line := range lines {
		row := splitAndToI(line, "")
		hmap = append(hmap, row)
	}

	lowPoints := findLowPoints(hmap)
	sum := 0
	for _, p := range lowPoints {
		sum += hmap[p.y][p.x] + 1
	}

	fmt.Println("PART ONE: ", sum)

	var basins []int
	for _, p := range lowPoints {
		m := make(map[Point]bool)
		basins = append(basins, basin(p, hmap, m))
	}
	sort.Ints(basins)

	product := 1
	for _, v := range basins[len(basins)-3:] {
		product *= v
	}

	fmt.Println("PART TWO (recursion FTW): ", product)
}

func findLowPoints(hmap [][]int) []Point {
	var low []Point
	checked := make(map[Point]bool)

	for y, row := range hmap {
		for x, v := range row {
			if checked[Point{x, y}] {
				continue
			}
			lowpoint := true

			var adj []Point
			//up
			if y-1 >= 0 {
				adj = append(adj, Point{x, y - 1})
			}
			//down
			if y < len(hmap)-1 {
				adj = append(adj, Point{x, y + 1})
			}
			// left
			if x-1 >= 0 {
				adj = append(adj, Point{x - 1, y})
			}
			//right
			if x < len(row)-1 {

				adj = append(adj, Point{x + 1, y})
			}

			for _, p := range adj {
				if hmap[p.y][p.x] > v {
					checked[p] = true
				} else {
					lowpoint = false
				}
			}

			if lowpoint {
				low = append(low, Point{x, y})
			}
		}
	}

	return low

}

func basin(p Point, hmap [][]int, mapped map[Point]bool) int {
	size := 0
	var adj []Point
	//up
	if p.y-1 >= 0 {
		adj = append(adj, Point{p.x, p.y - 1})
	}
	//down
	if p.y < len(hmap)-1 {
		adj = append(adj, Point{p.x, p.y + 1})
	}
	// left
	if p.x-1 >= 0 {
		adj = append(adj, Point{p.x - 1, p.y})
	}
	//right
	if p.x < len(hmap[p.y])-1 {

		adj = append(adj, Point{p.x + 1, p.y})
	}

	for _, a := range adj {
		// have we mapped it already?
		if seen := mapped[a]; !seen {
			// is it max height??
			if hmap[a.y][a.x] != 9 {
				mapped[a] = true
				size += 1 + basin(a, hmap, mapped)
			}
		}
	}

	return size

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

func splitAndToI(line string, sep string) []int {
	s := strings.Split(line, sep)
	var p []int
	for _, inst := range s {
		p = append(p, toI(inst))
	}

	return p
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

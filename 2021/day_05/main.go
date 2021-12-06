package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Line struct {
	p1         Point
	p2         Point
	orthagonal bool
}

type Graph map[Point]int

var parseRegEx = regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)

func main() {
	lines := getLines("input.txt")
	graph := newGraph(lines, false)

	fmt.Println("PART ONE: points with more than 1 intersection: ", graph.intersections())

	graph2 := newGraph(lines, true)

	fmt.Println("PART TWO: points with more than 1 intersection: ", graph2.intersections())
}

func (g Graph) intersections() int {
	intersections := 0
	for _, intersect := range g {
		if intersect > 1 {
			intersections++
		}
	}
	return intersections
}

func newGraph(input []string, useDiag bool) Graph {
	g := make(Graph)
	for _, l := range input {
		line := parseLine(l)

		if !useDiag && !line.orthagonal {
			continue
		}

		if line.p1.x == line.p2.x {
			x := line.p1.x
			y := []int{line.p1.y, line.p2.y}
			sort.Ints(y)

			for i := y[0]; i <= y[1]; i++ {
				g[Point{x, i}]++
			}
		} else if line.p1.y == line.p2.y {
			y := line.p1.y
			x := []int{line.p1.x, line.p2.x}
			sort.Ints(x)

			for i := x[0]; i <= x[1]; i++ {
				g[Point{i, y}]++
			}
		} else {
			x1, y1, x2, y2 := line.p1.x, line.p1.y, line.p2.x, line.p2.y
			// negative slope
			if (x1 > x2 && y1 > y2) || (x1 < x2 && y1 < y2) {
				y := []int{y1, y2}
				x := []int{x1, x2}
				sort.Ints(y)
				sort.Ints(x)

				for i := 0; i <= (y[1] - y[0]); i++ {
					p := Point{x[0] + i, y[0] + i}

					g[p]++

				}
			} else {
				var xx, yy, size int
				if x1 < x2 {
					xx, yy = x1, y1
					size = x2 - x1
				} else {
					xx, yy = x2, y2
					size = x1 - x2
				}

				for i := 0; i <= size; i++ {
					x := xx + i
					y := yy - i
					p := Point{x, y}
					g[p]++
				}

			}
		}

	}
	return g
}

func parseLine(l string) Line {
	match := parseRegEx.FindAllStringSubmatch(l, -1)
	line := Line{
		p1: Point{x: toI(match[0][1]), y: toI(match[0][2])},
		p2: Point{x: toI(match[0][3]), y: toI(match[0][4])},
	}

	if line.p1.x == line.p2.x || line.p1.y == line.p2.y {
		line.orthagonal = true
	}
	return line
}

//helpers
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

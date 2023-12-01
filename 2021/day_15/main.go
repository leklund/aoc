package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

var dirs = []Point{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func main() {
	cave, dest := buildCave(getLines("input.txt"))

	path := navigate(cave, dest)
	risk := calcRisk(dest, path, cave)

	fmt.Println("PART ONE: ", risk)

	bigCave, dest2 := expandCave(cave, dest)

	path2 := navigate(bigCave, dest2)
	risk2 := calcRisk(dest2, path2, bigCave)

	fmt.Println("PART TWO: ", risk2)

}

func navigate(cave map[Point]int, dest Point) map[Point]Point {
	var pq PriorityQueue
	origin := Point{0, 0}
	est := heuristic(origin, dest)

	heap.Push(&pq, &Item{point: origin, priority: est})

	seen := make(map[Point]bool)
	riskMap := map[Point]int{origin: 0}
	path := make(map[Point]Point)

	i := 0
	for {
		// fmt.Println(i)
		// for _, item := range pq {
		// 	fmt.Println(item)
		// }
		current := heap.Pop(&pq)

		loc := current.(*Item).point

		// fmt.Println("loc: ", loc, "p: ", current.(*Item).priority)

		// we made it!
		if loc == dest {
			return path
		}

		seen[loc] = true
		currentRisk := riskMap[loc]

		for _, move := range loc.adj() {

			// out of range?
			if move.x < 0 || move.x > dest.x || move.y < 0 || move.y > dest.y {
				continue
			}

			// been there done that
			if seen[move] {
				continue
			}

			newRisk := currentRisk + cave[loc]

			// skip if more risky
			if riskMap[move] != 0 && newRisk > riskMap[move] {
				continue
			}

			riskMap[move] = newRisk
			path[move] = loc

			newEst := newRisk + heuristic(loc, dest)

			// if pq.Exists(move) {
			// 	pq.update(current.(*Item), loc, newEst)
			// } else {
			heap.Push(&pq, &Item{point: move, priority: newEst})
			// }

		}

		i++
	}
}

func calcRisk(dest Point, path map[Point]Point, cave map[Point]int) int {
	// walk backwards
	o := Point{0, 0}
	risk := cave[dest]

	prevPoint := path[dest]

	for prevPoint != o {
		risk += cave[prevPoint]
		prevPoint = path[prevPoint]
	}

	return risk

}

func heuristic(loc Point, dest Point) int {
	return abs(loc.x-dest.x) + abs(loc.y-dest.y)
}

func buildCave(lines []string) (map[Point]int, Point) {
	var mx, my int
	out := make(map[Point]int)
	for y, line := range lines {
		if y > my {
			my = y
		}
		chars := splitAndToI(line)
		for x, risk := range chars {
			if x > mx {
				mx = x
			}
			p := Point{x, y}
			out[p] = risk

		}

	}

	return out, Point{mx, my}
}

func expandCave(cave map[Point]int, br Point) (map[Point]int, Point) {
	var mx, my int
	out := cave
	for x := 0; x <= br.x; x++ {
		for y := 0; y <= br.y; y++ {
			srcRisk := cave[Point{x, y}]
			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					if i == 0 && j == 0 {
						continue
					}

					newRisk := (srcRisk + i + j) % 9
					if newRisk == 0 {
						newRisk = 9
					}

					nx := x + (i * (br.x + 1))
					ny := y + (j * (br.y + 1))
					out[Point{nx, ny}] = newRisk
				}
			}
		}
	}

	mx = br.x + (4 * (br.x + 1))
	my = br.y + (4 * (br.y + 1))

	return out, Point{mx, my}
}

func printCave(cave map[Point]int, dest Point) {
	for y := 0; y <= dest.x; y++ {
		for x := 0; x <= dest.x; x++ {
			fmt.Print(cave[Point{x, y}])
		}
		fmt.Print("\n")
	}
}

func (p Point) adj() []Point {
	points := make([]Point, 4)

	for _, d := range dirs {
		points = append(points, p.add(d))
	}
	return points
}

func (p Point) add(p2 Point) Point {
	return Point{p.x + p2.x, p.y + p2.y}
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

func splitAndToI(line string) []int {
	s := strings.Split(line, "")
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

func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}

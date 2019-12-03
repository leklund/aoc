package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type coord struct {
	x int
	y int
}

type pointSet map[coord]int

func main() {
	wires := getLines("input.txt")

	wire1 := strings.Split(wires[0], ",")
	wire2 := strings.Split(wires[1], ",")

	pointSetA := makePath(wire1)
	pointSetB := makePath(wire2)

	manhattan, steps := findIntersections(pointSetA, pointSetB)

	sort.Ints(manhattan)
	fmt.Println("___ PART ONE ___")
	fmt.Println(manhattan[0])

	sort.Ints(steps)
	fmt.Println("___ PART TWO ___")
	fmt.Println(steps[0])
}

func makePath(wire []string) pointSet {
	points := make(pointSet)

	var x, y, steps int
	for _, v := range wire {
		direction, distance := parse(v)

		for i := 0; i < distance; i++ {
			steps++
			switch direction {
			case 'U':
				y++
			case 'R':
				x++
			case 'D':
				y--
			case 'L':
				x--
			}
			p := coord{x: x, y: y}
			points[p] = steps
		}
	}

	return points
}

func findIntersections(path1, path2 pointSet) ([]int, []int) {
	var mds, steps []int

	for point, s1 := range path1 {
		if s2, ok := path2[point]; ok {
			mds = append(mds, abs(point.x)+abs(point.y))
			steps = append(steps, s1+s2)
		}
	}

	return mds, steps
}

func parse(instruction string) (byte, int) {
	d := instruction[0]
	l, _ := strconv.Atoi(instruction[1:])
	return d, l
}

// hacky
func abs(x int) int {
	y := x >> 63
	return (x ^ y) - y
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

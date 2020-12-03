package main

import (
	"bufio"
	"fmt"
	"os"
)

type Slope struct {
	dx int
	dy int
}

func main() {
	file := "input.txt"
	lines := getLines(file)

	treecount := GoSledding(lines, 3, 1)

	fmt.Printf("PART ONE: I hit %d trees\n", treecount)

	trees := 1
	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	for _, slope := range slopes {
		trees *= GoSledding(lines, slope.dx, slope.dy)
	}

	fmt.Printf("PART TWO: trees:  %d", trees)
}

func GoSledding(lines []string, dx int, dy int) int {
	var x, trees int
	bottom := len(lines) - 1
	width := len(lines[0])

	for y := dy; y < len(lines); y += dy {
		line := lines[y]
		x += dx

		if x >= width {
			x = x - width
		}

		if string(line[x]) == "#" {
			trees++
		}

		if y == bottom {
			fmt.Println("hit bottom")
			break
		}
	}

	return trees
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

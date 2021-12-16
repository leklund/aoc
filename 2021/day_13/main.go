package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

type Fold struct {
	axis string
	pos  int
}

type Paper struct {
	points map[Point]bool
	w      int
	h      int
}

var foldRe = regexp.MustCompile(`fold along (\w)=(\d+)`)

func main() {
	lines := getLines("input.txt")
	paper, instructions := initPaper(lines)

	paper.fold(instructions[0])

	fmt.Println("Part One", paper.pointCount())

	for _, fold := range instructions[1:] {
		paper.fold(fold)
	}
	paper.print()
}

func initPaper(lines []string) (*Paper, []Fold) {
	points := make(map[Point]bool)
	inst := []Fold{}
	var w int
	var h int

	for _, line := range lines {
		if line == "" {
			continue
		}
		if strings.Contains(line, "fold along") {
			m := foldRe.FindAllStringSubmatch(line, -1)
			inst = append(inst, Fold{axis: m[0][1], pos: toI(m[0][2])})
		} else {
			coords := splitAndToI(line)
			x := coords[0]
			y := coords[1]
			if x > w {
				w = x
			}
			if y > h {
				h = y
			}
			points[Point{x, y}] = true
		}
	}
	p := &Paper{
		points: points,
		w:      w + 1,
		h:      h + 1,
	}
	return p, inst
}

func (p *Paper) fold(f Fold) {
	x0 := 0
	y0 := 0
	if f.axis == "y" {
		y0 = f.pos + 1
	} else if f.axis == "x" {
		x0 = f.pos + 1
	}

	for x := x0; x < p.w; x++ {
		for y := y0; y < p.h; y++ {
			pt := Point{x, y}

			if p.points[pt] {
				p.points[pt] = false
				if f.axis == "y" {
					np := Point{x, p.h - 1 - y}
					p.points[np] = true
				} else if f.axis == "x" {
					np := Point{p.w - 1 - x, y}
					p.points[np] = true
				}
			}
		}
	}

	if f.axis == "y" {
		p.h = ((p.h + 1) / 2) - 1
	} else if f.axis == "x" {
		p.w = ((p.w + 1) / 2) - 1
	}
}

func (p *Paper) pointCount() int {
	sum := 0
	for _, v := range p.points {
		if v {
			sum++
		}
	}
	return sum
}

func (p *Paper) print() {
	fmt.Print(p.toS())
	fmt.Print("\n")
}

func (p *Paper) toS() string {
	s := ""
	for y := 0; y < p.h; y++ {
		for x := 0; x < p.w; x++ {
			if p.points[Point{x, y}] {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
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

func splitAndToI(line string) []int {
	s := strings.Split(line, ",")
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

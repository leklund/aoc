package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Octopus struct {
	pos       Point
	energy    int
	neighbors Octopi
}

type Octopi []*Octopus

type Tracker struct {
	step    int
	flashes int
	octs    Octopi
}

func main() {
	lines := getLines("input.txt")

	tracker := initTracker(lines)
	for i := 0; i < 100; i++ {
		tracker.run()

	}

	fmt.Println("PART ONE: ", tracker.flashes)

	tracker2 := initTracker(lines)
	step := tracker2.runUntilAll()

	fmt.Println("PART TWO: all flashing on step ", step)
}

func (t *Tracker) run() {
	t.step++

	//energy plus one
	for _, o := range t.octs {
		o.energy++
	}

	// FLASHERING
	t.checkFlash()

}

func (t *Tracker) runUntilAll() int {
	lastFlash := 0

	for {
		t.run()
		if t.flashes-lastFlash == 100 {
			return t.step
		} else {
			lastFlash = t.flashes
		}
	}
}

func (t *Tracker) checkFlash() {
	for _, o := range t.octs {
		if o.energy >= 10 {
			// OH SNAP WE FLASHED
			o.flash(t)
		}
	}
}

func (o *Octopus) flash(t *Tracker) {
	t.flashes++
	o.energy = 0
	// neighbors feeling the heat
	for _, n := range o.neighbors {
		if n.energy == 0 {
			// I ALREADY FLASHED MAN
			continue
		}
		n.energy++
		if n.energy >= 10 {
			n.flash(t)
		}
	}
}

func initTracker(lines []string) *Tracker {
	o := builder(lines)

	return &Tracker{
		octs: o,
	}

}

func builder(lines []string) Octopi {
	octos := Octopi{}
	omap := make(map[Point]*Octopus)
	for y, line := range lines {
		for x, e := range line {

			o := &Octopus{
				energy: toI(string(e)),
				pos:    Point{x, y},
			}
			octos = append(octos, o)
			omap[o.pos] = o
		}
	}

	// let's get neigborly
	for _, oct := range octos {
		oct.neighbors = findNeighbors(oct, omap)
	}

	return octos
}

func findNeighbors(o *Octopus, omap map[Point]*Octopus) Octopi {
	octos := Octopi{}

	ox, oy := o.pos.x, o.pos.y

	for x := ox - 1; x <= ox+1; x++ {
		if x < 0 || x > 9 {
			continue
		}
		for y := oy - 1; y <= oy+1; y++ {
			if y < 0 || y > 9 {
				continue
			}
			if x == ox && y == oy {
				continue
			}
			// hi nieghbor

			n := omap[Point{x, y}]

			octos = append(octos, n)

		}

	}
	return octos
}

func (octs Octopi) print() string {
	for i := 1; i <= 100; i++ {
		o := octs[i-1]
		fmt.Print(o.energy)
		if i%10 == 0 {
			fmt.Print("\n")
		}
	}
	return ""
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

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

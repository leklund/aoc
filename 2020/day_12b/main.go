package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Ferry struct {
	x       int
	y       int
	wpx     int
	wpy     int
	heading int
}

type Action struct {
	action byte
	value  int
}

var funcMap = map[byte]func(Action, *Ferry){
	'N': func(a Action, f *Ferry) { f.wpy += a.value },
	'S': func(a Action, f *Ferry) { f.wpy -= a.value },
	'E': func(a Action, f *Ferry) { f.wpx += a.value },
	'W': func(a Action, f *Ferry) { f.wpx -= a.value },
	'R': func(a Action, f *Ferry) { rotate(f, a) },
	'L': func(a Action, f *Ferry) { rotate(f, a) },
}

func main() {
	file := "../day_12/input.txt"

	lines := getLines(file)

	actionList := makeActions(lines)

	ferry := &Ferry{
		x:       0,
		y:       0,
		wpx:     10,
		wpy:     1,
		heading: 90,
	}

	ferry.navigate(actionList)

	fmt.Println("Manhattan Distance from origin:", abs(ferry.x)+abs(ferry.y))
	fmt.Println("Ferry x, Ferry y", ferry.x, ferry.y)
}

func (f *Ferry) navigate(actions []Action) {
	for _, a := range actions {
		switch a.action {
		case 'F':
			f.x += a.value * f.wpx
			f.y += a.value * f.wpy
		default:
			funcMap[a.action](a, f)
		}
		// fmt.Printf("action: %s%d\n", string(a.action), a.value)
		// fmt.Printf("%d: x,y: %d, %d, wpx, wpy: %d, %d\n", i, f.x, f.y, f.wpx, f.wpy)
	}
}

func rotate(f *Ferry, a Action) {
	v := a.value
	if a.action == 'L' {
		if a.value == 270 {
			v = 90
		} else if a.value == 90 {
			v = 270
		}
	}

	switch v {
	case 90:
		f.wpx, f.wpy = f.wpy, -f.wpx
	case 180:
		f.wpx, f.wpy = -f.wpx, -f.wpy
	case 270:
		f.wpx, f.wpy = -f.wpy, f.wpx
	}
}

func makeActions(lines []string) []Action {
	actions := []Action{}

	for _, l := range lines {
		a := Action{
			action: l[0],
			value:  toI(l[1:]),
		}

		actions = append(actions, a)
	}

	return actions
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

func abs(x int) int {
	y := x >> 63
	return (x ^ y) - y
}

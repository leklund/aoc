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
	heading int
}

type Action struct {
	action byte
	value  int
}

var compass = map[int]byte{
	0:   'N',
	90:  'E',
	180: 'S',
	270: 'W',
}

const C = 360

var funcMap = map[byte]func(Action, *Ferry){
	'N': func(a Action, f *Ferry) { f.y += a.value },
	'S': func(a Action, f *Ferry) { f.y -= a.value },
	'E': func(a Action, f *Ferry) { f.x += a.value },
	'W': func(a Action, f *Ferry) { f.x -= a.value },
	'R': func(a Action, f *Ferry) { h := f.heading + a.value; f.heading = ((h % C) + C) % C },
	'L': func(a Action, f *Ferry) { h := f.heading - a.value; f.heading = ((h % C) + C) % C },
}

func main() {
	file := "input.txt"

	lines := getLines(file)

	actionList := makeActions(lines)

	ferry := &Ferry{
		x:       0,
		y:       0,
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
			d, ok := compass[f.heading]
			if !ok {
				fmt.Println("compass fail. f heading", f.heading)
			}
			a2 := Action{d, a.value}
			if _, ok := funcMap[d]; !ok {
				fmt.Println("-------------- NOT OK", "d", string(d), "a.action", string(a.action), "a2.action", string(a2.action))
			}
			funcMap[d](a2, f)
		default:
			funcMap[a.action](a, f)
		}
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

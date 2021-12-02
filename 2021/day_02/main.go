package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Sub struct {
	depth int
	x     int
	aim   int
}

var matcher = regexp.MustCompile(`^(\w+) (\d+)`)

var cmdMap = map[string]func(int, *Sub){
	"forward": func(x int, s *Sub) { s.x += x },
	"up":      func(x int, s *Sub) { s.depth -= x },
	"down":    func(x int, s *Sub) { s.depth += x },
}

var cmdMap2 = map[string]func(int, *Sub){
	"forward": func(x int, s *Sub) { s.x += x; s.depth += x * s.aim },
	"up":      func(x int, s *Sub) { s.aim -= x },
	"down":    func(x int, s *Sub) { s.aim += x },
}

func main() {
	file := "input.txt"
	lines := getLines(file)

	sub := &Sub{}

	depth, x := Dive1(sub, lines)

	fmt.Println(depth, x)
	fmt.Println("Part One: ", depth*x)

	sub = &Sub{}
	depth, x = Dive2(sub, lines)
	fmt.Println("Part Two: ", depth*x)
}

func Dive1(sub *Sub, cmds []string) (int, int) {
	//var x, y int

	for _, cmd := range cmds {
		match := matcher.FindStringSubmatch(cmd)

		op := match[1]
		val := toI(match[2])

		cmdMap[op](val, sub)
	}
	return sub.depth, sub.x
}

func Dive2(sub *Sub, cmds []string) (int, int) {
	//var x, y int

	for _, cmd := range cmds {
		match := matcher.FindStringSubmatch(cmd)

		op := match[1]
		val := toI(match[2])

		cmdMap2[op](val, sub)
	}
	return sub.depth, sub.x
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

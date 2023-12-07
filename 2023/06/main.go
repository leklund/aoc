package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type race struct {
	t int
	d int
}

func main() {
	input := []race{
		{35, 212},
		{93, 2060},
		{73, 1201},
		{66, 1044},
	}
	in2 := race{35937366, 212206012011044}
	fmt.Println("ONE: ", one(input))

	fmt.Println("TWO: ", waysToWin(in2))
}

func one(races []race) int {
	ways := 1
	for _, race := range races {
		ways *= waysToWin(race)
	}
	return ways
}

func waysToWin(r race) (count int) {
	left := (r.d / r.t) + 1
	right := r.t - left

	for left*(r.t-left) <= r.d {
		left++
	}

	for right*(r.t-right) <= r.d {
		right--
	}
	//fmt.Printf("race t/d %d/%d, left/right %d/%d\n", r.t, r.d, left, right)

	return right - left + 1
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

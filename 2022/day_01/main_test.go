package main

import (
	"fmt"
	"sort"
	"testing"
)

var testCases = []struct {
	input []string
	part1 int
	part2 int
}{
	{
		[]string{"1000", "2000", "3000", "", "4000", "", "5000", "6000", "", "7000", "8000", "9000", "", "10000"},
		24000,
		45000,
	},
}

func TestElves(t *testing.T) {
	for _, testcase := range testCases {
		elves := elves(testcase.input)

		sort.Ints(elves)
		fmt.Println(elves)
		x := elves[len(elves)-1]
		if x != testcase.part1 {
			t.Errorf("1: expected %d, got %d", testcase.part1, x)
		}

		z := topThree(elves)

		if z != testcase.part2 {
			t.Errorf("2: expected %d, got %d", testcase.part2, z)
		}
	}
}

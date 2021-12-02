package main

import "testing"

var testCases = []struct {
	input []string
	part1 int
	part2 int
}{
	{
		[]string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
		150,
		900,
	},
}

func TestDive(t *testing.T) {
	for _, tc := range testCases {
		sub := &Sub{}
		x, y := Dive1(sub, tc.input)

		if x*y != tc.part1 {
			t.Errorf("1: expected %d, got %d", tc.part1, x*y)
		}
	}
}

func TestDive2(t *testing.T) {
	for _, tc := range testCases {
		sub := &Sub{}
		x, y := Dive2(sub, tc.input)

		if x*y != tc.part2 {
			t.Errorf("1: expected %d, got %d -- pos %d, depth %d", tc.part2, x*y, y, x)
		}
	}
}

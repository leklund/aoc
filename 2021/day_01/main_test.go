package main

import "testing"

var testCases = []struct {
	input []int
	part1 int
	part2 int
}{
	{
		[]int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
		7,
		5,
	},
}

func TestCountIncreases(t *testing.T) {
	for _, testCase := range testCases {
		x := CountIncreases(testCase.input)

		if x != testCase.part1 {
			t.Errorf("1: expected %d, got %d", testCase.part1, x)
		}
	}
}

func TestCountIncreaseSums(t *testing.T) {
	for _, testCase := range testCases {
		x := CountIncreaseSums(testCase.input)

		if x != testCase.part2 {
			t.Errorf("2: expected %d, got %d", testCase.part2, x)
		}
	}
}

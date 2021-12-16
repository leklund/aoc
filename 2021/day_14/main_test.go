package main

import (
	"testing"
)

var testCases = []struct {
	input string
	min   int
	max   int
	steps int
}{
	{
		"input_test.txt",
		161,
		1749,
		10,
	},
	{
		"input_test.txt",
		3849876073,
		2192039569602,
		40,
	},
}

func TestOne(t *testing.T) {
	for _, tc := range testCases {
		p, r := parseInput(tc.input)

		np := run(p, tc.steps, r)

		min, max := np.minMax()

		if min != tc.min {
			t.Errorf("expected %d, got %d", tc.min, min)
		}
		if max != tc.max {
			t.Errorf("expected %d, got %d", tc.max, max)
		}

	}
}

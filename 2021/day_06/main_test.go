package main

import (
	"testing"
)

var testCases = []struct {
	input []int
	days  int
	count int
}{
	{
		[]int{3, 4, 3, 1, 2},
		18,
		26,
	},
	{
		[]int{3, 4, 3, 1, 2},
		80,
		5934,
	},
	{
		[]int{3, 4, 3, 1, 2},
		256,
		26984457539,
	},
}

func TestLantern(t *testing.T) {
	for _, tc := range testCases {
		fish := fish(tc.input)
		total := Run(tc.days, fish)

		if total != tc.count {
			t.Errorf("Expected %d lantern fish, got %d", tc.count, total)
		}
	}
}

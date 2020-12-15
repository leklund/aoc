package main

import (
	"testing"
)

var testCases = []struct {
	ts      int
	busses  []int
	nextBus int
}{
	{
		939,
		[]int{7, 13, 59, 31, 19},
		59,
	},
	{
		939,
		[]int{7, 13, -1, -1, 59, -1, 31, 19},
		59,
	},
}

func TestNextBus(t *testing.T) {
	for _, tc := range testCases {
		nb := findNextBus(tc.ts, tc.busses)

		if nb != tc.nextBus {
			t.Errorf("expected %d got %d", tc.nextBus, nb)
		}
	}
}

func TestSequence(t *testing.T) {
	for _, tc := range testCases {
		seq := sequence(tc.busses)

		if seq != 1 {
			t.Errorf("asd %d", seq)
		}
	}
}

package main

import (
	"sort"
	"testing"
)

var testCases = []struct {
	adapters []int
	diff     int
	perm     int
}{
	{
		[]int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4},
		35,
		8,
	},
	{
		[]int{28, 33, 18, 42, 31, 14, 46, 20, 48, 47, 24, 23, 49, 45, 19, 38, 39, 11, 1, 32, 25, 35, 8, 17, 7, 9, 4, 2, 34, 10, 3},
		220,
		19208,
	},
}

func TestFindGaps(t *testing.T) {
	for _, tc := range testCases {
		sort.Ints(tc.adapters)
		gaps := findGaps(tc.adapters)

		if gaps[1]*gaps[3] != tc.diff {
			t.Errorf("expected %d, got %d -- %v", tc.diff, gaps[1]*gaps[3], gaps)
		}
	}
}

func TestPermute(t *testing.T) {
	for _, tc := range testCases {

		sort.Ints(tc.adapters)

		res := permute(tc.adapters)

		if res != tc.perm {
			t.Errorf("expected %d got %d", res, tc.perm)
		}
	}
}

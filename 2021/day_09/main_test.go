package main

import (
	"fmt"
	"sort"
	"testing"
)

var testCases = []struct {
	input [][]int
	risk  int
	prisk int
}{
	{
		[][]int{{2, 1, 9, 9, 9, 4, 3, 2, 1, 0},
			{3, 9, 8, 7, 8, 9, 4, 9, 2, 1},
			{9, 8, 5, 6, 7, 8, 9, 8, 9, 2},
			{8, 7, 6, 7, 8, 9, 6, 7, 8, 9},
			{9, 8, 9, 9, 9, 6, 5, 6, 7, 8},
		},
		15,
		1134,
	},
}

func TestFindLow(t *testing.T) {
	for _, tc := range testCases {
		lowpoints := findLowPoints(tc.input)
		fmt.Println(lowpoints)
		sum := 0
		for _, p := range lowpoints {
			sum += tc.input[p.y][p.x] + 1
		}

		if sum != tc.risk {
			t.Errorf("expected %d got %d", tc.risk, sum)
		}
	}
}

func TestBasing(t *testing.T) {
	for _, tc := range testCases {
		lowpoints := findLowPoints(tc.input)

		basins := []int{}

		for _, p := range lowpoints {
			m := make(map[Point]bool)
			bsize := basin(p, tc.input, m)
			basins = append(basins, bsize)
		}
		sort.Ints(basins)

		product := 1
		for _, v := range basins[len(basins)-3:] {
			product *= v
		}
		if product != tc.prisk {
			t.Errorf("Expected %d, got %d", tc.prisk, product)
		}
	}
}

package main

import (
	"testing"
)

var lines = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

var testCases = []struct {
	lines []string
	trees int
}{
	{
		lines,
		7,
	},
}

var testCases2 = []struct {
	lines  []string
	trees  int
	slopes []Slope
}{
	{
		lines,
		336,
		[]Slope{
			{1, 2},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		},
	},
}

func TestGoSledding(t *testing.T) {
	for _, tc := range testCases {
		trees := GoSledding(tc.lines, 3, 1)

		if trees != tc.trees {
			t.Errorf("expected %d, got %d ", tc.trees, trees)
		}
	}

}

func TestGoSleddingMult(t *testing.T) {
	for _, tc := range testCases2 {
		prod := 1
		for _, slope := range tc.slopes {
			trees := GoSledding(tc.lines, slope.dx, slope.dy)

			prod *= trees
		}
		if prod != tc.trees {
			t.Errorf("expected %d got %d", tc.trees, prod)
		}
	}

}

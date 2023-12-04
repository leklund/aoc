package main

import (
	"testing"
)

var testCases = []struct {
	schematic string
	sum       int
}{
	{
		"/Users/leklund/projects/aoc/2023/03/input_test.txt",
		4361,
	},
}

func TestSumParts(t *testing.T) {
	for _, testcase := range testCases {
		lines := getLines(testcase.schematic)

		sym, num, _ := analyzeSchematic(lines)

		if _, ok := sym[P{3, 1}]; !ok {
			t.Errorf("should be a symbol at 3,1")
		}
		if num[P{0, 0}] != "467" {
			t.Errorf("457 should be at pos 0,0")
		}

		sum := sumParts(sym, num)

		if sum != testcase.sum {
			t.Errorf("got sum %d, expected %d", sum, testcase.sum)
		}

	}
}

func Test_isPartNumber(t *testing.T) {
	//type args struct {
	//	num       string
	//	loc       P
	//	symbolMap map[P]bool
	//}
	tests := []struct {
		num  string
		loc  P
		want bool
	}{
		{
			"467",
			P{0, 0},
			true,
		},
		{
			"114",
			P{5, 0},
			false,
		},
		{
			"58",
			P{7, 5},
			false,
		},
		{
			"592",
			P{2, 6},
			true,
		},
	}

	for _, tt := range tests {
		lines := getLines("/Users/leklund/projects/aoc/2023/03/input_test.txt")

		sym, _, _ := analyzeSchematic(lines)
		if got := isPartNumber(tt.num, tt.loc, sym); got != tt.want {
			t.Errorf("isPartNumber() = %v, want %v", got, tt.want)
		}

	}
}

func TestGears(t *testing.T) {
	for _, testcase := range testCases {
		lines := getLines(testcase.schematic)

		sym, n, nm := analyzeSchematic(lines)

		g := gears(sym, n, nm)

		want := 467835
		if calcRatios(g) != want {
			t.Errorf("got %d, want %d", calcRatios(g), want)
		}
	}
}

package main

import (
	"testing"
)

var testCases = []struct {
	problem   string
	solution  int
	solution2 int
}{
	{
		"1 + 2 * 3 + 4 * 5 + 6",
		71,
		231,
	},
	{
		"2 * 3 + (4 * 5)",
		26,
		46,
	},
	{
		"1 + (2 * 3) + (4 * (5 + 6))",
		51,
		51,
	},
	{
		"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
		12240,
		669060,
	},
	{
		"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
		13632,
		23340,
	},
	{
		"(1 + (2) + 3)",
		6,
		6,
	},
}

func TestAccum(t *testing.T) {

	for _, tc := range testCases {
		chars := splitter.FindAllString(tc.problem, -1)

		ans := accum(0, chars, 0)

		if ans != tc.solution {
			t.Errorf("expected %d got %d", tc.solution, ans)
		}
	}
}

func TestAccum2(t *testing.T) {

	for _, tc := range testCases {
		chars := splitter.FindAllString(tc.problem, -1)

		ans := accum2(chars)

		if ans != tc.solution {
			t.Errorf("expected %d got %d", tc.solution, ans)
		}
	}
}

func TestAccum2Part2(t *testing.T) {

	for _, tc := range testCases {
		chars := splitter.FindAllString(tc.problem, -1)

		ans := accum2Part2(chars)

		if ans != tc.solution2 {
			t.Errorf("expected %d got %d", tc.solution2, ans)
		}
	}
}

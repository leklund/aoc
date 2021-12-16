package main

import (
	"testing"
)

var testCases = []struct {
	input string
	part1 int
}{
	{
		"input_test.txt",
		17,
	},
}

func TestFold(t *testing.T) {
	for _, tc := range testCases {
		lines := getLines(tc.input)

		paper, inst := initPaper(lines)
		paper.fold(inst[0])

		if paper.pointCount() != tc.part1 {
			t.Errorf("Expecting %d, got %d", tc.part1, paper.pointCount())
		}

	}
}

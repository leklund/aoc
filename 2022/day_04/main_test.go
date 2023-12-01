package main

import "testing"

var testCases1 = []struct {
	line      string
	contained bool
	overlap   bool
}{
	{
		"2-4,6-8", false, false,
	},
	{
		"2-4,6-8", false, false,
	},
	{
		"5-7,7-9", false, true,
	},
	{
		"2-8,3-7", true, true,
	},
	{
		"6-6,4-6", true, true,
	},
	{
		"2-6,4-8", false, true,
	},
	{
		"4-8,4-8", true, true,
	},
	{
		"3-7,2-8", true, true,
	},
}

func TestIt(t *testing.T) {
	for _, tc := range testCases1 {
		p1, p2 := makePairs(tc.line)

		got := contained(p1, p2)

		if got != tc.contained {
			t.Errorf("expected pairs isContained to be %v, got %v", tc.contained, got)
		}

		got = overlap(p1, p2)

		if got != tc.overlap {
			t.Errorf("expected pairs %v, %v, overlap to be %v, got %v", p1, p2, tc.overlap, got)
		}
	}
}

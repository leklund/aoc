package main

import "testing"

var testCases = []struct {
	input []int
	count int
	res   int
}{
	{
		[]int{0, 3, 6},
		10,
		0,
	},
	{
		[]int{0, 3, 6},
		5,
		3,
	},
	{
		[]int{1, 3, 2},
		2020,
		1,
	},
	{
		[]int{2, 1, 3},
		2020,
		10,
	},
	{
		[]int{0, 3, 6},
		30000000,
		175594,
	},
}

func TestGamePlay(t *testing.T) {
	for _, tc := range testCases {
		g := &Game{
			input:   tc.input,
			lastMap: make(map[int]*Num),
		}

		res := g.play(tc.count)

		if res != tc.res {
			t.Errorf("expected %d, got %d, %v", tc.res, res, g)
		}
	}
}

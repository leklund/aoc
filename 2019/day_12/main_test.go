package main

import (
	"testing"
)

var testCases = []struct {
	in     []*Moon
	out    []*Moon
	steps  int
	energy int
}{
	{
		[]*Moon{
			{Point{-1, 0, 2}, Point{0, 0, 0}},
			{Point{2, -10, -7}, Point{0, 0, 0}},
			{Point{4, -8, 8}, Point{0, 0, 0}},
			{Point{3, 5, -1}, Point{0, 0, 0}},
		},
		[]*Moon{
			{Point{2, -1, 1}, Point{3, -1, -1}},
			{Point{3, -7, -4}, Point{1, 3, 3}},
			{Point{1, -7, 5}, Point{-3, 1, -3}},
			{Point{2, 2, 0}, Point{-1, -3, 1}},
		},
		1,
		229,
	},
	{
		[]*Moon{
			{Point{-1, 0, 2}, Point{0, 0, 0}},
			{Point{2, -10, -7}, Point{0, 0, 0}},
			{Point{4, -8, 8}, Point{0, 0, 0}},
			{Point{3, 5, -1}, Point{0, 0, 0}},
		},
		[]*Moon{
			{Point{2, 1, -3}, Point{-3, -2, 1}},
			{Point{1, -8, 0}, Point{-1, 1, 3}},
			{Point{3, -6, 1}, Point{3, 2, -3}},
			{Point{2, 0, 4}, Point{1, -1, -1}},
		},
		10,
		179,
	},
	{
		[]*Moon{
			{Point{-8, -10, 0}, Point{0, 0, 0}},
			{Point{5, 5, 10}, Point{0, 0, 0}},
			{Point{2, -7, 3}, Point{0, 0, 0}},
			{Point{9, -8, -3}, Point{0, 0, 0}},
		},
		[]*Moon{
			{Point{8, -12, -9}, Point{-7, 3, 0}},
			{Point{13, 16, -3}, Point{3, -11, -5}},
			{Point{-29, -11, -1}, Point{-3, 7, 4}},
			{Point{16, -13, 23}, Point{7, 1, 1}},
		},
		100,
		1940,
	},
}

func TestStep(t *testing.T) {
	for _, testCase := range testCases {
		for i := 0; i < testCase.steps; i++ {
			step(testCase.in)
		}

		for i, moon := range testCase.in {
			ex := testCase.out[i]
			if moon.pos.x != ex.pos.x || moon.pos.y != ex.pos.y || moon.pos.z != ex.pos.z || moon.vel.x != ex.vel.x || moon.vel.y != ex.vel.y || moon.vel.z != ex.vel.z {

				t.Errorf("Error, moon mismatch. Got %v, expected %v. step count: %d", *moon, *ex, testCase.steps)
			}
		}

		energy := energy(testCase.in)

		if energy != testCase.energy {
			t.Errorf("Error, energy mismatch. Expected %d, got %d", energy, testCase.energy)
		}
	}
}

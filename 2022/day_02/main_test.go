package main

import "testing"

var testCases1 = []struct {
	p1    Pick
	p2    Pick
	score int
}{
	{
		Rock, Paper, 1,
	},
	{
		Paper, Paper, 5,
	},
	{
		Scissors, Paper, 9,
	},
	{
		Rock, Rock, 4,
	},
}

func TestShoot(t *testing.T) {
	for _, tc := range testCases1 {
		got := shoot(tc.p1, tc.p2)
		if got != tc.score {
			t.Errorf("shoot: expected %d, got %d", tc.score, got)
		}
	}
}

func TestOne(t *testing.T) {
	inp := []string{"A Y", "B X", "C Z"}

	score := partOne(inp)

	if score != 15 {
		t.Errorf("got %d, expecting %d", score, 15)
	}
}

func TestDos(t *testing.T) {
	inp := []string{"A Y", "B X", "C Z"}

	score := partTwo(inp)

	if score != 12 {
		t.Errorf("got %d, expecting %d", score, 12)
	}
}

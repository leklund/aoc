package main

import (
	"testing"
)

//Time:      7  15   30
//Distance:  9  40  200

var testCases = []struct {
	race race
	wins int
}{
	{
		race{7, 9},
		4,
	},
	{
		race{15, 40},
		8,
	},
	{
		race{30, 200},
		9,
	},
	{
		race{71530, 940200},
		71503,
	},
}

func TestRaceWinCount(t *testing.T) {
	for _, testcase := range testCases {
		count := waysToWin(testcase.race)

		if count != testcase.wins {
			t.Errorf("Got win count %d, expected %d", count, testcase.wins)
		}

	}
}

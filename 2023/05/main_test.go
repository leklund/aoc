package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

var testCases = []struct {
	seed int
	loc  int
}{
	{
		79,
		82,
	},
	{
		14,
		43,
	},
	{
		55,
		86,
	},
	{
		13,
		35,
	},
}

func TestParseInput(t *testing.T) {
	lines := getLines("/Users/leklund/projects/aoc/2023/05/input_test.txt")
	parseInput(lines)

	wantSeeds := []int{79, 14, 55, 13}

	if !cmp.Equal(wantSeeds, SeedsOne) {
		t.Errorf("got %v seeds, expected %v", SeedsOne, wantSeeds)
	}

	Dest(SeedToSoil, 55)
	if Dest(SeedToSoil, 55) != 57 || Dest(SeedToSoil, 98) != 50 || Dest(SeedToSoil, 123) != 123 {
		t.Errorf("seed 55 should equal 57, got %d, and seed 98 should equal 50, got %d, 123 fail %d", Dest(SeedToSoil, 55), Dest(SeedToSoil, 98), Dest(SeedToSoil, 123))
	}

}
func TestSeedLoc(t *testing.T) {
	lines := getLines("/Users/leklund/projects/aoc/2023/05/input_test.txt")
	parseInput(lines)

	if minLoc() != 35 {
		t.Errorf("got min loc %d expected %d", minLoc(), 35)
	}

	for _, testcase := range testCases {
		seedLoc := SeedToLocation(testcase.seed)
		if seedLoc != testcase.loc {
			t.Errorf("got seed:loc %d:%d, expected %d:%d", testcase.seed, seedLoc, testcase.seed, testcase.loc)
		}
	}
}

func TestSeedLocTwo(t *testing.T) {
	lines := getLines("/Users/leklund/projects/aoc/2023/05/input_test.txt")
	parseInput(lines)

	if minLocSmart() != 46 {
		t.Errorf("got min loc %d expected %d", minLoc(), 46)
	}

}

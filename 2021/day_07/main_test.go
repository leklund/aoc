package main

import (
	"sort"
	"testing"
)

var testCases = []struct {
	input []int
	cheap int
	fuel  int
}{
	{
		[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
		2,
		37,
	},
	{
		[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
		1,
		41,
	},
	{
		[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
		3,
		39,
	},
}
var testCases2 = []struct {
	input []int
	cheap int
	fuel  int
}{
	{
		[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
		5,
		168,
	},
	{
		[]int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
		2,
		206,
	},
}

func TestCheapest(t *testing.T) {
	for i, tc := range testCases {
		if i > 0 {
			continue
		}
		cheapest := cheapestOutcomeOne(tc.input)
		if cheapest != tc.cheap {
			t.Errorf("Expected %d  got %d", tc.cheap, cheapest)
		}
	}
}

func TestFuel(t *testing.T) {
	for _, tc := range testCases {

		fuel := calcFuel(tc.cheap, tc.input)
		if fuel != tc.fuel {
			t.Errorf("Expected %d  got %d", tc.fuel, fuel)
		}
	}
}

func TestCheapestTwo(t *testing.T) {
	for i, tc := range testCases2 {
		if i > 0 {
			continue
		}
		pos, fuel := cheapestOutcomeTwo(tc.input)
		if pos != tc.cheap {
			t.Errorf("Expected %d  got %d", tc.cheap, pos)
		}
		if fuel != tc.fuel {
			t.Errorf("Expected %d  got %d", tc.fuel, fuel)
		}
	}
}

func TestFuelPartTwo(t *testing.T) {
	for _, tc := range testCases2 {
		sort.Ints(tc.input)
		fuel := calcFuelTwo(tc.cheap, tc.input)
		if fuel != tc.fuel {
			t.Errorf("Expected %d  got %d", tc.fuel, fuel)
		}
	}
}

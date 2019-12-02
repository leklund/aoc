package main

import "testing"

var testCases = []struct {
	mod  int
	fuel int
}{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

var testCases2 = []struct {
	mod  int
	fuel int
}{
	{14, 2},
	{1969, 966},
	{100756, 50346},
}

func TestCalcFuel(t *testing.T) {
	for _, testCase := range testCases {
		result := calcFuel(testCase.mod)
		if result != testCase.fuel {
			t.Errorf("Error, expected %d, got %d", testCase.fuel, result)
		}
	}
}

func TestCalcFuelInclusive(t *testing.T) {
	for _, testCase := range testCases2 {
		result := calcFuelInclusive(testCase.mod)
		if result != testCase.fuel {
			t.Errorf("Error, expected %d, got %d", testCase.fuel, result)
		}
	}
}

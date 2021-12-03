package main

import "testing"

var testCases1 = []struct {
	input   []string
	gamma   int
	epsilon int
	o2      int
	co2     int
}{
	{
		[]string{"00100",
			"11110",
			"10110",
			"10111",
			"10101",
			"01111",
			"00111",
			"11100",
			"10000",
			"11001",
			"00010",
			"01010",
		},
		22,
		9,
		23,
		10,
	},
}

func TestSomething(t *testing.T) {
	for _, tc := range testCases1 {
		g, e := PartOne(tc.input)
		if g != tc.gamma {
			t.Errorf("1: expected %d, got %d", tc.gamma, g)
		}

		if e != tc.epsilon {
			t.Errorf("1: expected %d, got %d", tc.epsilon, e)
		}
	}
}

func TestOxygen(t *testing.T) {
	for _, tc := range testCases1 {
		o := oxygenScrubber(tc.input)
		if o != tc.o2 {
			t.Errorf("1: expected %d, got %d", tc.o2, o)
		}
	}
}

func TestCO2(t *testing.T) {
	for _, tc := range testCases1 {
		co2 := co2Scrubber(tc.input)
		if co2 != tc.co2 {

			t.Errorf("1: expected %d, got %d", tc.co2, co2)
		}
	}
}

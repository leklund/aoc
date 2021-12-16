package main

import (
	"testing"
)

var testCases = []struct {
	input string
	ten   int
	hundy int
	all   int
}{
	{
		"input_test.txt",
		204,
		1656,
		195,
	},
}

func TestLantern(t *testing.T) {
	for _, tc := range testCases {
		lines := getLines(tc.input)

		tracker := initTracker(lines)

		for i := 0; i < 10; i++ {
			tracker.run()

		}

		if tracker.flashes != tc.ten {
			t.Errorf("Expected %d, got %d", tc.ten, tracker.flashes)
		}

		for i := 0; i < 90; i++ {
			tracker.run()

		}
		if tracker.flashes != tc.hundy {
			t.Errorf("Expected %d, got %d", tc.hundy, tracker.flashes)
		}
	}
}

func TestAllFlash(t *testing.T) {
	for _, tc := range testCases {
		lines := getLines(tc.input)

		tracker := initTracker(lines)

		step := tracker.runUntilAll()

		if step != tc.all {
			t.Errorf("expected %d, got %d", tc.all, step)
		}
	}
}

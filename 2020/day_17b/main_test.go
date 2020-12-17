package main

import (
	"strings"
	"testing"
)

var testCases = []struct {
	input       string
	activeCount int
}{
	{
		`.#.
..#
###`,
		848,
	},
}

func TestInit(t *testing.T) {
	for _, tc := range testCases {
		c := initCube(strings.Split(tc.input, "\n"))

		if c.xmin != -1 || c.xmax != 3 || !c.state[Point{2, 1, 0, 0}] {
			t.Errorf("initial padding failed %v", c.state)
		}
	}
}

func TestRun(t *testing.T) {
	for _, tc := range testCases {
		c := initCube(strings.Split(tc.input, "\n"))

		c.run(6)
		if c.active() != tc.activeCount {
			t.Errorf("expected %d, got %d", tc.activeCount, c.active())
		}
	}
}

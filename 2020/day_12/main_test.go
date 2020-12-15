package main

import (
	"strings"
	"testing"
)

var testCases = []struct {
	lines string
	md    int
}{
	{
		`F10
N3
F7
R90
F11`, 25,
	},
}

func TestNavigate(t *testing.T) {
	for _, tc := range testCases {
		lines := strings.Split(tc.lines, "\n")
		actionList := makeActions(lines)

		ferry := Ferry{x: 0, y: 0, heading: 90}

		ferry.navigate(actionList)

		if ferry.x+ferry.y != tc.md {
			t.Errorf("expecting %d, got %d -- %v", tc.md, ferry.x+ferry.y, ferry)
		}
	}
}

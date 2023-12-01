package main

import "testing"

var testCases1 = []struct {
	line      string
	leftSide  string
	rightSide string
	shared    string
	score     int
}{
	{
		"vJrwpWtwJgWrhcsFMMfFFhFp", "vJrwpWtwJgWr", "hcsFMMfFFhFp", "p", 16,
	},
	{
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "jqHRNqRjqzjGDLGL", "rsFMfFZSrLrFZsSL", "L", 38,
	},
	{
		"PmmdzqPrVvPwwTWBwg", "PmmdzqPrV", "vPwwTWBwg", "P", 42,
	},
	{
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "wMqvLMZHhHMvwLH", "jbvcjnnSBnvTQFn", "v", 22,
	},
}

func TestHalf(t *testing.T) {
	for _, tc := range testCases1 {
		l, r := half(tc.line)
		if l != tc.leftSide || r != tc.rightSide {
			t.Errorf("expected %s and %s, got %s and %s", tc.leftSide, tc.rightSide, l, r)
		}
	}
}

func TestCommon(t *testing.T) {
	for _, tc := range testCases1 {
		c := common(tc.line)
		if c != tc.shared {
			t.Errorf("expected %s, got %s", tc.shared, c)
		}
	}
}
func TestPriority(t *testing.T) {
	for _, tc := range testCases1 {
		p := priority(tc.line)
		if p != tc.score {
			t.Errorf("expected %d, got %d", tc.score, p)
		}
	}
}

func TestTwo(t *testing.T) {
	lines := []string{"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw"}

	p := partTwo(lines)
	if p != 70 {
		t.Errorf("p2 -- got %d, expected 70", p)
	}

}

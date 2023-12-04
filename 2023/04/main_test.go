package main

import (
	"testing"
)

var testCases = []struct {
	card  string
	val   int
	count int
}{
	{
		"Card 1: 41 48 83 86 17 |  83 86  6 31 17  9 48 53",
		8,
		4,
	},
	{
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		2,
		2,
	},
	{
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		2,
		2,
	},
	{
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		1,
		1,
	},
	{
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		0,
		0,
	},
	{
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
		0,
		0,
	},
}

func TestCardVal(t *testing.T) {
	for _, testcase := range testCases {
		winners, have := parseLine(testcase.card)

		v := cardVal(winners, have)

		if v != testcase.val {
			t.Errorf("Got %d, expected %d", v, testcase.val)
		}

		c := cardCount(winners, have)
		if c != testcase.count {
			t.Errorf("Got %d, expected %d", c, testcase.count)
		}

	}
}

func TestAllCards(t *testing.T) {
	allCards := []string{}

	for _, testcase := range testCases {
		allCards = append(allCards, testcase.card)

	}

	cc := ScratchOffCount(allCards)

	if cc != 30 {
		t.Errorf("got card count of %d, expecting %d", cc, 30)
	}
}

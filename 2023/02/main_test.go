package main

import (
	"testing"
)

var testCases = []struct {
	game  string
	valid bool
	power int
}{
	{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		true,
		48,
	},
	{
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		true,
		12,
	},
	{
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		false,
		1560,
	},
	{
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		false,
		630,
	},
	{
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 greenn",
		true,
		36,
	},
}

func TestValid(t *testing.T) {
	for _, testcase := range testCases {
		if validGame(testcase.game, bag) != testcase.valid {
			t.Errorf("game `%s` should should be valid: %v, got %v", testcase.game, testcase.valid, !testcase.valid)
		}
	}
}

func TestPower(t *testing.T) {
	for _, testcase := range testCases {
		if power(testcase.game) != testcase.power {
			t.Errorf("game `%s` should should have power: %d, got %d", testcase.game, testcase.power, power(testcase.game))
		}
	}
}

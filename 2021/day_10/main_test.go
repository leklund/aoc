package main

import (
	"testing"
)

var testCases2 = []struct {
	lines []string
	score int
}{
	{
		[]string{
			"[({(<(())[]>[[{[]{<()<>>",
			"[(()[<>])]({[<{<<[]>>(",
			"{([(<{}[<>[]}>{[]{[(<()>",
			"(((({<>}<{<{<>}{[]{[]{}",
			"[[<[([]))<([[{}[[()]]]",
			"[{[{({}]{}}([{[{{{}}([]",
			"{<[[]]>}<{[{[{[]{()[[[]",
			"[<(<(<(<{}))><([]([]()",
			"<{([([[(<>()){}]>(<<{{",
			"<{([{{}}[<[[[<>{}]]]>[]]",
		},
		26397,
	},
}

var testCases = []struct {
	line  string
	ok    bool
	exp   rune
	unexp rune
}{
	{
		"{([(<{}[<>[]}>{[]{[(<()>",
		false,
		']',
		'}',
	},
	{
		"[[<[([]))<([[{}[[()]]]",
		false,
		']',
		')',
	},
	{
		"[{[{({}]{}}([{[{{{}}([]",
		false,
		')',
		']',
	},
	{
		"[<(<(<(<{}))><([]([]()",
		false,
		'>',
		')',
	},
	{
		"<{([([[(<>()){}]>(<<{{",
		false,
		']',
		'>',
	},
}
var testCases3 = []struct {
	line  string
	score int
}{
	{
		"[({(<(())[]>[[{[]{<()<>>",
		288957,
	},
	{
		"<{([{{}}[<[[[<>{}]]]>[]]",
		294,
	},
}

func TestParseLine(t *testing.T) {
	for _, tc := range testCases {
		ok, e, u := parseLine(tc.line)
		if ok != tc.ok {
			t.Errorf("expected %v got %v", tc.ok, ok)
		}
		if e != tc.exp {
			t.Errorf("expected char: expected %c, got %c", tc.exp, e)
		}
		if u != tc.unexp {
			t.Errorf("unexpected char: expected %c, got %c", tc.unexp, u)
		}
	}
}

func TestCheckSyntax(t *testing.T) {
	for _, tc := range testCases2 {
		e := CheckSyntax(tc.lines)
		if e != tc.score {
			t.Errorf("expected %d, got %d", tc.score, e)
		}

	}
}

func TestCompletion(t *testing.T) {
	for _, tc := range testCases3 {
		score := scoreCompletionString(tc.line)

		if score != tc.score {
			t.Errorf("expected %d, got %d", tc.score, score)
		}
	}
}

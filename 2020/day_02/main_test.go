package main

import "testing"

var testCases = []struct {
	lines  string
	valid  []string
	valid2 []string
}{
	{
		"1-3 a: abcde\n1-3 b: cdefg\n2-9 c: ccccccccc",
		[]string{"1-3 a: abcde", "2-9 c: ccccccccc"},
		[]string{"1-3 a: abcde"},
	},
}

func TestValidPasswords(t *testing.T) {
	for _, tc := range testCases {
		v := ValidPasswords(tc.lines)

		if len(v) != len(tc.valid) {
			t.Errorf("expected %d, got %d from %v", len(tc.valid), len(v), v)
		}
	}
}

func TestValidPasswordsTwo(t *testing.T) {
	for _, tc := range testCases {
		v := ValidPasswordsTwo(tc.lines)

		if len(v) != len(tc.valid2) {
			t.Errorf("expected %d, got %d from %v", len(tc.valid2), len(v), v)
		}
	}
}

package main

import "testing"

var testCases = []struct {
	line string
	pass Pass
}{
	{"FBFBBFFRLR", Pass{44, 5}},
	{"BFFFBBFRRR", Pass{70, 7}},
	{"FFFBBBFRRR", Pass{14, 7}},
	{"BBFFBBFRLL", Pass{102, 4}},
}

func TestMakePass(t *testing.T) {
	for _, tc := range testCases {
		res := makePass(tc.line)

		if res.col != tc.pass.col && res.row != tc.pass.row {
			t.Errorf("err expected %v got %v", tc.pass, res)
		}
	}
}

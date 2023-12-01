package main

import (
	"testing"
)

var testCases = []struct {
	input string
	r1    int
	r2    int
}{
	{
		"input_test.txt",
		40,
		315,
	},
}

func TestOne(t *testing.T) {
	for _, tc := range testCases {
		lines := getLines(tc.input)

		cave, dest := buildCave(lines)

		// fmt.Println(dest)

		r := navigate(cave, dest)
		//fmt.Println(r)

		risk := calcRisk(dest, r, cave)
		// fmt.Println(risk)

		if risk != tc.r1 {
			t.Errorf("expected %d got %d", tc.r1, risk)
		}

		bigCave, dest := expandCave(cave, dest)

		r2 := navigate(bigCave, dest)
		risk2 := calcRisk(dest, r2, bigCave)

		if risk2 != tc.r2 {
			t.Errorf("expected %d got %d", tc.r2, risk2)
		}
	}
}

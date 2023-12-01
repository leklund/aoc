package main

import (
	"testing"
)

var testCases = []struct {
	input []string
	part1 int
	part2 int
}{
	{
		[]string{"1abc2", "pqr3stu8vwx", "a1b2c3d4e5f", "treb7uchet"},
		142,
		-1,
	},
	{
		[]string{"two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"},
		-1,
		281,
	},
}

func TestCalibrate(t *testing.T) {
	for _, testcase := range testCases {
		if testcase.part1 > 0 {
			res := One(testcase.input)

			if testcase.part1 > 0 && res != testcase.part1 {
				t.Errorf("1: expected %d, got %d", testcase.part1, res)
			}
		}

		if testcase.part2 > 0 {
			res2 := Two(testcase.input)
			if testcase.part2 > 0 && res2 != testcase.part2 {
				t.Errorf("2: expected %d, got %d", testcase.part2, res2)
			}
		}
	}
}

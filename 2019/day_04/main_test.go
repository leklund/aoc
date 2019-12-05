package main

import (
	"testing"
)

var sliceTest = []struct {
	code  int
	array [6]int
}{
	{123456, [6]int{1, 2, 3, 4, 5, 6}},
	{666666, [6]int{6, 6, 6, 6, 6, 6}},
}

var validPassTest = []struct {
	code  int
	valid bool
}{
	{111111, true},
	{223450, false},
	{123789, false},
	{122345, true},
	{111123, true},
}

var validPassTestTwo = []struct {
	code  int
	valid bool
}{
	{112233, true},
	{123444, false},
	{111122, true},
}

func TestSlicer(t *testing.T) {
	for _, testCase := range sliceTest {
		out := slicer(testCase.code)

		if out != testCase.array {
			t.Errorf("Error, expected %d, got %d ", testCase.array, out)
		}
	}

}

func TestValidPass(t *testing.T) {
	for _, testCase := range validPassTest {
		isValid := validPass(slicer(testCase.code))

		if isValid != testCase.valid {
			t.Errorf("Error, expected %t, got %t for code %d", testCase.valid, isValid, testCase.code)
		}
	}
}

func TestValidPassTwo(t *testing.T) {
	for _, testCase := range validPassTestTwo {
		isValid := validPassTwo(slicer(testCase.code))

		if isValid != testCase.valid {
			t.Errorf("Error, expected %t, got %t for code %d", testCase.valid, isValid, testCase.code)
		}
	}
}

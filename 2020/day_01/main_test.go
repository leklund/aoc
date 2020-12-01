package main

import "testing"

var testCases = []struct {
	er    []int
	prod2 int
	prod3 int
}{
	{
		[]int{1721, 979, 366, 299, 675, 1456},
		514579,
		241861950,
	},
}

func TestFind2020(t *testing.T) {
	for _, testCase := range testCases {
		x, y := Find2020(testCase.er)

		if x*y != testCase.prod2 {
			t.Errorf("expected %d, got %d", testCase.prod2, x*y)
		}
	}
}

func TestFindThree2020(t *testing.T) {
	for _, testCase := range testCases {
		x, y, z := FindThree2020(testCase.er)

		if x*y*z != testCase.prod3 {
			t.Errorf("expected %d, got %d", testCase.prod3, x*y*z)
		}
	}
}

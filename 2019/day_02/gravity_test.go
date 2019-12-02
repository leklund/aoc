package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	in  []int
	out []int
}{
	{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
	{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

func TestRun(t *testing.T) {
	for _, testCase := range testCases {
		run(testCase.in)
		result := testCase.in
		for i, v := range result {
			if v != testCase.out[i] {
				fmt.Println("got: ", result)
				fmt.Println("exp: ", testCase.out)
				t.Errorf("Error, expected %d, got %d at index %d", testCase.out[i], v, i)

			}
		}
	}
}

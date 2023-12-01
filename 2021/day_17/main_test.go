package main

import (
	"fmt"
	"testing"
)

var testCases = []struct {
	target   Box
	velocity Vector
}{
	{
		Box{Point{20, -5}, Point{30, -10}},
		Vector{6, 9},
	},
}

func TestOne(t *testing.T) {
	for _, tc := range testCases {
		v := findMaxY(tc.target)
		fmt.Println(v)
	}
}

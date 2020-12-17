package main

import (
	"fmt"
	"strings"
	"testing"
)

var testCases = []struct {
	input     string
	errorRate int
}{
	{
		`class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`,
		71,
	},
}

func TestErrorRate(t *testing.T) {
	for _, tc := range testCases {
		lines := strings.Split(tc.input, "\n")

		ranges, _, tickets := parse(lines)
		rate := errorRate(tickets, ranges)

		if rate != tc.errorRate {
			t.Errorf("expected %d, got %d", tc.errorRate, rate)
		}
	}
}

func TestValid(t *testing.T) {
	for _, tc := range testCases {
		lines := strings.Split(tc.input, "\n")

		ranges, _, tickets := parse(lines)
		v := valid(tickets, ranges)

		if v[0].fields[2] != 47 {
			t.Errorf("invalid valid %v", v[0].fields)
		}
	}
}

func TestIndexer(t *testing.T) {
	for _, tc := range testCases {
		lines := strings.Split(tc.input, "\n")

		ranges, _, tickets := parse(lines)
		fmt.Println(ranges)
		v := valid(tickets, ranges)

		fmt.Println("----------------")
		for _, vv := range v {
			fmt.Println(vv)
		}
		fmt.Println("----------------")
		mapping := indexer(v, ranges)

		t.Error(mapping)

	}
}

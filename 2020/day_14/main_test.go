package main

import (
	"strings"
	"testing"
)

var testCases1 = []struct {
	program string
	res     int
}{
	{
		`mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0`,
		165,
	},
}
var testCases2 = []struct {
	program string
	res     int
}{
	{
		`mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`,
		208,
	},
}

func TestDecodeV1(t *testing.T) {
	for _, tc := range testCases1 {
		lines := strings.Split(tc.program, "\n")

		mem := decodeV1(lines)

		tot := 0
		for _, x := range mem {
			tot += x
		}
		if tot != tc.res {

			t.Errorf("expected %d got %d", tc.res, tot)
		}
	}
}

func TestDecodeV2(t *testing.T) {
	for _, tc := range testCases2 {
		lines := strings.Split(tc.program, "\n")
		mem := decodeV2(lines)
		tot := 0
		for _, x := range mem {
			tot += x
		}
		if tot != tc.res {

			t.Errorf("expected %d got %d", tc.res, tot)
		}

	}

}

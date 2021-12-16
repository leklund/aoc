package main

import (
	"testing"
)

var testCases = []struct {
	input  [][]string
	count  int
	digits []int
}{
	{
		[][]string{{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
			{"fcgedb", "cgb", "dgebacf", "gc"},
			{"cg", "cg", "fdcagb", "cbg"},
			{"efabcd", "cedba", "gadfec", "cb"},
			{"gecf", "egdcabf", "bgf", "bfgea"},
			{"gebdcfa", "ecba", "ca", "fadegcb"},
			{"cefg", "dcbef", "fcge", "gbcadfe"},
			{"ed", "bcgafe", "cdgba", "cbgef"},
			{"gbdfcae", "bgc", "cg", "cgb"},
			{"fgae", "cfgab", "fg", "bagce"},
		},
		26,
		[]int{8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315},
	},
}
var testCases2 = []struct {
	input  []string
	count  int
	digits []int
}{
	{
		[]string{"acedgfb", "cdfbe", "gcdfa", "fbcad", "dab", "cefabd", "cdfgeb", "eafb", "cagedb", "ab"},
		26,
		[]int{8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315},
	},
	{
		[]string{"edbfga", "begcd", "cbg", "gc", "gcadebf", "fbgde", "acbgfd", "abcde", "gfcbed", "gfec"},
		26,
		[]int{8394, 9781, 1197, 9361, 4873, 8418, 4548, 1625, 8717, 4315},
	},
}

func TestOne(t *testing.T) {
	for _, tc := range testCases {

		count := countUniqs(tc.input)
		if count != tc.count {
			t.Errorf("Expected %d  got %d", tc.count, count)
		}
	}
}

// func TestDigitBilder(t *testing.T) {
// 	base := map[string]string{
// 		"acedgfb": "8",
// 		"cdfbe":   "5",
// 		"gcdfa":   "2",
// 		"fbcad":   "3",
// 		"dab":     "7",
// 		"cefabd":  "9",
// 		"cdfgeb":  "6",
// 		"eafb":    "4",
// 		"cagedb":  "0",
// 		"ab":      "1",
// 	}
// 	dm := digitMap(base)
// 	for _, tc := range testCases {
// 		for i, entry := range tc.input {
// 			digit := buildNumber(entry, dm)

// 			if digit != tc.digits[i] {
// 				t.Errorf("Exepcted %d got %d for entry %v", tc.digits[i], digit, entry)
// 			}
// 		}

// 	}
// }

func TestDecoder(t *testing.T) {
	for _, tc := range testCases2 {

		decoder(tc.input)

	}
}

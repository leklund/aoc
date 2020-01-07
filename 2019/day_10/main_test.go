package main

import "testing"

var testCases = []struct {
	path  string
	best  Point
	count int
}{
	{"test1.txt", Point{x: 5, y: 8}, 33},
	{"test2.txt", Point{x: 1, y: 2}, 35},
	{"test3.txt", Point{x: 6, y: 3}, 41},
	{"test4.txt", Point{x: 11, y: 13}, 210},
}

func TestMakeAsteroidMap(t *testing.T) {
	for _, testCase := range testCases {
		a := readFile(testCase.path)
		m := makeAsteroidMap(a)

		if !m[testCase.best] {
			t.Errorf("Error - expected point %v to exist", testCase.best)
		}

	}
}

func TestScan(t *testing.T) {
	for _, testCase := range testCases {
		a := readFile(testCase.path)
		m := makeAsteroidMap(a)

		best, count, _ := scanForBest(m)

		if best != testCase.best {
			t.Errorf("Error - expected best to be %v got %v", testCase.best, best)
		}

		if count != testCase.count {
			t.Errorf("Error - expected count to be %d got %d", testCase.count, count)
		}
	}
}

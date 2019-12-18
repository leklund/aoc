package main

import (
	"fmt"
	"reflect"
	"testing"
)

var mapTests = []struct {
	in    []string
	out   map[string]string
	count int
	steps int
}{
	{
		[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"},
		map[string]string{
			"B": "COM",
			"C": "B",
			"D": "C",
			"E": "D",
			"F": "E",
			"G": "B",
			"H": "G",
			"I": "D",
			"J": "E",
			"K": "J",
			"L": "K",
		},
		42,
		0,
	},
	{
		[]string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN"},
		map[string]string{
			"B":   "COM",
			"C":   "B",
			"D":   "C",
			"E":   "D",
			"F":   "E",
			"G":   "B",
			"H":   "G",
			"I":   "D",
			"J":   "E",
			"K":   "J",
			"L":   "K",
			"YOU": "K",
			"SAN": "I",
		},
		54,
		4,
	},
}

func TestMakeMap(t *testing.T) {
	for _, testCase := range mapTests {
		m := makeMap(testCase.in)

		if !reflect.DeepEqual(testCase.out, m) {
			t.Errorf("Error, expected %v, got %v", testCase.out, m)
		}
	}
}

func TestAncestorCount(t *testing.T) {
	for _, testCase := range mapTests {
		m := makeMap(testCase.in)

		count := ancestorCount(m)

		if count != testCase.count {
			t.Errorf("Error, expected count %d, got %d", testCase.count, count)
		}
	}
}

func TestIntersections(t *testing.T) {
	testCase := mapTests[1]
	m := makeMap(testCase.in)

	youPath := makePath("YOU", m)
	santaPath := makePath("SAN", m)

	fmt.Println(youPath, santaPath)

	steps := findTransfers(youPath, santaPath)

	if steps != testCase.steps {
		t.Errorf("Error, expected steps %d, got %d", testCase.steps, steps)
	}
}

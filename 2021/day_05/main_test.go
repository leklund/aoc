package main

import (
	"testing"
)

var testCase = struct {
	input []string
	p1    int
	p2    int
}{
	[]string{
		"0,9 -> 5,9",
		"8,0 -> 0,8",
		"9,4 -> 3,4",
		"2,2 -> 2,1",
		"7,0 -> 7,4",
		"6,4 -> 2,0",
		"0,9 -> 2,9",
		"3,4 -> 1,4",
		"0,0 -> 8,8",
		"5,5 -> 8,2",
	},
	5,
	12,
}

func TestIntersections(t *testing.T) {
	g := newGraph(testCase.input, false)

	if g.intersections() != testCase.p1 {
		t.Errorf("expected %d got %d", testCase.p1, g.intersections())
	}

}

func TestNewGraph(t *testing.T) {
	graph := newGraph(testCase.input, false)

	i := graph[Point{7, 0}]

	if i != 1 {
		t.Errorf("expected 1 got %d", i)
	}
	if graph[Point{0, 9}] != 2 {
		t.Errorf("expected 2 got %d", graph[Point{0, 9}])
	}
	for j := 3; j <= 9; j++ {
		p := Point{j, 4}

		if _, ok := graph[p]; !ok {
			t.Errorf("expecting point to be graphed at %v", p)
		}
	}
}

func TestNewGraphDiagonals(t *testing.T) {
	graph := newGraph(testCase.input, true)

	for j := 0; j <= 8; j++ {
		p := Point{j, j}

		if _, ok := graph[p]; !ok {
			t.Errorf("expecting point to be graphed at %v", p)
		}
	}

	if graph.intersections() != testCase.p2 {
		t.Errorf("expected %d got %d", testCase.p2, graph.intersections())
	}
}

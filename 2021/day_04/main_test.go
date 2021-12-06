package main

import (
	"testing"
)

var testCases1 = []struct {
	input []string
	sets  [][]int
}{
	{
		[]string{"22 13 17 11  0",
			"	8  2 23  4 24",
			"21  9 14 16  7",
			"	6 10  3 18  5",
			"	1 12 20 15 19",
		},
		[][]int{{22, 13, 17, 11, 0}, {8, 2, 23, 4, 24}, {21, 9, 14, 16, 7}, {6, 10, 3, 18, 5}, {1, 12, 20, 15, 19},
			{22, 8, 21, 6, 1}, {13, 2, 9, 10, 12}, {17, 23, 14, 3, 20}, {11, 4, 16, 18, 15}, {0, 24, 7, 5, 19}},
	},
}

var testCases2 = []struct {
	boards [][]string
	winner []string
	nums   []int
	score  int
}{
	{
		[][]string{{"22 13 17 11  0",
			"	8  2 23  4 24",
			"21  9 14 16  7",
			"	6 10  3 18  5",
			"	1 12 20 15 19",
		},
			{" 3 15  0  2 22",
				"	9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
			},
			{"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				" 2  0 12  3  7",
			},
		},
		[]string{"14 21 17 24  4",
			"10 16 15  9 19",
			"18  8 23 26 20",
			"22 11 13  6  5",
			" 2  0 12  3  7",
		},
		[]int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
		4512,
	},
}

var testCases3 = []struct {
	boards [][]string
	winner []string
	nums   []int
	score  int
}{
	{
		[][]string{{"22 13 17 11  0",
			"	8  2 23  4 24",
			"21  9 14 16  7",
			"	6 10  3 18  5",
			"	1 12 20 15 19",
		},
			{" 3 15  0  2 22",
				"	9 18 13 17  5",
				"19  8  7 25 23",
				"20 11 10 24  4",
				"14 21 16 12  6",
			},
			{"14 21 17 24  4",
				"10 16 15  9 19",
				"18  8 23 26 20",
				"22 11 13  6  5",
				" 2  0 12  3  7",
			},
		},
		[]string{" 3 15  0  2 22",
			"	9 18 13 17  5",
			"19  8  7 25 23",
			"20 11 10 24  4",
			"14 21 16 12  6",
		},
		[]int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
		1924,
	},
}

func TestParser(t *testing.T) {
	for _, tc := range testCases1 {
		b := NewCard(tc.input)

		b.parseBoard()

		for i, set := range b.sets {
			for j, x := range tc.sets[i] {
				if _, ok := set.nums[x]; !ok {
					t.Errorf("expected %d, got %d in row position %d, %d", x, tc.sets[i][j], i, j)
				}
			}

		}
	}
}

func TestPlay(t *testing.T) {
	for _, tc := range testCases2 {
		expected := NewCard(tc.winner)

		cards := Cards{}
		for _, b := range tc.boards {
			cards = append(cards, NewCard(b))
		}

		winner := cards.Play(tc.nums)
		winner.score()

		for i, line := range winner.input {
			if expected.input[i] != line {
				t.Errorf("expected %s to match %s", expected.input[i], line)
			}
		}

		if winner.finalScore*winner.winningNum != tc.score {
			t.Errorf("expected scrore %d, got %d", tc.score, winner.finalScore*winner.winningNum)
		}
	}
}

func TestPlayUntilLast(t *testing.T) {
	for _, tc := range testCases3 {
		expected := NewCard(tc.winner)

		cards := Cards{}
		for _, b := range tc.boards {
			cards = append(cards, NewCard(b))
		}

		winner := cards.PlayUntilLast(tc.nums)
		winner.score()

		for i, line := range winner.input {
			if expected.input[i] != line {
				t.Errorf("expected %s to match %s", expected.input[i], line)
			}
		}

		if winner.finalScore*winner.winningNum != tc.score {
			t.Errorf("expected scrore %d, got %d", tc.score, winner.finalScore*winner.winningNum)
		}
	}
}

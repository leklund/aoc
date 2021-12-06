package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Set struct {
	i        int
	hitCount int
	nums     map[int]bool
}

type BingoCard struct {
	input      []string
	sets       map[int]*Set
	winner     bool
	winningNum int
	finalScore int
}

type Cards []*BingoCard

func main() {
	file := "input.txt"
	groups := getGroups(file)
	firstline, groups := groups[0], groups[1:]

	numbers := splitAndToIComma(firstline)

	cards := Cards{}

	for _, s := range groups {
		input := strings.Split(s, "\n")
		cards = append(cards, NewCard(input))
	}

	// Part One
	// 	winner := cards.Play(numbers)
	// 	winner.score()

	// 	fmt.Println("PART ONE: BINGO!!", winner.finalScore, winner.winningNum, winner.finalScore*winner.winningNum)

	// Part Two
	winner := cards.PlayUntilLast(numbers)
	winner.score()

	fmt.Println("PART TWO: BINGO!!", winner.finalScore, winner.winningNum, winner.finalScore*winner.winningNum)
}

func (cards Cards) Play(numbers []int) *BingoCard {
	for _, n := range numbers {
		for _, c := range cards {
			if c.markCard(n) {
				//BINGO
				c.winningNum = n
				return c
			}
		}
	}
	return &BingoCard{}
}

func (cards Cards) PlayUntilLast(numbers []int) *BingoCard {
	cardCount := len(cards)
	winnerCount := 0
	for _, n := range numbers {
		for _, c := range cards {
			// skip the card if it's a winner
			if c.winner {
				continue
			}

			if c.markCard(n) {
				//BINGO
				c.winningNum = n
				c.winner = true
				winnerCount++

				if winnerCount == cardCount {
					return c
				}
			}
		}
	}
	return &BingoCard{}
}

func (b *BingoCard) score() {
	var s int
	for i := 0; i < 5; i++ {
		set := b.sets[i]
		for num, found := range set.nums {
			if !found {
				s += num
			}
		}
	}
	b.finalScore = s
}

func (b *BingoCard) markCard(n int) bool {
	for i := 0; i < 10; i++ {
		set := b.sets[i]
		if _, ok := set.nums[n]; ok {
			set.hitCount++
			set.nums[n] = true
		}
		if set.hitCount == 5 {
			// BINGO
			return true
		}
	}
	return false
}

func NewCard(s []string) *BingoCard {
	b := &BingoCard{
		input: s,
	}
	b.parseBoard()
	return b
}

func (b *BingoCard) parseBoard() {
	b.sets = make(map[int]*Set)
	rows := make([][]int, 5)
	cols := make([][]int, 5)

	for i := 0; i < 5; i++ {
		cols[i] = make([]int, 5)
	}

	for i, line := range b.input {
		line = strings.TrimSpace(line)
		rows[i] = splitAndToI(line)
		for j, x := range rows[i] {
			cols[j][i] = x
		}
	}

	for i, row := range rows {
		s := &Set{
			i:    i,
			nums: map[int]bool{},
		}
		for _, x := range row {
			s.nums[x] = false
		}
		b.sets[i] = s
	}

	for i, col := range cols {
		s := &Set{
			i:    i + 5,
			nums: map[int]bool{},
		}
		for _, x := range col {
			s.nums[x] = false
		}
		b.sets[i+5] = s
	}

}

// helpers
func getString(path string) string {
	data, _ := ioutil.ReadFile(path)

	return string(data)
}

func getGroups(path string) []string {
	file := getString(path)

	return strings.Split(file, "\n\n")
}

func splitAndToI(line string) []int {
	s := strings.Fields(line)
	var p []int
	for _, inst := range s {
		p = append(p, toI(inst))
	}

	return p
}

func splitAndToIComma(line string) []int {
	s := strings.Split(line, ",")
	var p []int
	for _, inst := range s {
		p = append(p, toI(inst))
	}

	return p
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

var cm = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"J": 11,
	"T": 10,
}

type Hand struct {
	cardlist string
	cards    []int
	Bid      int
	rank     int
	jokers   bool
}

func (h *Hand) Rank() int {
	if h.rank == 0 {
		h.SetRank()
	}
	return h.rank
}

func (h *Hand) Cards() []int {
	if len(h.cards) != 5 {
		h.cards = cardToSlice(h.cardlist, h.jokers)
	}
	return h.cards
}

func (h *Hand) SetRank() {
	counts := make(map[int]int)

	for _, c := range h.Cards() {
		counts[c]++
	}

	var rank int
	if h.jokers {
		jokers := counts[1]
		delete(counts, 1)
		if len(counts) == 0 {
			h.rank = 25
			return
		}

		allCounts := maps.Values(counts)

		slices.Sort(allCounts)
		slices.Reverse(allCounts)

		allCounts[0] += jokers

		for _, x := range allCounts {
			rank += x * x
		}
	} else {

		for _, x := range counts {
			rank += x * x
		}
	}
	h.rank = rank
}

type Hands []*Hand

func (h Hands) Len() int {
	return len(h)
}

func (h Hands) Less(i, j int) bool {
	if h[i].Rank() == h[j].Rank() {
		for k, x := range h[i].Cards() {
			y := h[j].Cards()[k]
			if x == y {
				continue
			} else {
				return x < y
			}
		}
	}
	return h[i].Rank() < h[j].Rank()
}

func (h Hands) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Hands) Sort() {
	sort.Sort(h)
}

func main() {
	input := "/Users/leklund/projects/aoc/2023/07/input.txt"
	lines := getLines(input)

	handsOne := makeHands(lines, false)

	fmt.Println("ONE: ", Winnings(handsOne))

	handsTwo := makeHands(lines, true)

	fmt.Println("Two: ", Winnings(handsTwo))
}

func Winnings(hands Hands) int {
	hands.Sort()

	out := 0

	for i, hand := range hands {
		out += (i + 1) * hand.Bid
	}
	return out
}

func makeHands(input []string, jokers bool) Hands {
	var h []*Hand

	for _, in := range input {
		parts := strings.Split(in, " ")
		cardlist := parts[0]
		bid := toI(parts[1])

		h = append(h, &Hand{
			cardlist: cardlist,
			Bid:      bid,
			jokers:   jokers,
		})
	}

	return h
}

func cardToSlice(cards string, jokers bool) []int {
	var cs []int

	for _, char := range cards {
		if x, ok := cm[string(char)]; ok {
			if jokers && char == 'J' {
				x = 1
			}
			cs = append(cs, x)
		} else {
			cs = append(cs, toI(string(char)))
		}
	}

	return cs
}

// boiler plate
func getLines(path string) []string {
	file, err := os.Open(path)

	var lines []string

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file := "/Users/leklund/projects/aoc/2023/04/input.txt"
	cards := getLines(file)

	fmt.Println("ONE: ", ScratchOffVal(cards))

	fmt.Println("TWO: ", ScratchOffCount(cards))

}

func cardCount(winners map[int]bool, have []int) int {
	var val int

	for _, x := range have {
		if winners[x] {
			val++
		}
	}
	return val
}
func ScratchOffCount(cc []string) int {
	allCards := make(map[int][]string)
	for i, c := range cc {
		allCards[i+1] = []string{c}
	}
	count := 0

	for CardIndex := 1; CardIndex <= len(allCards); CardIndex++ {
		cards := allCards[CardIndex]
		count += len(cards)

		winCount := cardCount(parseLine(cards[0]))

		for k := 0; k < len(cards); k++ {
			for j := 0; j < winCount; j++ {
				newIndex := CardIndex + j + 1
				allCards[newIndex] = append(allCards[newIndex], allCards[newIndex][0])
			}
		}
	}
	return count
}

func ScratchOffVal(cards []string) int {
	var totes int
	for _, card := range cards {
		totes += ScratchVal(card)
	}
	return totes
}

func ScratchVal(card string) int {
	return cardVal(parseLine(card))
}

func cardVal(winners map[int]bool, have []int) int {
	var val int

	for _, x := range have {
		if winners[x] {
			if val == 0 {
				val = 1
			} else {
				val *= 2
			}
		}
	}
	return val
}

func parseLine(line string) (map[int]bool, []int) {
	//Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	trimmer := regexp.MustCompile(`Card\s+\d+:\s+`)
	splitter := regexp.MustCompile(`\s+`)
	pipeSplitter := regexp.MustCompile(`\s+\|\s+`)
	line = trimmer.ReplaceAllString(line, "")

	parts := pipeSplitter.Split(line, -1)

	w := splitter.Split(parts[0], -1)
	winners := make(map[int]bool)

	for _, x := range w {
		winners[toI(x)] = true
	}

	var have []int
	for _, z := range splitter.Split(parts[1], -1) {
		have = append(have, toI(z))
	}

	return winners, have
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
		fmt.Printf("ERR -%s- %v\n", s, err)
		panic(err)
	}

	return i
}

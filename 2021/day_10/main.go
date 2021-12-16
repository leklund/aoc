package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var L = map[rune]rune{
	'{': '}',
	'[': ']',
	'(': ')',
	'<': '>',
}

var ErrorVal = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var Val = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func main() {
	lines := getLines("input.txt")
	fmt.Println("PART ONE: ", CheckSyntax(lines))

	inc := incompleteLines(lines)

	score := []int{}

	for _, i := range inc {
		score = append(score, scoreCompletionString(i))
	}

	sort.Ints(score)
	mid := score[(len(score)-1)/2]
	fmt.Println("PART TWO: ", mid)

}

func CheckSyntax(lines []string) int {
	errorScore := 0

	for _, line := range lines {
		if ok, _, u := parseLine(line); !ok {
			errorScore += ErrorVal[u]
		}
	}

	return errorScore
}

func incompleteLines(lines []string) []string {
	var inc []string

	for _, l := range lines {
		if ok, _, _ := parseLine(l); ok {
			inc = append(inc, l)
		}
	}
	return inc
}

func scoreCompletionString(line string) int {
	c := completionRunes(line)
	score := 0
	for _, r := range c {
		score = (score * 5) + Val[r]
	}
	return score

}

func completionRunes(line string) []rune {
	var completers []rune
	var chars []rune
	var x rune
	for _, r := range line {
		// fmt.Println(string(r))
		if _, ok := L[r]; ok {
			// push it real good
			chars = append(chars, r)
		} else {
			// pop
			chars = chars[:len(chars)-1]
		}
	}

	for {
		if len(chars) > 0 {
			// pop n lock it
			//z, x = x[len(x)-1], x[:len(x)-1]
			x, chars = chars[len(chars)-1], chars[:len(chars)-1]
			completers = append(completers, L[x])
		} else {
			break
		}
	}
	return completers

}

func parseLine(line string) (ok bool, expected rune, unexpected rune) {
	var chars []rune
	var x rune
	for _, r := range line {
		// fmt.Println(string(r))
		if _, ok := L[r]; ok {
			// push
			chars = append(chars, r)
		} else {
			// pop
			x, chars = chars[len(chars)-1], chars[:len(chars)-1]
			if L[x] != r {
				// fmt.Println("next char", string(r), "popped", string(x), "L map (expected)", string(L[x]))
				return false, L[x], r
			}
		}
	}
	return true, 0, 0
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

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

var ruleRE = regexp.MustCompile(`(\w\w) -> (\w)`)

type Pairs map[string]int

func main() {
	p, r := parseInput("input.txt")

	p = run(p, 10, r)

	min, max := p.minMax()

	fmt.Println("PART ONE", max-min)

	p2, r := parseInput("input.txt")
	p2 = run(p2, 40, r)
	min, max = p2.minMax()

	fmt.Println("PART ONE", max-min)
}

func run(p Pairs, steps int, rules map[string][]string) Pairs {
	newPairs := p
	for i := 0; i < steps; i++ {
		newPairs = step(newPairs, rules)
	}
	return newPairs
}

func step(p Pairs, rules map[string][]string) Pairs {
	newPairs := make(Pairs)
	for pair, count := range p {
		if newP, ok := rules[pair]; !ok {
			fmt.Println("NO RULE FOR ", pair)
		} else {
			newPairs[newP[0]] += count
			newPairs[newP[1]] += count
		}
	}
	return newPairs
}

func (p Pairs) countByElement() map[string]int {
	m := make(map[rune]int)
	out := make(map[string]int)
	for pair, count := range p {
		for _, r := range pair {
			m[r] += count
		}
	}

	for r, x := range m {
		out[string(r)] = (x + 1) / 2
	}

	return out
}

func (p Pairs) minMax() (int, int) {
	counted := p.countByElement()
	counts := []int{}

	for _, count := range counted {
		counts = append(counts, count)
	}

	sort.Ints(counts)

	return counts[0], counts[len(counts)-1]
}

func parseInput(file string) (pairs Pairs, rules map[string][]string) {
	rules = make(map[string][]string)
	pairs = make(Pairs)
	lines := getLines(file)

	templateString := lines[0]

	for i, r := range templateString {
		if i == len(templateString)-1 {
			break
		}
		pair := string(r) + string(templateString[i+1])
		pairs[pair]++
	}

	for _, rule := range lines[2:] {

		m := ruleRE.FindAllStringSubmatch(rule, -1)
		in := m[0][1]
		out := []string{
			string(in[0]) + m[0][2],
			m[0][2] + string(in[1]),
		}
		rules[in] = out
	}
	return pairs, rules
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

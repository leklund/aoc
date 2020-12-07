package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Contents struct {
	color string
	count int
}

func main() {
	file := "input.txt"
	lines := getLines(file)

	tree, tree2 := parse(lines)

	uniqs := make(map[string]interface{})
	uniqs = bagCount("shiny gold", tree, uniqs)

	fmt.Println("ONE:", len(uniqs))

	bagCount := contentsBagCount("shiny gold", tree2) - 1
	fmt.Println("TWO:", bagCount)

}

func bagCount(color string, tree map[string][]string, uniqs map[string]interface{}) map[string]interface{} {
	if parents, ok := tree[color]; ok {
		for _, p := range parents {
			uniqs[p] = struct{}{}
			uniqs = bagCount(p, tree, uniqs)
		}
	}
	return uniqs
}

func contentsBagCount(color string, tree map[string][]Contents) int {
	count := 1
	if children, ok := tree[color]; ok {
		for _, child := range children {
			count += child.count * contentsBagCount(child.color, tree)
		}
	}
	return count
}

func parse(lines []string) (map[string][]string, map[string][]Contents) {
	var noBagsRe = regexp.MustCompile(`^(.*) bags contain no other bags\.$`)
	var withBagsRe = regexp.MustCompile(`^(.*) bags contain (.+\.)$`)
	var contentRe = regexp.MustCompile(`(\d+) (.+?) bags?`)

	tree := map[string][]string{}
	tree2 := map[string][]Contents{}

	for _, rule := range lines {
		if m := noBagsRe.FindStringSubmatch(rule); m != nil {
			// do nothing
		} else {
			parts := withBagsRe.FindStringSubmatch(rule)
			color := parts[1]

			contents := contentRe.FindAllStringSubmatch(parts[2], -1)

			for _, cont := range contents {
				tree[cont[2]] = append(tree[cont[2]], color)
				tree2[color] = append(tree2[color], Contents{color: cont[2], count: toI(cont[1])})
			}
		}
	}

	return tree, tree2
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

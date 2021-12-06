package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	line := getLine("input.txt")
	init := splitAndToI(line)

	total := Run(80, fish(init))

	fmt.Println("PART ONE -- total: ", total)

	total = Run(256, fish(init))

	fmt.Println("PART TWO -- total: ", total)
}

func Run(days int, fish [9]int) (total int) {
	for i := 0; i < days; i++ {
		zeroCount := fish[0]

		for i := 1; i < len(fish); i++ {
			fish[i-1] = fish[i]
		}
		fish[8] = zeroCount
		fish[6] += zeroCount
	}

	for _, x := range fish {
		total += x
	}
	return total
}

func fish(input []int) [9]int {
	var fishies [9]int
	for _, i := range input {
		fishies[i]++
	}
	return fishies
}

// helpers
func getLine(path string) string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	s.Scan()
	return s.Text()
}

func splitAndToI(line string) []int {
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

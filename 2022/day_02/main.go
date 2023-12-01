package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pick int

const (
	_ Pick = iota
	Rock
	Paper
	Scissors
)

var winMap = map[Pick]Pick{
	Rock:     Scissors,
	Paper:    Rock,
	Scissors: Paper,
}

var loseMap = map[Pick]Pick{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var inpMap = map[string]Pick{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

// TWO
// X == lose
// Y == draw
// Z == win

func main() {
	file := "input.txt"

	lines := getLines(file)

	fmt.Println("Part One: ", partOne(lines))

	fmt.Println("Part Two: ", partTwo(lines))
}

func partOne(lines []string) (totes int) {
	for _, line := range lines {
		totes += shoot(parseRound(line))
	}
	return totes
}

func partTwo(lines []string) (totes int) {
	for _, line := range lines {
		dir, p2 := parseRound(line)
		var p1 Pick
		if dir == Rock { // X -- lose {
			p1 = winMap[p2]
		} else if dir == Paper { // Y -- tie
			p1 = p2
		} else {
			p1 = loseMap[p2]
		}
		totes += shoot(p1, p2)
	}
	return totes
}

func shoot(p1 Pick, p2 Pick) (score int) {
	score = int(p1)
	if p1 == p2 {
		score += 3
	} else if winMap[p1] == p2 {
		score += 6
	}

	return score
}

func parseRound(line string) (Pick, Pick) {
	return inpMap[line[2:]], inpMap[line[0:1]]
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

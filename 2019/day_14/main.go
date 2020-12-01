package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Chemical struct {
	chem string
	amt  int
}

type Formula struct {
	chemical Chemical
	reqs     []Chemical
}

func main() {
	input := getLines("input.txt")

	reactions := makeReactions(input)
	excess := make(map[string]int)
	ore := produce("FUEL", 1, reactions, excess)

	fmt.Println("PART ONE:")
	fmt.Println(ore)

	fmt.Println("Part Two:")

	max := maxFuel(reactions, 1000000000000)
	fmt.Println(max)
}

func maxFuel(reactions map[string]Formula, ore int) int {
	min := 0
	max := ore
	fuel := 0
	c := 0

	for {
		c++
		excess := make(map[string]int)
		//binary search
		fuel = (max-min)/2 + min

		reqOre := produce("FUEL", fuel, reactions, excess)

		if reqOre == ore {
			//winner
			break
		}
		if reqOre > ore {
			max = fuel
		} else {
			min = fuel
		}

		// so this seems to loop infinitely
		if c > 10000 {
			break
		}
	}

	return fuel
}

func makeReactions(reactions []string) map[string]Formula {
	f := make(map[string]Formula)

	for _, reaction := range reactions {
		reactionss := strings.Split(reaction, " => ")
		lhs, rhs := reactionss[0], reactionss[1]

		rhss := strings.Split(rhs, " ")
		chem := Chemical{rhss[1], toI(rhss[0])}

		lhss := strings.Split(lhs, ", ")

		freqs := []Chemical{}

		for _, req := range lhss {
			reqs := strings.Split(req, " ")
			rchem := Chemical{reqs[1], toI(reqs[0])}

			freqs = append(freqs, rchem)
		}

		f[chem.chem] = Formula{chemical: chem, reqs: freqs}
	}

	return f
}

func produce(chem string, amt int, reactions map[string]Formula, excess map[string]int) int {
	// want ORE? cool we have a lot
	if chem == "ORE" {
		return amt
	}

	// have we made enoughy already?
	if excess[chem] >= amt {
		excess[chem] -= amt
		return 0
	}

	// have me made some of it?
	if excess[chem] > 0 {
		amt -= excess[chem]
		excess[chem] = 0
	}

	formula := reactions[chem]
	batchCount := int(math.Ceil(float64(amt) / float64(formula.chemical.amt)))

	// feed the machine
	ore := 0

	for _, req := range formula.reqs {
		ore += produce(req.chem, req.amt*batchCount, reactions, excess)
	}

	amtMade := batchCount * formula.chemical.amt
	excess[chem] += amtMade - amt

	return ore
}

// setup
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

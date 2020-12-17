package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Ticket struct {
	fields []int
	valid  []bool
}

func main() {
	file := "input.txt"
	lines := getLines(file)
	ranges, _, tickets := parse(lines)
	fmt.Println("part one: ", errorRate(tickets, ranges))

	ranges2, myTicket, tickets2 := parse(lines)
	v := valid(tickets2, ranges2)

	mapping := indexer(v, ranges2)

	fmt.Println(mapping)

	fmt.Println("Two", calcIt(myTicket, departureFieldIndexes(mapping)))
}

func errorRate(tickets []*Ticket, ranges map[string][][]int) int {
	//0, 0 0,1 1,0 1,0
	for _, ticket := range tickets {
		for z, val := range ticket.fields {
			for _, rr := range ranges {
				if (val >= rr[0][0] && val <= rr[0][1]) || (val >= rr[1][0] && val <= rr[1][1]) {
					ticket.fields[z] = 0
					break
				}
			}
		}
	}

	sum := 0
	for _, ticket := range tickets {
		for _, val := range ticket.fields {
			sum += val
		}
	}
	return sum
}

func calcIt(ticket, dep []int) int {
	x := 1

	for _, i := range dep {
		x *= ticket[i]
	}

	return x
}

func departureFieldIndexes(m map[string]int) []int {
	x := []int{}

	for f, i := range m {
		if strings.HasPrefix(f, "departure") {
			x = append(x, i)
		}
	}

	return x
}

func indexer(tickets []*Ticket, ranges map[string][][]int) map[string]int {
	fieldIdxPoss := make(map[string]map[int]bool)

	for field, rr := range ranges {
		fieldIdxPoss[field] = make(map[int]bool)

		for j := 0; j < len(ranges); j++ {
			for i := 0; i < len(tickets); i++ {
				val := tickets[i].fields[j]
				if (val >= rr[0][0] && val <= rr[0][1]) || (val >= rr[1][0] && val <= rr[1][1]) {
					fieldIdxPoss[field][j] = true
				} else {
					fieldIdxPoss[field][j] = false
					break
				}

			}
		}
	}
	mapFieldToIdx := make(map[string]int)

	// oof
	for {
		done := true
		idx := -1
		for f, possibles := range fieldIdxPoss {
			sz := 0
			for _, y := range possibles {
				if y {
					sz++
				}
			}
			if sz == 1 {
				// set it
				done = false
				for i, y := range possibles {
					if y {
						idx = i
					}
				}
				mapFieldToIdx[f] = idx
			}
		}

		if done {
			break
		}
		for _, possibles := range fieldIdxPoss {

			possibles[idx] = false
		}

	}

	return mapFieldToIdx
}

func valid(tickets []*Ticket, ranges map[string][][]int) []*Ticket {
	//0, 0 0,1 1,0 1,0
	for _, ticket := range tickets {

		for _, val := range ticket.fields {
			for _, rr := range ranges {
				if (val >= rr[0][0] && val <= rr[0][1]) || (val >= rr[1][0] && val <= rr[1][1]) {
					//valid
					ticket.valid = append(ticket.valid, true)
					break
				}
			}

		}

	}
	v := []*Ticket{}

	for _, t := range tickets {
		if len(t.fields) == len(t.valid) {
			v = append(v, t)
		}
	}
	return v
}

func parse(lines []string) (map[string][][]int, []int, []*Ticket) {
	ranges := make(map[string][][]int)
	myTicket := []int{}
	tickets := []*Ticket{}
	// departure location: 40-261 or 279-955
	matcher := regexp.MustCompile(`^(.+?): (\d+)-(\d+) or (\d+)-(\d+)$`)

	section := 1
	for _, line := range lines {
		if line == "" {
			section++
			continue
		}
		if line == "your ticket:" || line == "nearby tickets:" {
			continue
		}

		if section == 1 {
			matches := matcher.FindStringSubmatch(line)

			field := matches[1]
			r1 := []int{toI(matches[2]), toI(matches[3])}
			r2 := []int{toI(matches[4]), toI(matches[5])}
			ranges[field] = [][]int{r1, r2}
		} else if section == 2 {
			myTicket = splitAndToI(line)
		} else if section == 3 {
			t := &Ticket{
				fields: splitAndToI(line),
			}
			tickets = append(tickets, t)
		}
	}

	return ranges, myTicket, tickets
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

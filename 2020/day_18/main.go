package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var splitter = regexp.MustCompile(`([^\s])`)

func main() {
	lines := getLines("input.txt")
	fmt.Println("PART ONE: ", accumAll(lines))
	fmt.Println("PART ONE (refactored): ", accumAll2(lines))

	fmt.Println("PART ONE (refactored): ", accumAll2Part2(lines))
}

func accumAll2Part2(lines []string) int {
	x := 0

	for _, line := range lines {
		chars := splitter.FindAllString(line, -1)
		x += accum2Part2(chars)
	}
	return x
}

func accumAll2(lines []string) int {
	x := 0

	for _, line := range lines {
		chars := splitter.FindAllString(line, -1)
		x += accum2(chars)
	}
	return x
}

func accumAll(lines []string) int {
	x := 0

	for _, line := range lines {
		chars := splitter.FindAllString(line, -1)
		x += accum(0, chars, 0)
	}
	return x
}

func accum2(s []string) int {
	parenIdxs := getOuterParens(s)

	// no parens we can do maths
	if len(parenIdxs) == 0 {
		x := 0
		op := ""

		for _, ins := range s {

			d, err := toI(ins)
			if err != nil {
				if ins == "*" || ins == "+" {
					op = ins
				}
			} else {
				if x == 0 {
					x += d
				} else if op == "*" {
					x *= d
				} else if op == "+" {
					x += d
				}
			}
		}
		return x
	}

	inner := s[parenIdxs[0]+1 : parenIdxs[1]]

	// collapse it
	// left side up to the paren
	left := s[:parenIdxs[0]]

	//eavalute and append the block
	left = append(left, strconv.Itoa(accum2(inner)))

	//right side up to closing paren
	outer := s[parenIdxs[1]+1:]
	left = append(left, outer...)

	return accum2(left)
}

func accum2Part2(s []string) int {
	parenIdxs := getOuterParens(s)

	// no parens we can do maths
	if len(parenIdxs) == 0 {
		pidx := firstPlus(s)

		// no plusses
		if pidx == 0 {
			x := 1

			for _, ins := range s {
				d, err := toI(ins)
				if err == nil {
					x *= d
				}
			}
			return x
		}

		left := s[:pidx-1]
		inner := s[pidx-1 : pidx+2]
		outer := s[pidx+2:]

		// I could toI these and sum them -- or just use my original method from part one
		reduced := append(left, strconv.Itoa(accum2(inner)))
		reduced = append(reduced, outer...)

		return accum2Part2(reduced)
	}

	// collapse it
	// left side up to the paren
	left := s[:parenIdxs[0]]
	inner := s[parenIdxs[0]+1 : parenIdxs[1]]
	outer := s[parenIdxs[1]+1:]

	//eavalute and append the block
	reduced := append(left, strconv.Itoa(accum2Part2(inner)))

	//right side up to closing paren
	reduced = append(reduced, outer...)

	return accum2Part2(reduced)
}

func getOuterParens(s []string) []int {
	depth := 0
	out := []int{}
	for i, c := range s {
		if c == "(" {
			if depth == 0 {
				out = append(out, i)
			}
			depth++
		}

		if c == ")" {
			depth--
			if depth == 0 {
				out = append(out, i)
				// only care about the first outer pair
				break
			}
		}
	}
	return out
}

func firstPlus(s []string) int {
	i := 0

	for i, c := range s {
		if c == "+" {
			return i
		}
	}

	return i
}

// in hidsight this would have probaly been easier by grouping the parentheses instead of tracking depth
// bad bad bad
// refactored into accum2
func accum(x int, s []string, d int) int {
	op := "+"
	depth := 0

	for i := 0; i < len(s); i++ {
		ins, s := s[i], s[i+1:]

		if depth > 0 {
			if ins == "(" {
				depth++
			}

			if ins != ")" {
				continue
			}
		}

		d, err := toI(ins)
		if err != nil {
			if ins == "*" || ins == "+" {
				op = ins
			} else if ins == "(" {
				depth++

				if op == "*" {
					x *= accum(0, s, depth)
				} else if op == "+" {
					x += accum(0, s, depth)
				}
			} else if ins == ")" {
				depth--
				if depth < 0 {

					return x
				}
			}
		} else {

			if x == 0 {
				x += d
			} else if op == "*" {
				x *= d
			} else if op == "+" {
				x += d
			}
		}

	}
	return x
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

func toI(s string) (int, error) {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	return i, err
}

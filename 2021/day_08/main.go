package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const pipe = " | "

func main() {
	lines := getLines("input.txt")
	uniqs, notes := getNotes(lines)
	count := countUniqs(notes)

	fmt.Println("PART ONE: ", count)

	accum := 0
	for i, uniq := range uniqs {
		mapping := decoder(uniq)

		n := buildNumber(notes[i], mapping)

		fmt.Println(notes[i], n)
		accum += n
	}

	fmt.Println("Part TWO :", accum)
}

func countUniqs(input [][]string) (count int) {
	for _, note := range input {
		for _, digit := range note {
			if len(digit) != 5 && len(digit) != 6 {
				count++
			}
		}
	}

	return count
}

func getNotes(lines []string) ([][]string, [][]string) {
	var notes [][]string
	var uniqs [][]string

	for _, line := range lines {
		n := strings.Split(line, pipe)
		u := strings.Fields(n[0])
		nn := strings.Fields(n[1])

		uniqs = append(uniqs, u)
		notes = append(notes, nn)

	}
	return uniqs, notes
}

func buildNumber(in []string, mapping map[string]string) int {
	var n []string
	for _, code := range in {
		d := mapping[lazySort(code)]
		n = append(n, d)
	}
	return toI(strings.Join(n, ""))
}

type code struct {
	len   int
	s     string
	val   int
	runes []rune
	digit string
}

func decoder(uniq []string) map[string]string {
	var codes []*code
	mapping := make(map[string]string)

	known := make(map[int]*code)
	for _, w := range uniq {
		val := wordVal(w)
		r := []rune(w)
		var d string
		switch len(w) {
		case 2:
			d = "1"
		case 3:
			d = "7"
		case 4:
			d = "4"
		case 7:
			d = "8"
		}
		c := &code{
			len:   len(w),
			s:     w,
			val:   val,
			digit: d,
			runes: r,
		}
		codes = append(codes, c)
		if d != "" {
			known[toI(d)] = c
		}
	}

	//   0
	// 1   2
	//   3
	// 4   5
	//   6
	segs := make(map[int]rune)
	zegz := make(map[rune]int)

	// step get segment 0
	z := known[7].val - known[1].val
	segs[0] = rune(z)
	zegz[rune(z)] = 0

	// get the len 6 codes and check for the sgements that are part of number 1
	for _, c := range codes {
		if c.len != 6 {
			continue
		}
		if !contains(c.runes, known[1].runes[0]) || !contains(c.runes, known[1].runes[1]) {

			known[6] = c
			c.digit = "6"
			if contains(c.runes, known[1].runes[0]) {
				segs[5] = known[1].runes[0]
				segs[2] = known[1].runes[1]
				zegz[known[1].runes[0]] = 5
				zegz[known[1].runes[1]] = 2
			} else {
				segs[2] = known[1].runes[0]
				segs[5] = known[1].runes[1]
				zegz[known[1].runes[0]] = 2
				zegz[known[1].runes[1]] = 5
			}
		}
	}

	//know we know 3 segments (0,5,2) and 5 number (1,4,6,7,8)
	four := known[4]
	var nz []*code
	for _, cc := range codes {
		if cc.len == 6 && cc.digit == "" {
			nz = append(nz, cc)
		}
	}
	for _, r := range four.runes {
		if _, ok := zegz[r]; !ok {
			// there are two segments in the 4 we don't know about. the segment that is only in one of 9 or 0 is segment 3
			if contains(nz[0].runes, r) && contains(nz[1].runes, r) {
				segs[1] = r
				zegz[r] = 1
			} else {
				segs[3] = r
				zegz[r] = 3
			}
		}
	}

	if contains(nz[0].runes, segs[3]) && contains(nz[0].runes, segs[1]) {
		known[9] = nz[0]
		nz[0].digit = "9"
		known[0] = nz[1]
		nz[1].digit = "0"
	} else {
		known[0] = nz[1]
		nz[1].digit = "9"
		known[9] = nz[0]
		nz[0].digit = "0"
	}
	// fmt.Println("-----------")
	// for _, c := range codes {
	// 	fmt.Println(c.digit, c.s)
	// }
	// fmt.Println("-----------")
	// known 0,1,4,6,7,8,9
	var l5 []*code
	var five *code
	for _, cc := range codes {
		if cc.len == 5 && cc.digit == "" {
			l5 = append(l5, cc)
			for _, r := range cc.runes {
				if r == segs[1] {
					known[5] = cc
					cc.digit = "5"
					five = cc
				}
			}
		}
	}
	for _, r := range five.runes {
		if _, ok := zegz[r]; !ok {
			segs[6] = r
			zegz[r] = 6
		}
	}

	for _, c := range l5 {
		if c.digit == "5" {
			continue
		} else {
			// for 3 we have them all mapped so only the two has a missing segment
			allokay := true
			for _, r := range c.runes {
				if _, ok := zegz[r]; !ok {
					segs[4] = r
					zegz[r] = 4
					c.digit = "2"
					known[2] = c
					allokay = false
				}
			}
			if allokay {
				c.digit = "3"
				known[3] = c
			}
		}
	}

	for _, c := range codes {
		// if c.digit == "" {
		// 	c.digit = "0"
		// 	known[0] = c
		// }
		// fmt.Println(c.digit, c.s)
		mapping[lazySort(c.s)] = c.digit
	}
	// for i := 0; i < 10; i++ {
	// 	c := known[i]
	// 	fmt.Println(c.s, c.digit)
	// }

	return mapping
}

func contains(s []rune, r rune) bool {
	for _, a := range s {
		if a == r {
			return true
		}
	}
	return false
}

// oof. still used to get sgement 0 but this broke badly for building the mapping
func wordVal(s string) (v int) {
	for _, r := range s {
		v += int(r)
	}
	return v
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

func lazySort(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

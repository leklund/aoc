package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	file := "input.txt"
	lines := getLines(file)

	mem := decodeV1(lines)

	tot := 0
	for _, x := range mem {
		tot += x
	}
	fmt.Println("Part 1 total:", tot)

	mem2 := decodeV2(lines)
	tot = 0
	for _, x := range mem2 {
		tot += x
	}
	fmt.Println("Part 2 total:", tot)
}

var matcher = regexp.MustCompile(`^(\w+)(?:\[(\d+)\])? = (.+)$`)

func decodeV1(lines []string) map[int]int {
	mem := make(map[int]int)
	mask := ""
	for _, line := range lines {
		match := matcher.FindStringSubmatch(line)
		if match[1] == "mask" {
			maskString := match[3]
			mask = ""
			for _, ch := range maskString {
				mask = string(ch) + mask
			}
			// fmt.Println("MASK:", mask)
		} else if match[1] == "mem" {
			loc := match[2]
			val := match[3]
			bval := strconv.FormatInt(int64(toI(val)), 2)
			//reverse it
			bs := ""
			for _, ch := range bval {
				bs = string(ch) + bs
			}
			// and split it
			bvals := strings.Split(bs, "")
			bvals = append(bvals, make([]string, 36-len(bvals))...)

			mv := 0
			for i, x := range mask {
				if x == '1' || (x != '0' && bvals[i] == "1") {
					mv += (1 << i)
				}
				mem[toI(loc)] = mv
			}
		}

	}
	return mem
}

func decodeV2(lines []string) map[int]int {
	mem := make(map[int]int)
	mask := ""
	for _, line := range lines {
		match := matcher.FindStringSubmatch(line)
		if match[1] == "mask" {
			maskString := match[3]
			mask = ""
			for _, ch := range maskString {
				mask = string(ch) + mask
			}
			// fmt.Println("MASK:", mask)
		} else if match[1] == "mem" {
			loc := toI(match[2])
			val := toI(match[3])

			bval := strconv.FormatInt(int64(loc), 2)
			//reverse it
			bs := ""
			for _, ch := range bval {
				bs = string(ch) + bs
			}
			// and split it
			bloc := strings.Split(bs, "")
			bloc = append(bloc, make([]string, 36-len(bloc))...)

			locs := []int{0}

			for i, x := range mask {
				if x == '0' {
					for j := range locs {
						if bloc[i] == "1" {
							locs[j] += 1 << i
						}
					}
				} else if x == '1' {
					for j := range locs {
						locs[j] += 1 << i
					}
				} else if x == 'X' {
					for j, a := range locs {
						locs = append(locs, a)
						locs[j] += 1 << i
					}
				}
			}
			for _, a := range locs {
				mem[a] = val
			}
		}
	}
	return mem

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

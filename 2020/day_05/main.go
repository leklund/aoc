package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pass struct {
	row int
	col int
}

func main() {
	file := "input.txt"

	input := getLines(file)

	ids := []int{}
	zids := []int{}
	for _, pass := range input {
		bstring := bs(pass)

		p := makePass(pass)
		zid := (p.row * 8) + p.col
		zids = append(zids, zid)
		id, _ := strconv.ParseUint(bstring, 2, 32)

		ids = append(ids, int(id))
	}

	sort.Ints(ids)
	sort.Ints(zids)
	fmt.Println(ids[len(ids)-1])
	fmt.Println(ids[len(zids)-1])
	for _, id := range zids {

		fmt.Println(id)
		// if ids[i+1] != ids[i]+1 {
		// 	fmt.Println(ids[i] + 1)
		// 	break
		// }
	}

}

func makePass(p string) Pass {
	fmt.Println(p)
	f := 0
	b := 127
	l := 0
	r := 7

	for _, char := range p {
		if char == 'F' {
			b = ((b - f) / 2) + f
		} else if char == 'B' {
			f = ((b - f) / 2) + f + 1
		} else if char == 'L' {
			r = ((r - l) / 2) + l
		} else if char == 'R' {
			l = ((r - l) / 2) + l + 1
		}
		fmt.Println(b, f, r, l)
	}
	return Pass{b, l}
}

func bs(line string) string {
	bstring := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1").Replace(line)
	return bstring
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

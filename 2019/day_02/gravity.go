package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// HALT is opcode 99
const HALT = 99

// ADD opcode 1
const ADD = 1

// MUL opcode is 2
const MUL = 2

func main() {
	line := getLine("input.txt")

	prog := splitAndToI(line)

	repair(prog)

	run(prog)

	fmt.Println("___PART ONE___")
	fmt.Println("index 0: ", prog[0])
}

func run(program []int) {
	for idx, opcode := range program {
		if idx%4 != 0 {
			continue
		}

		if opcode == HALT {
			return
		}

		p1, p2, op := program[idx+1], program[idx+2], program[idx+3]
		v1, v2 := program[p1], program[p2]

		if opcode == ADD {
			program[op] = v1 + v2
		} else if opcode == MUL {
			program[op] = v1 * v2
		}
	}
}

func repair(program []int) {
	// To do this, before running the program, replace position 1 with the value 12 and replace position 2 with the value 2.
	program[1] = 12
	program[2] = 2
}

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

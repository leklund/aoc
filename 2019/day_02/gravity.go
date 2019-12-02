package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const HLT = 99
const ADD = 1
const MUL = 2

func main() {
	line := getLine("input.txt")

	prog := splitAndToI(line)

	// part one
	testprog := setup(prog, 13, 2)
	run(testprog)

	fmt.Println("___PART ONE___")
	fmt.Println("index 0: ", testprog[0])

	// part two
	// brute force
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			testprog := setup(prog, noun, verb)

			run(testprog)

			if testprog[0] == 19690720 {
				fmt.Println("___ PART TWO ___")
				fmt.Print(100*noun + verb)
			}
		}
	}

}

func run(program []int) {
	for idx, opcode := range program {
		if idx%4 != 0 {
			continue
		}

		p1, p2, op := program[idx+1], program[idx+2], program[idx+3]
		v1, v2 := program[p1], program[p2]

		switch opcode {
		case HLT:
			return
		case ADD:
			program[op] = v1 + v2
		case MUL:
			program[op] = v1 * v2
		}

	}
}

func setup(program []int, noun int, verb int) []int {
	// To do this, before running the program, replace position 1 with the value 12 and replace position 2 with the value 2.
	newslice := make([]int, len(program))
	copy(newslice, program)
	newslice[1] = noun
	newslice[2] = verb
	return newslice
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

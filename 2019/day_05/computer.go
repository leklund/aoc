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
const INP = 3
const OUT = 4
const JIT = 5
const JIF = 6
const LT = 7
const EQL = 8

var pointerIncr = map[int]int{
	ADD: 4,
	MUL: 4,
	INP: 2,
	OUT: 2,
	HLT: 0,
	JIT: 3,
	JIF: 3,
	LT:  4,
	EQL: 4,
}

func main() {
	line := getLine("input.txt")

	prog := splitAndToI(line)

	prog2 := make([]int, len(prog))
	copy(prog2, prog)

	input := 1
	diagnostic := run(prog, input)

	fmt.Println("--- Part one ---")
	fmt.Println(diagnostic)

	input2 := 5
	diagnostic2 := run(prog2, input2)

	fmt.Println("--- Part Two ---")
	fmt.Println(diagnostic2)
}

func run(program []int, input int) []int {
	var outputCodes []int
	var pointer int

	for {
		var params [3]int
		var valuePointer int

		opcode, modes := parseOpcode(program[pointer])

		incr := pointerIncr[opcode]

		for i := 1; i < incr; i++ {
			if modes[i-1] == 0 {
				valuePointer = program[pointer+i]
			} else {
				valuePointer = pointer + i
			}

			params[i-1] = program[valuePointer]

		}
		// DEBUG
		// fmt.Println(pointer, opcode, modes, params)

		switch opcode {
		case HLT:
			return outputCodes
		case ADD:
			program[valuePointer] = params[0] + params[1]
		case MUL:
			program[valuePointer] = params[0] * params[1]
		case INP:
			program[valuePointer] = input
		case OUT:
			outputCodes = append(outputCodes, program[valuePointer])
		case JIT:
			if params[0] > 0 {
				pointer = params[1]
				continue
			}
		case JIF:
			if params[0] == 0 {
				pointer = params[1]
				continue
			}
		case LT:
			v := 0
			if params[0] < params[1] {
				v = 1
			}
			program[valuePointer] = v
		case EQL:
			v := 0
			if params[0] == params[1] {
				v = 1
			}
			program[valuePointer] = v
		}

		pointer += incr
	}
}

func parseOpcode(opcode int) (int, [3]int) {
	var modes [3]int
	op := opcode % 100
	opcode /= 100

	for i := 0; i < 3; i++ {
		modes[i] = opcode % 10
		opcode /= 10
	}
	return op, modes
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

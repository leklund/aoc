package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// HLT Halt opcode
const HLT = 99

// ADD - addition opcode
const ADD = 1

// MUL - multiplication opcode
const MUL = 2

// INP - input opcode
const INP = 3

// OUT - output opcode
const OUT = 4

// JIT - Jump if true opcode
const JIT = 5

// JIF - Jump if false opcode
const JIF = 6

// LT - less than opcode
const LT = 7

// EQL 0 equality opcode
const EQL = 8

// ARB Adjust Relative Base: changes the relative base for the relative mode
const ARB = 9

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
	ARB: 2,
}

// Program type so that I don't have to keep typing int64 over and over
type Program map[int64]int64

func makeProg(path string) Program {
	line := getLine("input.txt")
	return progFromString(line)
}

func progFromString(line string) Program {
	slicey := splitAndToI(line)

	return progFromSlice(slicey)
}

func progFromSlice(s []int64) Program {
	prog := make(map[int64]int64)
	for i, x := range s {
		prog[int64(i)] = x
	}
	return prog
}

func (program Program) dup() Program {
	duplicate := make(Program)

	for idx, op := range program {
		duplicate[idx] = op
	}
	return duplicate
}

func (program Program) run(in, out chan int64) {
	defer close(out)
	var pointer, relativeBase int64

	for program[pointer] != HLT {
		var params [3]int64
		var valuePointer int64

		opcode, modes := parseOpcode(program[pointer])

		incr := pointerIncr[opcode]
		for i := int64(1); i < int64(incr); i++ {
			switch modes[i-1] {
			case 0: // position mode
				valuePointer = program[pointer+i]
			case 1: // immediate mode
				valuePointer = pointer + i
			case 2: // relative mode
				valuePointer = program[pointer+i] + relativeBase
			}

			params[i-1] = program[valuePointer]
		}

		switch opcode {
		case HLT:
			break
		case ADD:
			program[valuePointer] = params[0] + params[1]
		case MUL:
			program[valuePointer] = params[0] * params[1]
		case INP:
			program[valuePointer] = <-in
		case OUT:
			out <- params[0]
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
			program[valuePointer] = int64(v)
		case EQL:
			v := 0
			if params[0] == params[1] {
				v = 1
			}
			program[valuePointer] = int64(v)
		case ARB:
			relativeBase += params[0]
		}

		pointer += int64(incr)
	}
}

func parseOpcode(opcode int64) (int, [3]int) {
	var modes [3]int
	op := opcode % 100
	opcode /= 100

	for i := 0; i < 3; i++ {
		modes[i] = int(opcode % 10)
		opcode /= 10
	}
	return int(op), modes
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

func splitAndToI(line string) []int64 {
	s := strings.Split(line, ",")
	var p []int64
	for _, inst := range s {
		p = append(p, toI(inst))
	}

	return p
}

func toI(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)

	if err != nil {
		panic(err)
	}

	return i
}

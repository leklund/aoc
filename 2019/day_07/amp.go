package amp

import (
	"bufio"
	"fmt"
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

	signal := amp(prog)

	fmt.Println(signal)
}

func amp(program []int) int {
	p0 := []int{0, 1, 2, 3, 4}
	phases := [][]int{}
	generatePermutations(len(p0), p0, &phases)

	signal := 0

	for _, permutation := range phases {
		i2 := 0
		for _, phase := range permutation {
			i2 = run(program, phase, i2)[0]
		}
		if i2 > signal {
			signal = i2
		}
	}
	return signal
}

func run(p []int, inputs ...int) []int {
	program := make([]int, len(p))
	copy(program, p)

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
			input := inputs[0]
			inputs = inputs[1:]
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

// Heap's algorithim (https://en.wikipedia.org/wiki/Heap%27s_algorithm)
func generatePermutations(k int, a []int, res *[][]int) {
	if k == 1 {
		aa := make([]int, len(a))
		copy(aa, a)
		*res = append(*res, aa)
	}

	for i := 0; i < k; i++ {
		generatePermutations(k-1, a, res)
		if k%2 == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}

	}
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

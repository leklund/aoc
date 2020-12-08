package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Program struct {
	code     []*Instruction
	acc      int
	curr     *Instruction
	pointer  int
	repaired bool
}

type Instruction struct {
	op  string
	arg int
	ex  bool
}

var swap = map[string]string{
	"jmp": "nop",
	"nop": "jmp",
}

func main() {
	file := "input.txt"
	program := makeProg(getLines(file))

	program.runUntilHalt()
	fmt.Println("PART ONE: ", program.acc)

	program.runUntilRepaired()
	fmt.Println("PART TWO:", program.acc)
}

func makeProg(lines []string) *Program {
	p := &Program{
		code: []*Instruction{},
	}
	for _, line := range lines {
		strs := strings.Split(line, " ")

		p.code = append(p.code, &Instruction{op: strs[0], arg: toI(strs[1])})
	}
	return p
}

func (program *Program) reset() {
	program.acc = 0
	program.curr = &Instruction{}
	program.pointer = 0
	program.repaired = false

	// I used a map to track instruction exceution but this extra iteration over the code doesn't make it any slower for solving part 2.
	for _, inst := range program.code {
		inst.ex = false
	}
}

func (program *Program) runUntilHalt() {
	for {
		program.curr = program.code[program.pointer]
		if program.curr.ex {
			break
		}
		program.execute(program.curr.op)
	}
}

func (program *Program) runUntilRepaired() {
	tested := make(map[*Instruction]bool)
	for {
		program.reset()

		for {
			if program.pointer >= len(program.code) {
				// fmt.Printf("returning at pointer %d out of range\n", program.pointer)
				return
			}

			program.curr = program.code[program.pointer]
			if program.curr.ex {
				break
			}

			op := program.curr.op
			swop, ok := swap[op]
			if !program.repaired && ok {
				if _, ok := tested[program.curr]; !ok {
					tested[program.curr] = true
					program.repaired = true

					// fmt.Println("attempting repair on ", program.curr)
					op = swop
				}
			}
			program.execute(op)
		}
	}
}

func (program *Program) execute(op string) {
	program.curr.ex = true
	switch op {
	case "nop":
		program.pointer++
	case "acc":
		program.acc += program.curr.arg
		program.pointer++
	case "jmp":
		program.pointer += program.curr.arg
	}
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

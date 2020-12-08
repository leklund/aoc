package main

import (
	"strings"
	"testing"
)

var sample = `nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6`

func TestMakeProg(t *testing.T) {
	program := makeProg(strings.Split(sample, "\n"))

	if program.code[0].op != "nop" {
		t.Errorf("expected nop got %v", program.code[0].op)
	}
}

func TestRunUntilHalt(t *testing.T) {
	program := makeProg(strings.Split(sample, "\n"))
	program.runUntilHalt()

	if program.acc != 5 {
		t.Errorf("expected acc to be 5, got %d", program.acc)
	}
}

func TestRunUntilRepaired(t *testing.T) {
	program := makeProg(strings.Split(sample, "\n"))
	program.runUntilRepaired()

	if program.acc != 8 {
		t.Errorf("expected acc to be 8, got %d", program.acc)
	}
}

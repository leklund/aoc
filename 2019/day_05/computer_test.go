package main

import "testing"

type Opcode struct {
	code  int
	modes [3]int
}

var opTest = []struct {
	opcode   int
	expected Opcode
}{
	{1002, Opcode{2, [3]int{0, 1, 0}}},
	{104, Opcode{4, [3]int{1, 0, 0}}},
	{10001, Opcode{1, [3]int{0, 0, 1}}},
}

func TestParseOpCode(t *testing.T) {
	for _, testCase := range opTest {
		code, modes := parseOpcode(testCase.opcode)

		if code != testCase.expected.code {
			t.Errorf("Error, expected %d, got %d", testCase.expected.code, code)
		}

		if modes != testCase.expected.modes {
			t.Errorf("Error, expected %v, got %v", testCase.expected.modes, modes)
		}
	}

}

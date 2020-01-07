package main

import (
	"testing"
)

var testCases = []struct {
	in     string
	inputs []int64
	out    []int64
}{
	{"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", []int64{}, []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}},
	{"1102,34915192,34915192,7,4,7,99,0", []int64{}, []int64{1219070632396864}},
	{"104,1125899906842624,99", []int64{}, []int64{1125899906842624}},
	{
		"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
		[]int64{1},
		[]int64{999},
	},
	{
		"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
		[]int64{8},
		[]int64{1000},
	},
	{
		"3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99",
		[]int64{10},
		[]int64{1001},
	},
}

func TestRun(t *testing.T) {
	for _, testCase := range testCases {
		in, out := make(chan int64), make(chan int64)
		program := progFromString(testCase.in)

		program.run(in, out)

		for _, z := range testCase.inputs {
			in <- z
		}

		res := []int64{}
		for x := range out {
			res = append(res, x)
		}

		if len(res) != len(testCase.out) {
			t.Errorf("Error, output len mismatch: expected %d, got %d", len(testCase.out), len(res))
		}
		for i, y := range res {
			if y != testCase.out[i] {
				t.Errorf("Error, expected slice index %d to match. Expected %d, got %d", i, testCase.out[i], y)
			}
		}
	}
}

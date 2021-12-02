package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file := "input.txt"
	ints := readInts(file)

	inc := CountIncreases(ints)
	fmt.Printf("Part One: %d increases \n", inc)

	inc = CountIncreaseSums(ints)
	fmt.Printf("Part Two: %d increases \n", inc)
}

func CountIncreases(ints []int) int {
	var last, inc int

	for i, x := range ints {
		if i == 0 {
			last = x
			continue
		}

		if x > last {
			inc++
		}
		last = x

	}
	return inc
}

func CountIncreaseSums(ints []int) int {
	var previousWindow, currentWindow, inc int

	windows := slidingWindow(3, ints)

	for i, w := range windows {
		if i == 0 {
			previousWindow = sumSlice(w)
			continue
		}
		currentWindow = sumSlice(w)

		if currentWindow > previousWindow {
			inc++
		}
		previousWindow = currentWindow
	}

	return inc
}

// helpers

func sumSlice(ints []int) int {
	var x int

	for _, y := range ints {
		x += y
	}
	return x
}

// thanks slice tricks! :)
func slidingWindow(size int, input []int) [][]int {
	// returns the input slice as the first element
	if len(input) <= size {
		return [][]int{input}
	}

	// allocate slice at the precise size we need
	r := make([][]int, 0, len(input)-size+1)

	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}

	return r
}

func readInts(path string) []int {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanWords)

	var nums []int
	for s.Scan() {
		nums = append(nums, toI(s.Text()))
	}

	return nums
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

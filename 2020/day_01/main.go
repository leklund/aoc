package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	file := "input.txt"
	ints := readInts(file)

	x, y := Find2020(ints)

	fmt.Printf("Part One\nx: %v, y: %v, product: %v \n", x, y, x*y)

	x, y, z := FindThree2020(ints)

	fmt.Printf("Part Two\nx: %v, y: %v, z: %v, product: %v \n", x, y, z, x*y*z)
}

func Find2020(ints []int) (int, int) {
	sort.Ints(ints)

	for _, x := range ints {
		for i := len(ints) - 1; i >= 0; i-- {
			y := ints[i]
			if x+y == 2020 {
				return x, y
			}
		}
	}
	return 0, 0
}

func FindThree2020(ints []int) (int, int, int) {
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	for _, x := range ints {
		for i := len(ints) - 1; i >= 0; i-- {
			y := ints[i]
			if x+y+ints[i-1] > 2020 {
				break
			}
			for j := i - 1; j >= 0; j-- {
				z := ints[j]

				if x+y+z == 2020 {
					return x, y, z
				}
			}

		}
	}
	return 1, 1, 1
}

// helpers
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

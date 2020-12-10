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
	xmas := readInts(file)

	err := findErr(xmas, 25)

	fmt.Println("PART ONE", err)

	cont := findCont(xmas, err)

	sort.Ints(cont)

	fmt.Println("PART TWO: ", cont[0]+cont[len(cont)-1])

}

func findErr(xmas []int, pre int) int {
	idx := pre

	for idx < len(xmas)-1 {
		i := idx - pre
		scan := xmas[i:idx]

		if !checkSum(scan, xmas[idx]) {
			return xmas[idx]
		}
		idx++
	}

	return -1
}

func findCont(xmas []int, err int) []int {
	for i := 0; i < len(xmas)-1; i++ {
		j := i

		for sum := xmas[j]; sum < err; j++ {
			sum += xmas[j+1]
			if sum == err {
				return xmas[i : j+2]
			}
		}
	}
	return []int{}
}

func checkSum(s []int, n int) bool {
	// copy and sort the slice
	scan := make([]int, len(s))
	copy(scan, s)
	sort.Ints(scan)

	i := 0
	j := len(scan) - 1
	for i < j {
		sum := scan[i] + scan[j]

		if sum == n {
			return true
		} else if sum < n {
			i++
		} else if sum > n {
			j--
		}
	}
	return false
}

// boiler plate
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

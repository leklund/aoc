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
	adapters := readInts(file)

	gaps := findGaps(adapters)

	fmt.Println("Part One: ", gaps[1]*gaps[3])

	perms := permute(adapters)

	fmt.Println("PART TWO", perms)

}

func findGaps(adapters []int) map[int]int {
	gaps := make(map[int]int)

	gaps[adapters[0]]++
	for i, a := range adapters {
		if i == 0 {
			continue
		}

		gap := a - adapters[i-1]
		gaps[gap]++
	}

	gaps[3]++
	return gaps
}

func permute(adapters []int) int {
	max := adapters[len(adapters)-1] + 3
	adapters = append(adapters, max)
	adapters = append([]int{0}, adapters...)

	cache := make(map[int]int)

	var candidates []int
	for i, x := range adapters {
		tracker := 0
		if i <= 1 {
			cache[x] = 1
			continue
		} else if i == 2 {
			candidates = adapters[i-2 : 2]
		} else {
			candidates = adapters[i-3 : i]
		}

		for _, c := range candidates {
			if c+3 >= x {
				tracker++
			}
		}

		l := len(candidates)
		if tracker == 3 {
			cache[x] = cache[candidates[0]] + cache[candidates[1]] + cache[candidates[2]]
		} else if tracker == 2 {
			cache[x] = cache[candidates[l-2]] + cache[candidates[l-1]]
		} else {
			cache[x] = cache[candidates[l-1]]
		}
	}
	return cache[max]
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
	sort.Ints(nums)
	return nums
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

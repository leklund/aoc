package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	modules := readInts("input.txt")

	total, totalWithFuel := 0, 0

	for _, mass := range modules {
		total += calcFuel(mass)
		totalWithFuel += calcFuelInclusive(mass)
	}

	fmt.Println("___ Part 1 ___")
	fmt.Println(total)
	fmt.Println("___ Part 2 ___")
	fmt.Println(totalWithFuel)
}

func calcFuel(mass int) int {
	return mass/3 - 2
}

func calcFuelInclusive(mass int) int {
	fuel := calcFuel(mass)

	if fuel <= 0 {
		return 0
	}

	return fuel + calcFuelInclusive(fuel)
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

package main

import "fmt"

func slicer(x int) [6]int {
	var res [6]int
	for i := 5; i >= 0; i-- {
		d := x % 10
		x /= 10
		res[i] = d
	}
	return res
}

func validPass(pass [6]int) bool {
	// always increasing
	paired := false
	for i := 0; i < 5; i++ {
		if pass[i] > pass[i+1] {
			return false
		}
		if pass[i] == pass[i+1] {
			paired = true
		}
	}
	return paired
}

func validPassTwo(pass [6]int) bool {
	// always increasing
	paired := make(map[int]int)
	for i := 0; i < 5; i++ {
		if pass[i] > pass[i+1] {
			return false
		}
		if pass[i] == pass[i+1] {
			paired[pass[i]]++
		}
	}
	double := false
	for _, c := range paired {
		if c == 1 {
			double = true
		}
	}
	return double
}

func main() {
	valid := 0
	for i := 136818; i <= 685979; i++ {
		if validPass(slicer(i)) {
			valid++
		}

	}

	fmt.Println("___ Part One ___")
	fmt.Println(valid)

	validTwo := 0
	for i := 136818; i <= 685979; i++ {
		if validPassTwo(slicer(i)) {
			validTwo++
		}

	}

	fmt.Println("___ Part Two ___")
	fmt.Println(validTwo)
}

package main

import (
	"fmt"
)

func main() {
	prog := makeProg("input.txt")

	signal := amp(prog)

	fmt.Println("PART ONE:")
	fmt.Println(signal)

	signal2 := ampFeedback(prog)
	fmt.Println("PART TWO:")
	fmt.Println(signal2)
}

func amp(program Program) int64 {
	p0 := []int{0, 1, 2, 3, 4}
	permutations := [][]int{}
	max := int64(0)
	in := make(chan int64)
	out := make(chan int64)

	generatePermutations(len(p0), p0, &permutations)
	fmt.Println("phases size:", len(permutations))

	for _, phases := range permutations {

		output := int64(0)
		for _, phase := range phases {
			p := program.dup()

			go p.run(in, out)
			in <- int64(phase)
			in <- output
			output = <-out
		}
		if output > max {
			max = output
		}

	}
	return max
}

func ampFeedback(program Program) int64 {
	p0 := []int{5, 6, 7, 8, 9}
	permutations := [][]int{}
	max := int64(0)

	// one input and output channel per amplifier
	// could likely be simplified to be one slice of chans
	in := make([]chan int64, 5)
	out := make([]chan int64, 5)

	generatePermutations(len(p0), p0, &permutations)

	for _, phases := range permutations {
		for i, phase := range phases {
			p := program.dup()

			in[i] = make(chan int64)
			out[i] = make(chan int64)

			go p.run(in[i], out[i])

			in[i] <- int64(phase)
		}

		output := int64(0)
		running := true

		for running {
			for i := range phases {
				select {
				case in[i] <- output:
				default:
					running = false
				}
				if !running {
					break
				}
				output = <-out[i]
			}
		}
		if output > max {
			max = output
		}
	}

	return max
}

// Heap's algorithim (https://en.wikipedia.org/wiki/Heap%27s_algorithm)
func generatePermutations(k int, a []int, res *[][]int) {
	if k == 1 {
		aa := make([]int, len(a))
		copy(aa, a)
		*res = append(*res, aa)
	}

	for i := 0; i < k; i++ {
		generatePermutations(k-1, a, res)
		if k%2 == 0 {
			a[i], a[k-1] = a[k-1], a[i]
		} else {
			a[0], a[k-1] = a[k-1], a[0]
		}

	}
}

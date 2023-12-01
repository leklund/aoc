package main

import "fmt"

type Point struct {
	x int
	y int
}

type Vector struct {
	x int
	y int
}
type Box struct {
	tl Point
	br Point
}

func main() {
	t := "target area: x=56..76, y=-162..-134"
	fmt.Println(t)
}

func findMaxY(target Box) Vector {

	return Vector{0, 0}
}

func nsum(n int) int {
	return (n * (n + 1) / 2)
}

package main

import "fmt"

func main() {
	one()
	two()
}

func one() {
	program := makeProg("input.txt")
	in, out := make(chan int64), make(chan int64)

	fmt.Println("RUN")
	go program.run(in, out)

	in <- int64(1)
	fmt.Println(<-out)
}

func two() {
	program := makeProg("input.txt")
	in, out := make(chan int64), make(chan int64)

	fmt.Println("RUN")
	go program.run(in, out)

	in <- int64(2)
	fmt.Println(<-out)
}

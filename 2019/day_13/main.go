package main

import "fmt"

type Point struct {
	x int64
	y int64
}

type Game map[Point]int64

func main() {
	game := one()

	counter := 0
	for _, id := range game {
		if id == 2 { // bolock tile
			counter++
		}
	}
	fmt.Println("PART ONE -- # of block tiles:", counter)
}

func one() Game {
	program := makeProg("input.txt")

	game := make(Game)
	in, out := make(chan int64, 1), make(chan int64, 2)

	go program.run(in, out)

	for x := range out {
		y := <-out
		id := <-out
		game[Point{x, y}] = id
	}
	return game
}

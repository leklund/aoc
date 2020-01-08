package main

import (
	"fmt"
	"strings"
)

type Point struct {
	x int64
	y int64
}

type Game map[Point]int64

var charMap = map[int64]string{
	0: " ", // empty
	1: "■", // wall
	2: "✩", // block
	3: "_", // paddle
	4: "◉", // ball
}

func main() {
	game := one()

	counter := 0
	for _, id := range game {
		if id == 2 { // bolock tile
			counter++
		}
	}
	fmt.Println("PART ONE -- # of block tiles:", counter)
	printGame(game, 0)

	two()
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

func two() {
	program := makeProg("input.txt")

	//insert quarter
	program[0] = 2

	game := make(Game)
	in, out := make(chan int64, 1), make(chan int64, 2)

	go program.run(in, out)

	var ball, paddle Point
	var score int64

	for x := range out {
		y := <-out
		id := <-out

		if x == -1 && y == 0 {
			score = id
			continue
		}
		pos := Point{x, y}
		game[pos] = id

		switch id {
		case 4: // ball
			ball = pos
			printGame(game, score)
			if ball.x > paddle.x {
				in <- 1
			} else if ball.x < paddle.x {
				in <- -1
			} else {
				in <- 0
			}
		case 3:
			paddle = pos
		}

		// add a tick to make the gameplay watchable.
		// time.Sleep(2500000)
	}
	printGame(game, score)
}

// i wish this looked better on iterm2
func printGame(game Game, score int64) {
	var ymax, xmax int64

	for tile := range game {
		xmax = max(xmax, tile.x)
		ymax = max(ymax, tile.y)
	}

	screen := &strings.Builder{}
	screen.WriteString("\033c")

	fmt.Fprintf(screen, "Score: %d\n", score)
	for y := int64(0); y <= ymax; y++ {
		for x := int64(0); x <= xmax; x++ {
			screen.WriteString(charMap[game[Point{x, y}]])
		}
		screen.WriteRune('\n')
	}

	fmt.Print(screen.String())
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}

	return b
}

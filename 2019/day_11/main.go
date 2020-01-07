package main

import (
	"container/ring"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"os"
	"sort"
)

type Point struct {
	x int
	y int
}

type Turtle struct {
	pos Point
	dir *ring.Ring
}

var colorMap = map[int64]color.RGBA{
	0: color.RGBA{0, 0, 0, 0xff},          // black
	1: color.RGBA{0xff, 0xff, 0xff, 0xff}, // white
}

func main() {
	painter(0)
	g := painter(1)

	makeImg(g)
}

func painter(initColor int64) map[Point]int64 {
	program := makeProg("input.txt")
	in, out := make(chan int64, 1), make(chan int64, 2)
	r := ring.New(4)
	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	turtle := Turtle{
		pos: Point{x: 0, y: 0},
		dir: r,
	}

	graph := make(map[Point]int64)
	graph[turtle.pos] = initColor

	done := program.run(in, out)

	//run until done
looper:
	for {
		select {
		case <-done:
			break looper

		default:
			//run the program
			//send the input
			in <- graph[turtle.pos]
			// get the outputs
			color, direction := <-out, <-out

			// PAINT
			graph[turtle.pos] = color
			// fmt.Println(graph)

			// TURN
			switch direction {
			case 0:
				turtle.dir = turtle.dir.Prev()
			case 1:
				turtle.dir = turtle.dir.Next()
			default:
				panic(fmt.Sprintf("got invalid direction: %d", direction))
			}

			// MOVE
			x, y := turtle.pos.x, turtle.pos.y
			switch turtle.dir.Value {
			case 0:
				y++
			case 1:
				x++
			case 2:
				y--
			case 3:
				x--
			default:
				panic("invalid direction")
			}
			turtle.pos = Point{x, y}

		}
	}

	fmt.Println("--ONE--")
	fmt.Println("painted a lot of panels:")
	fmt.Println(len(graph))

	return graph
}

func makeImg(graph map[Point]int64) {
	// find min/max x/y to normalize to only positive integers from 0,0
	var x, y []int

	for p := range graph {
		x = append(x, p.x)
		y = append(y, p.y)
	}
	sort.Ints(x)
	sort.Ints(y)

	xoffset := x[0]
	yoffset := y[0]

	xmax := x[len(x)-1] - xoffset
	ymax := y[len(y)-1] - yoffset

	origin := image.Point{0, 0}
	lowerRight := image.Point{xmax + 1, ymax + 1}

	fmt.Println("rect", origin, lowerRight)

	rgb := image.NewRGBA(image.Rectangle{origin, lowerRight})

	for p, c := range graph {
		p.x -= xoffset
		p.y -= yoffset

		// invert the y. painter started from bottom but img is from top down.
		rgb.Set(p.x, ymax-p.y, colorMap[c])
	}

	f, _ := os.Create("out.png")
	png.Encode(f, rgb)
}

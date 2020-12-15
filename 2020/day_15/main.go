package main

import "fmt"

type Num struct {
	count int
	first int
	last  int
}

type Game struct {
	input   []int
	lastMap map[int]*Num
	spoken  int
}

func main() {
	input := []int{7, 14, 0, 17, 11, 1, 2}

	game := &Game{
		input:   input,
		lastMap: make(map[int]*Num),
	}
	game2 := &Game{
		input:   input,
		lastMap: make(map[int]*Num),
	}

	fmt.Println(game.play(2020))
	fmt.Println(game2.play(30000000))
}

func (g *Game) play(count int) int {
	turn := 1

	for _, x := range g.input {
		g.spoken = x
		n := &Num{
			count: 1,
			first: turn,
			last:  turn,
		}
		g.lastMap[x] = n
		// fmt.Println(turn, g.spoken)
		turn++
	}

	for {
		if num, ok := g.lastMap[g.spoken]; !ok {
			g.lastMap[g.spoken] = &Num{
				count: 1,
				first: turn,
				last:  turn,
			}

			g.spoken = 0
		} else {
			if num.count <= 1 {
				g.sayIt(0, turn)
			} else {
				z := num.last - num.first
				num.first = num.last
				g.sayIt(z, turn)
			}
			num.count++
		}

		// fmt.Println(turn, g.spoken)
		if turn == count {
			return g.spoken
		}

		turn++
	}
}

func (g *Game) sayIt(n, t int) {
	g.spoken = n
	if num, ok := g.lastMap[n]; !ok {
		nn := &Num{
			count: 1,
			first: t,
			last:  t,
		}
		g.lastMap[n] = nn
	} else {
		num.last = t
		num.count++
	}
}

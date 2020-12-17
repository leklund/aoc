package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
	z int
	w int
}

type Cube struct {
	state map[Point]bool
	xmin  int
	xmax  int
	ymin  int
	ymax  int
	zmin  int
	zmax  int
	wmin  int
	wmax  int
}

func main() {
	lines := getLines("input.txt")

	cubes := initCube(lines)

	cubes.run(6)

	fmt.Println("Part Two: ", cubes.active())
}

func (cube *Cube) run(steps int) {
	for i := 0; i < steps; i++ {
		newState := make(map[Point]bool)

		for point, active := range cube.state {
			newState[point] = cube.stateChange(point, active)
		}

		cube.state = newState
		cube.padding()

	}

}

func (cube *Cube) stateChange(p Point, a bool) bool {
	activeNeighbors := 0

	for x := p.x - 1; x <= p.x+1; x++ {
		for y := p.y - 1; y <= p.y+1; y++ {
			for z := p.z - 1; z <= p.z+1; z++ {
				for w := p.w - 1; w <= p.w+1; w++ {
					n := Point{x, y, z, w}
					if n == p {
						continue
					}
					if cube.state[n] {
						activeNeighbors++
					}
				}
			}
		}

	}

	if a {
		if activeNeighbors == 2 || activeNeighbors == 3 {
			return true
		}
	} else {
		if activeNeighbors == 3 {
			return true
		}
	}
	return false
}

func (cube *Cube) active() int {
	a := 0
	for _, active := range cube.state {
		if active {
			a++
		}
	}
	return a
}

func initCube(lines []string) *Cube {
	state := make(map[Point]bool)
	cube := &Cube{state: state}
	z, w := 0, 0
	for y, line := range lines {
		if y > cube.ymax {
			cube.ymax = y
		}
		for x, ch := range line {
			if x > cube.xmax {
				cube.xmax = x
			}
			active := ch == '#'

			cube.state[Point{x, y, z, w}] = active
		}
	}

	cube.padding()

	return cube
}

func (cube *Cube) padding() {
	for p, a := range cube.state {
		// expand -x
		if p.x <= cube.xmin && a {
			cube.xmin = p.x - 1
			x := cube.xmin
			for y := cube.ymin; y <= cube.ymax; y++ {
				for z := cube.zmin; z <= cube.zmax; z++ {
					for w := cube.wmin; w <= cube.wmax; w++ {
						cube.state[Point{x, y, z, w}] = false
					}
				}
			}
		}

		// expand +x
		if p.x >= cube.xmax && a {
			cube.xmax = p.x + 1
			x := cube.xmax
			for y := cube.ymin; y <= cube.ymax; y++ {
				for z := cube.zmin; z <= cube.zmax; z++ {
					for w := cube.wmin; w <= cube.wmax; w++ {
						cube.state[Point{x, y, z, w}] = false
					}
				}
			}
		}

		// expand -y
		if p.y <= cube.ymin && a {
			cube.ymin = p.y - 1
			y := cube.ymin
			for x := cube.xmin; x <= cube.xmax; x++ {
				for z := cube.zmin; z <= cube.zmax; z++ {
					for w := cube.wmin; w <= cube.wmax; w++ {
						cube.state[Point{x, y, z, w}] = false
					}
				}
			}
		}

		// expand +y
		if p.y >= cube.ymax && a {
			cube.ymax = p.y + 1
			y := cube.ymax
			for x := cube.xmin; x <= cube.xmax; x++ {
				for z := cube.zmin; z <= cube.zmax; z++ {
					for w := cube.wmin; w <= cube.wmax; w++ {
						cube.state[Point{x, y, z, w}] = false
					}
				}
			}
		}

		// expand -z
		if p.z <= cube.zmin && a {
			cube.zmin = p.z - 1
			z := cube.zmin
			for y := cube.ymin; y <= cube.ymax; y++ {
				for x := cube.xmin; x <= cube.xmax; x++ {
					for w := cube.wmin; w <= cube.wmax; w++ {
						cube.state[Point{x, y, z, w}] = false
					}
				}
			}
		}

		// expand +z
		if p.z >= cube.zmax && a {
			cube.zmax = p.z + 1
			z := cube.zmax
			for y := cube.ymin; y <= cube.ymax; y++ {
				for x := cube.xmin; x <= cube.xmax; x++ {
					for w := cube.wmin; w <= cube.wmax; w++ {
						cube.state[Point{x, y, z, w}] = false
					}
				}
			}
		}

		// expand -w
		if p.w <= cube.wmin && a {
			cube.wmin = p.w - 1
			w := cube.wmin
			for y := cube.ymin; y <= cube.ymax; y++ {
				for x := cube.xmin; x <= cube.xmax; x++ {
					for z := cube.zmin; z <= cube.zmax; z++ {
						cube.state[Point{x, y, z, w}] = false
					}
				}
			}
		}

		// expand +w
		if p.w >= cube.wmax && a {
			cube.wmax = p.w + 1
			w := cube.wmax
			for y := cube.ymin; y <= cube.ymax; y++ {
				for x := cube.xmin; x <= cube.xmax; x++ {
					for z := cube.zmin; z <= cube.zmax; z++ {
						cube.state[Point{x, y, z, w}] = false
					}
				}
			}
		}
	}
}

// boiler plate
func getLines(path string) []string {
	file, err := os.Open(path)

	var lines []string

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

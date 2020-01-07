package main

import "fmt"

type Point struct {
	x int
	y int
	z int
}

type Moon struct {
	pos Point
	vel Point
}

func main() {
	moons := getInput()

	for i := 0; i < 1000; i++ {
		step(moons)
	}

	fmt.Println("Part One: energy = ", energy(moons))

	// part two
	m1 := getInput()
	m2 := getInput()
	fmt.Println("**TWO: ", two(m1, m2))
}

// thanks for the hints reddit.
func two(moons, initMoons []*Moon) int64 {
	//cycle each axis
	cycleCount := []int64{1, 1, 1}
	axisRepeated := make([]bool, 3)

	fin := axisRepeated[0] && axisRepeated[1] && axisRepeated[2]

	fmt.Println(moons[0], moons[1], moons[2], moons[3])

	for !fin {
		step(moons)

		for axis := 0; axis < 3; axis++ {
			if !axisRepeated[axis] {
				done := true

				for i, moon := range moons {
					if !moonRepeat(moon, initMoons[i], axis) {
						done = false
					}
				}

				if done {
					axisRepeated[axis] = true
				} else {
					cycleCount[axis]++
				}
			}
		}
		fin = axisRepeated[0] && axisRepeated[1] && axisRepeated[2]
	}

	fmt.Println(cycleCount)
	return lcm(cycleCount[0], cycleCount[1], cycleCount[2])
}

func moonRepeat(a *Moon, b *Moon, axis int) bool {
	switch axis {
	case 0:
		return a.pos.x == b.pos.x && a.vel.x == b.vel.x
	case 1:
		return a.pos.y == b.pos.y && a.vel.y == b.vel.y
	case 2:
		return a.pos.z == b.pos.z && a.vel.z == b.vel.z
	}
	return false
}

func step(moons []*Moon) {
	// gravity
	for i := 0; i < len(moons)-1; i++ {
		for j := i + 1; j < len(moons); j++ {
			a, b := moons[i], moons[j]

			// x
			if a.pos.x > b.pos.x {
				b.vel.x++
				a.vel.x--
			} else if a.pos.x != b.pos.x {
				a.vel.x++
				b.vel.x--
			}

			// x
			if a.pos.y > b.pos.y {
				b.vel.y++
				a.vel.y--
			} else if a.pos.y != b.pos.y {
				a.vel.y++
				b.vel.y--
			}

			// z
			if a.pos.z > b.pos.z {
				b.vel.z++
				a.vel.z--
			} else if a.pos.z != b.pos.z {
				a.vel.z++
				b.vel.z--
			}
		}
	}

	//move
	for _, moon := range moons {
		moon.pos.x += moon.vel.x
		moon.pos.y += moon.vel.y
		moon.pos.z += moon.vel.z
	}
}

func getInput() []*Moon {
	return []*Moon{
		{Point{1, 3, -11}, Point{0, 0, 0}},
		{Point{17, -10, -8}, Point{0, 0, 0}},
		{Point{-1, -15, 2}, Point{0, 0, 0}},
		{Point{12, -4, -4}, Point{0, 0, 0}},
	}
}

func energy(moons []*Moon) int {
	energy := 0
	for _, moon := range moons {
		pe := abs(moon.pos.x) + abs(moon.pos.y) + abs(moon.pos.z)
		ke := abs(moon.vel.x) + abs(moon.vel.y) + abs(moon.vel.z)
		energy += pe * ke
	}
	return energy
}

// hacky
func abs(x int) int {
	y := x >> 63
	return (x ^ y) - y
}

// greatest common divisor (GCD) via Euclidean algorithm
func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func lcm(a, b int64, integers ...int64) int64 {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = lcm(result, integers[i])
	}

	return result
}

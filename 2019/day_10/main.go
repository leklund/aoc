package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

// Point in space
type Point struct {
	x int
	y int
}

// Target blow it up
type Target struct {
	point    Point
	dist     int
	quadrant int
	angle    float64
}

type Targets []Target

func (t Targets) Len() int {
	return len(t)
}
func (t Targets) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}
func (t Targets) Less(i, j int) bool {
	return t[i].angle < t[j].angle
}

func main() {
	asteroids := readFile("input.txt")
	asteroidMap := makeAsteroidMap(asteroids)
	bestPoint, count, visible := scanForBest(asteroidMap)

	fmt.Println("part one:")
	fmt.Println("Best point, count", bestPoint, count)

	// scan and destroy
	p := Point{}
	counter := 0
	for len(asteroidMap) > 1 {
		p = fireLasers(visible, asteroidMap, bestPoint, counter)

		if p != bestPoint {
			break
		}

		visible = scanFromPoint(bestPoint, asteroidMap)
	}
	fmt.Println("part two:")
	fmt.Println("200th point: ", p)
}

func fireLasers(asteroids map[Point]Target, asteroidMap map[Point]bool, origin Point, counter int) Point {
	targets := Targets{}

	for _, t := range asteroids {
		targets = append(targets, t)
	}
	sort.Sort(targets)

	for _, t := range targets {
		// FIRE
		if _, ok := asteroidMap[t.point]; ok {
			delete(asteroidMap, t.point)
			counter++
			if counter == 200 {
				return t.point
			}
		}

	}
	return origin
}

func makeAsteroidMap(s string) map[Point]bool {
	m := make(map[Point]bool)
	lines := strings.Split(s, "\n")

	for y, line := range lines {
		for x, c := range line {
			if string(c) == "#" {
				m[Point{x: x, y: y}] = true
			}
		}
	}
	return m
}

func scanForBest(asteroidMap map[Point]bool) (Point, int, map[Point]Target) {
	var best Point
	var count int
	var currVisible, visible map[Point]Target

	for origin := range asteroidMap {
		currVisible = scanFromPoint(origin, asteroidMap)

		if len(currVisible) > count {
			best = origin
			count = len(currVisible)
			visible = currVisible
		}
	}

	return best, count, visible
}

func scanFromPoint(origin Point, asteroidMap map[Point]bool) map[Point]Target {
	visible := make(map[Point]Target)

	for asteroid := range asteroidMap {
		dx, dy := asteroid.x-origin.x, origin.y-asteroid.y

		// it me
		if dx == 0 && dy == 0 {
			continue
		}

		dist := abs(dx) + abs(dy)
		q := 0

		if dx >= 0 && dy >= 0 {
			q = 0
		} else if dx >= 0 && dy <= 0 {
			q = 1
		} else if dx <= 0 && dy <= 0 {
			q = 3
		} else if dx <= 0 && dy >= 0 {
			q = 4
		}

		d := gcd(dx, dy)

		dx /= d
		dy /= d

		reduced := Point{x: dx, y: dy}
		angle := math.Atan2(float64(reduced.x), float64(reduced.y))

		if angle < 0 {
			angle = angle + (math.Pi * 2)
		}

		newT := Target{point: asteroid, dist: dist, quadrant: q, angle: angle}
		if targ, ok := visible[reduced]; ok {
			if dist < targ.dist {
				visible[reduced] = newT
			}
		} else {
			visible[reduced] = newT
		}
	}
	return visible
}

func readFile(path string) string {
	asteroids, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(asteroids)
}

// thanks Euclid
func gcd(x, y int) int {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return abs(x)
}

// hacky
func abs(x int) int {
	y := x >> 63
	return (x ^ y) - y
}

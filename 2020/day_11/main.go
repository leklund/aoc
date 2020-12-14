package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

type Seat struct {
	Point
	occupied bool
	next     bool
	floor    bool
}

type SeatMap map[Point]*Seat

var dirs = []Point{
	Point{-1, 0},
	Point{-1, -1},
	Point{0, -1},
	Point{1, -1},
	Point{1, 0},
	Point{1, 1},
	Point{0, 1},
	Point{-1, 1},
}

func main() {
	file := "input2.txt"

	lines := getLines(file)

	seatMap := MakeSeatMap(lines)
	LoadPlane(seatMap, false)

	//seatMap.print()
	fmt.Println(seatMap.passengerCount())

	// seatMap2 := MakeSeatMap(lines)
	// LoadPlane(seatMap2, true)

	// //seatMap2.print()
	// fmt.Println(seatMap2.passengerCount())
}

func LoadPlane(smap SeatMap, useLos bool) {
	var images []*image.Paletted
	var delays []int
	images = append(images, smap.toImage())
	delays = append(delays, 0)
	for {
		stable := true
		var adj, maxAdj int

		if useLos {
			maxAdj = 5
		} else {
			maxAdj = 4
		}

		for _, seat := range smap {
			if seat.floor {
				continue
			}
			if useLos {
				adj = seat.adjacentLos(smap)
			} else {
				adj = seat.adjacent(smap)
			}

			if !seat.occupied && adj == 0 {
				seat.next = true
				stable = false
			} else if seat.occupied && adj >= maxAdj {
				seat.next = false
				stable = false
			}
		}
		if stable {
			f, _ := os.OpenFile("giffy.gif", os.O_WRONLY|os.O_CREATE, 0600)
			defer f.Close()
			gif.EncodeAll(f, &gif.GIF{
				Image: images,
				Delay: delays,
			})
			return
		}

		for _, seat := range smap {
			seat.occupied = seat.next
		}
		// visualization:
		// smap.printScreen()
		// time.Sleep(25000000)
		images = append(images, smap.toImage())
		delays = append(delays, 0)
	}
}

func (sm SeatMap) passengerCount() int {
	c := 0
	for _, s := range sm {
		if s.occupied {
			c++
		}
	}
	return c
}

func (sm SeatMap) print() {
	var w, h int
	for _, s := range sm {
		w = max(w, s.x)
		h = max(h, s.y)
	}

	for i := 0; i <= w; i++ {
		for j := 0; j <= h; j++ {
			seat := sm[Point{i, j}]
			if seat.floor {
				fmt.Print(".")
			} else if seat.occupied {
				fmt.Print("#")
			} else {
				fmt.Print("L")

			}

		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func (sm SeatMap) printScreen() {
	var w, h int
	for _, s := range sm {
		w = max(w, s.x)
		h = max(h, s.y)
	}

	screen := &strings.Builder{}
	screen.WriteString("\033c")

	for i := 0; i <= w; i++ {
		for j := 0; j <= h; j++ {
			seat := sm[Point{i, j}]
			if seat.floor {
				screen.WriteString(".")
			} else if seat.occupied {
				screen.WriteString("#")
			} else {
				screen.WriteString("L")

			}

		}
		screen.WriteRune('\n')
	}
	screen.WriteRune('\n')
	fmt.Print(screen.String())

}

func (seat *Seat) adjacent(smap SeatMap) int {
	count := 0

	for _, d := range dirs {
		np := Point{seat.x + d.x, seat.y + d.y}
		if n, ok := smap[np]; ok {
			if n.occupied {
				count++
			}
		}
	}
	return count
}

func (seat *Seat) adjacentLos(smap SeatMap) int {
	count := 0

	for _, d := range dirs {
		dx, dy := 0, 0

		// This should be a recursive func but I'm tired
		for {
			dx += d.x
			dy += d.y
			np := Point{seat.x + dx, seat.y + dy}
			if n, ok := smap[np]; ok {
				if n.floor {
					continue
				} else {
					if n.occupied {
						count++
					}
					break
				}
			} else {
				break
			}
		}
	}
	return count
}

func MakeSeatMap(s []string) SeatMap {
	smap := make(map[Point]*Seat)
	for i, r := range s {
		for j, c := range r {
			p := Point{i, j}
			if c == 'L' {
				seat := &Seat{Point: p}
				smap[p] = seat
			} else if c == '.' {
				seat := &Seat{Point: p, floor: true}
				smap[p] = seat
			}
		}
	}

	return smap
}

func (sm SeatMap) toImage() *image.Paletted {
	var w, h int
	for _, s := range sm {
		w = max(w, s.x)
		h = max(h, s.y)
	}

	var palette = []color.Color{
		color.RGBA{0x00, 0x00, 0x00, 0xff}, // black
		color.RGBA{0x00, 0x00, 0xff, 0xff}, // blue
		color.RGBA{0x00, 0xff, 0x00, 0xff}, // green
		color.RGBA{0x00, 0xff, 0xff, 0xff}, // teal
		color.RGBA{0xff, 0x00, 0x00, 0xff}, // red
		color.RGBA{0xff, 0x00, 0xff, 0xff}, // pink
		color.RGBA{0xff, 0xff, 0x00, 0xff}, // yellow
		color.RGBA{0xff, 0xff, 0xff, 0xff}, // white
	}

	img := image.NewPaletted(image.Rect(0, 0, (w+1)*10, (h+1)*10), palette)

	for _, seat := range sm {
		c := color.RGBA{0x00, 0x00, 0x00, 0xff}

		if !seat.floor && seat.occupied {
			c = color.RGBA{0xff, 0xff, 0x00, 0xff}
		} else if !seat.floor && !seat.occupied {
			c = color.RGBA{0x00, 0x00, 0xff, 0xff}
		}

		r := image.Rect((seat.x*10)+1, (seat.y*10)+1, (seat.x+1)*10, (seat.y+1)*10)
		draw.Draw(img, r, &image.Uniform{c}, image.Point{}, draw.Src)

		// Rect(img, seat.x, seat.y, seat.x+10, seat.y+10)
	}
	return img
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

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

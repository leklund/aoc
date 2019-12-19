package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"sort"
)

type Layer struct {
	pixels    [][]byte
	histogram map[byte]int
}

type Image []Layer

var colorMap = map[byte]color.RGBA{
	'0': color.RGBA{0, 0, 0, 0xff},          // black
	'1': color.RGBA{0xff, 0xff, 0xff, 0xff}, // white
	'2': color.RGBA{0xff, 0xff, 0xff, 0},    //transparent
}

var palette = color.Palette{color.RGBA{0, 0, 0, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0xff}, color.RGBA{0xff, 0xff, 0xff, 0}}

func main() {
	line := getLine("input.txt")

	img := generateImage(line, 25, 6)

	sortedImg := Image{}

	copy(sortedImg, img)
	sort.Slice(sortedImg, func(i, j int) bool {
		return sortedImg[i].histogram['0'] < sortedImg[j].histogram['0']
	})

	fmt.Println("---PART ONE---")
	fmt.Println(img[0].histogram['1'] * img[0].histogram['2'])

	fmt.Println("---Part Two---")
	fmt.Println("open image.gif and anim.gif")

	renderImage(img, "image.gif")
}

func renderImage(img Image, filename string) {
	flattened := flattenImage(img)

	rgb := makeImg(flattened)

	file, _ := os.Create(filename)
	o := gif.Options{}
	gif.Encode(file, &rgb, &o)
}

func makeImg(layer Layer) image.Paletted {
	height, width := len(layer.pixels), len(layer.pixels[0])

	origin := image.Point{0, 0}
	lowerRight := image.Point{width, height}

	rgb := image.NewPaletted(image.Rectangle{origin, lowerRight}, palette)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			rgb.Set(x, y, colorMap[layer.pixels[y][x]])
		}
	}

	return *rgb
}

func flattenImage(img Image) Layer {
	width, height := len(img[0].pixels[0]), len(img[0].pixels)
	var flat Layer
	flat.pixels = makePixels(width, height)
	images := []*image.Paletted{}
	delays := []int{}

	// top down
	for layerIdx, layer := range img {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if layer.pixels[i][j] != '2' && flat.pixels[i][j] == 0 {
					flat.pixels[i][j] = layer.pixels[i][j]
				}
			}
		}
		frame := makeImg(flat)
		images = append(images, &frame)
		delay := 2
		if layerIdx == len(img)-1 {
			delay = 10000
		}
		delays = append(delays, delay)
	}

	file, _ := os.Create("anim.gif")

	gif.EncodeAll(file, &gif.GIF{
		Image: images,
		Delay: delays,
	})
	return flat
}

func generateImage(line string, width, height int) Image {
	var img Image

	for len(line) > 0 {
		layer := Layer{
			histogram: make(map[byte]int),
			pixels:    makePixels(width, height),
		}
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				char := line[0]
				layer.pixels[i][j] = char
				layer.histogram[char]++
				line = line[1:]
			}
		}
		img = append(img, layer)
	}
	return img
}

func makePixels(w, h int) [][]byte {
	layer := make([][]byte, h)
	for i := range layer {
		layer[i] = make([]byte, w)
	}

	return layer
}

func getLine(path string) string {
	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	s := bufio.NewScanner(file)
	s.Split(bufio.ScanLines)

	s.Scan()
	return s.Text()
}

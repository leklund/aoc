package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Mappy struct {
	SrcStart int
	SrcEnd   int
	Offset   int
}
type Seed struct {
	Start int
	Range int
}

func Dest(name string, src int) int {
	for _, m := range SeedMap[name] {
		dst := m.Dest(src)
		if dst > 0 {
			return dst
		}
	}
	return src
}

func (m *Mappy) Dest(x int) int {
	if x >= m.SrcStart && x <= m.SrcEnd {
		return x + m.Offset
	} else {
		return -1
	}
}

var (
	SeedMap            = make(map[string][]*Mappy)
	SeedsOne           []int
	Seeds              []Seed
	SeedToSoil         = "SeedToSoil"
	SoilToFertilizer   = "SoilToFertilizer"
	FertilizerToWater  = "FertilizerToWater"
	WaterToLight       = "WaterToLight"
	LightToTemp        = "LightToTemp"
	TempToHumidity     = "TempToHumidity"
	HumidityToLocation = "HumidityToLocation"

	digit = regexp.MustCompile(`\d+`)
	mapRE = regexp.MustCompile(` map:`)
)

func main() {
	file := "/Users/leklund/projects/aoc/2023/05/input.txt"
	parseInput(getLines(file))

	fmt.Println("ONE: ", minLoc())

	fmt.Println("TWO: ", minLocSmart())
}

func minLocSmart() int {
	var min int
	for _, seed := range Seeds {
		for i := seed.Start; i < seed.Start+seed.Range; i++ {

			loc := SeedToLocation(i)
			if min == 0 {
				min = loc
			} else if loc < min {
				min = loc
			}
		}
	}
	return min
}

func minLoc() int {
	var min int
	for _, seed := range SeedsOne {
		loc := SeedToLocation(seed)
		if min == 0 {
			min = loc
		} else if loc < min {
			min = loc
		}

	}
	return min
}

func parseInput(lines []string) {
	var current string
	for i, line := range lines {
		if i == 0 {
			m := digit.FindAllString(line, -1)
			for _, x := range m {
				SeedsOne = append(SeedsOne, toI(x))
			}

			for i := 1; i < len(m); i += 2 {
				s := Seed{
					Start: toI(m[i-1]),
					Range: toI(m[i]),
				}
				Seeds = append(Seeds, s)
			}

		} else if len(line) == 0 {
			continue
		} else if mapRE.MatchString(line) {
			switch line {
			case "seed-to-soil map:":
				current = SeedToSoil
			case "soil-to-fertilizer map:":
				current = SoilToFertilizer
			case "fertilizer-to-water map:":
				current = FertilizerToWater
			case "water-to-light map:":
				current = WaterToLight
			case "light-to-temperature map:":
				current = LightToTemp
			case "temperature-to-humidity map:":
				current = TempToHumidity
			case "humidity-to-location map:":
				current = HumidityToLocation
			}
		} else {
			m := digit.FindAllString(line, -1)
			dst := toI(m[0])
			src := toI(m[1])
			rng := toI(m[2])

			mm := &Mappy{
				SrcStart: src,
				SrcEnd:   src + rng - 1,
				Offset:   dst - src,
			}

			SeedMap[current] = append(SeedMap[current], mm)
		}
	}
}

func SeedToLocation(seed int) int {
	return Dest(HumidityToLocation,
		Dest(TempToHumidity,
			Dest(LightToTemp,
				Dest(WaterToLight,
					Dest(FertilizerToWater,
						Dest(SoilToFertilizer,
							Dest(SeedToSoil, seed),
						),
					),
				),
			),
		),
	)
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

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		fmt.Printf("ERR -%s- %v\n", s, err)
		panic(err)
	}

	return i
}

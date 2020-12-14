package main

import (
	"strings"
	"testing"
)

var testCases = []struct {
	seats  string
	final  int
	final2 int
}{
	{
		`L.LL.LL.LL
LLLLLLL.LL
L.L.L..L..
LLLL.LL.LL
L.LL.LL.LL
L.LLLLL.LL
..L.L.....
LLLLLLLLLL
L.LLLLLL.L
L.LLLLL.LL`, 37, 26,
	},
}

func TestFunc(t *testing.T) {
	for _, tc := range testCases {
		lines := strings.Split(tc.seats, "\n")

		seatMap := MakeSeatMap(lines)
		LoadPlane(seatMap, false)

		if seatMap.passengerCount() != tc.final {
			t.Errorf("got: %d, expecting: %d", seatMap.passengerCount(), tc.final)
		}

		seatMap2 := MakeSeatMap(lines)
		LoadPlane(seatMap2, true)

		if seatMap2.passengerCount() != tc.final2 {
			t.Errorf("got: %d, expecting: %d", seatMap2.passengerCount(), tc.final2)
		}

	}
}

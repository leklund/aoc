package main

import (
	"sort"
	"testing"
)

var sample = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}
var s2 = []int{354867, 322795, 387458, 389019, 342751, 596297, 415398, 391769, 373553, 395672, 582578, 421581, 441837, 433015, 547570, 567526, 466769, 762572, 533914, 596841, 774496, 602534, 659853, 884628, 976087, 764632, 665546, 716304, 1029312, 1011695, 765322, 769225, 813350, 1417924, 817253, 854596, 1213116, 2189203, 2067712, 1136448, 1034295, 1418542, 1130755}

func TestFindErr(t *testing.T) {
	res := findErr(sample, 5)

	if res != 127 {
		t.Errorf("nope. expected 127 got %d", res)
	}

	res = findErr(s2, 25)
	if res != -1 {
		t.Errorf("nope.  expected -1 got %d", res)
	}
}

func TestFindCont(t *testing.T) {
	res := findErr(sample, 5)

	cont := findCont(sample, res)

	sort.Ints(cont)

	a := cont[0] + cont[len(cont)-1]
	if a != 62 {
		t.Errorf("expected 62, got %d (%v)", a, cont)
	}
}

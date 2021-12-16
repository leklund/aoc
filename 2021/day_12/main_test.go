package main

import (
	"strings"
	"testing"
)

var testCases = []struct {
	input      string
	pathcount  int
	pathcount2 int
}{
	{
		`start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
		10,
		36,
	},
	{
		`dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
		19,
		103,
	},
	{
		`fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
		226,
		3509,
	},
}

func TestPathCountOne(t *testing.T) {
	for _, tc := range testCases {
		lines := strings.Split(tc.input, "\n")
		cavemap := parse(lines)

		visit := make(map[string]int)
		paths := FindPaths(START, cavemap, visit, true)

		if len(paths) != tc.pathcount {
			t.Errorf("Expected %d, got %d paths", tc.pathcount, len(paths))
		}

	}
}

func TestPathCountTwo(t *testing.T) {
	for _, tc := range testCases {
		lines := strings.Split(tc.input, "\n")
		cavemap := parse(lines)

		visit := make(map[string]int)
		paths := FindPaths(START, cavemap, visit, false)

		if len(paths) != tc.pathcount2 {
			t.Errorf("Expected %d, got %d paths", tc.pathcount2, len(paths))
		}

	}
}

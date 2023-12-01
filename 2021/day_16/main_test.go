package main

import (
	"strings"
	"testing"
)

var testCases = []struct {
	input       string
	sumVersions int
}{
	{
		"8A004A801A8002F478",
		16,
	},
	{
		"620080001611562C8802118E34",
		12,
	},
	{
		"C0015000016115A2E0802F182340",
		23,
	},
	{
		"A0016C880162017C3686B18A3D4780",
		31,
	},
}
var testCases2 = []struct {
	input string
	res   int
}{
	{
		"C200B40A82",
		3,
	},
	{
		"04005AC33890",
		54,
	},
	{
		"880086C3E88112",
		7,
	},
	{
		"CE00C43D881120",
		9,
	},
	{
		"D8005AC2A8F0",
		1,
	},
	{
		"F600BC2D8F",
		0,
	},
	{
		"9C005AC2F8F0",
		0,
	},
	{
		"9C0141080250320F1802104A08",
		1,
	},
}

func TestOne(t *testing.T) {
	for _, tc := range testCases {
		bs := buildString(tc.input)

		pt := 0
		p := parsePacket(strings.Split(bs, ""), &pt)

		vs := p.versionSum()

		if vs != tc.sumVersions {
			t.Errorf("Expected %d, got %d", tc.sumVersions, vs)
		}
	}
}

func TestTwo(t *testing.T) {
	for _, tc := range testCases2 {
		bs := buildString(tc.input)

		pt := 0
		p := parsePacket(strings.Split(bs, ""), &pt)

		v := p.eval()

		if v != tc.res {
			t.Errorf("Expected %d, got %d", tc.res, v)
		}
	}
}

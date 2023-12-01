package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Packet struct {
	version    int
	typeId     int
	val        int
	subPackets []Packet
}

func main() {
	l := getLine("input.txt")
	binary := buildString(l)
	pointer := 0
	packet := parsePacket(strings.Split(binary, ""), &pointer)

	fmt.Println("PART ONE: version sum - ", packet.versionSum())

	fmt.Println("PART TWO: eval() - ", packet.eval())
}

func (p Packet) eval() int {
	var v int

	switch p.typeId {
	case 0: // sum
		for _, sp := range p.subPackets {
			v += sp.eval()
		}
	case 1: // product
		v = 1
		for _, sp := range p.subPackets {
			v *= sp.eval()
		}
	case 2: // min
		for i, sp := range p.subPackets {
			if i == 0 {
				v = sp.eval()
			} else {
				z := sp.eval()
				if z < v {
					v = z
				}
			}
		}
	case 3: // max
		for i, sp := range p.subPackets {
			if i == 0 {
				v = sp.eval()
			} else {
				z := sp.eval()
				if z > v {
					v = z
				}
			}
		}
	case 4:
		return p.val
	case 5: // >
		if p.subPackets[0].eval() > p.subPackets[1].eval() {
			v = 1
		}
	case 6: // <
		if p.subPackets[0].eval() < p.subPackets[1].eval() {
			v = 1
		}
	case 7: // ==
		if p.subPackets[0].eval() == p.subPackets[1].eval() {
			v = 1
		}
	}

	return v
}

func (p Packet) versionSum() int {
	v := p.version
	for _, sp := range p.subPackets {
		v += sp.versionSum()

	}
	return v
}

func buildString(input string) string {
	var out strings.Builder
	for _, c := range input {
		num, _ := strconv.ParseInt(string([]rune{c}), 16, 64)

		out.WriteString(fmt.Sprintf("%04b", num))
	}
	return out.String()
}

func parsePacket(bin []string, pointer *int) Packet {
	v, t := packetHeaders(bin, pointer)
	p := Packet{}

	if t == 4 {
		p = parseLiteral(bin, pointer)
	} else {
		p = parseOperator(bin, pointer)
	}

	p.version = v
	p.typeId = t

	return p
}

func packetHeaders(bin []string, pointer *int) (int, int) {
	v := btoi(strings.Join(bin[*pointer:*pointer+3], ""))
	*pointer += 3

	t := btoi(strings.Join(bin[*pointer:*pointer+3], ""))
	*pointer += 3

	return int(v), int(t)
}

func parseLiteral(bin []string, p *int) Packet {
	sval := []string{}
	last := false
	for !last {
		if bin[*p] == "0" {
			last = true
		}
		chunk := bin[*p+1 : *p+5]
		sval = append(sval, chunk...)
		*p += 5
	}
	val := btoi(strings.Join(sval, ""))

	return Packet{val: int(val)}
}

func parseOperator(bin []string, p *int) Packet {
	lenId, size := parseLength(bin, p)
	subs := []Packet{}

	if lenId == 0 {
		subPacketEnd := *p + size

		for *p < subPacketEnd {
			subs = append(subs, parsePacket(bin, p))
		}
	} else {
		for i := 0; i < size; i++ {
			subs = append(subs, parsePacket(bin, p))
		}
	}

	return Packet{subPackets: subs}

}

func parseLength(bin []string, p *int) (int, int) {
	if bin[*p] == "0" {
		*p++
		l := btoi(strings.Join(bin[*p:*p+15], ""))
		*p += 15
		return 0, l
	} else {
		*p++
		l := btoi(strings.Join(bin[*p:*p+11], ""))
		*p += 11
		return 1, l
	}
}

func btoi(s string) int {
	i, _ := strconv.ParseInt(s, 2, 64)
	return int(i)
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

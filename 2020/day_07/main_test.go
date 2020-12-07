package main

import (
	"strings"
	"testing"
)

var input = `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

var input2 = `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

func TestBagCount(t *testing.T) {
	rules := strings.Split(input, "\n")
	tree, tree2 := parse(rules)

	uniqs := make(map[string]interface{})
	uniqs = bagCount("shiny gold", tree, uniqs)

	count := len(uniqs)
	if count != 4 {
		t.Errorf("expected 4, got %d", count)
	}

	count = contentsBagCount("shiny gold", tree2) - 1

	if count != 32 {
		t.Errorf("expected 32, got %d", count)
	}

	rules = strings.Split(input2, "\n")
	_, tree2 = parse(rules)
	count = contentsBagCount("shiny gold", tree2) - 1
	if count != 126 {
		t.Errorf("expected 126, got %d", count)
	}

}

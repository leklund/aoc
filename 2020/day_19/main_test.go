package main

import (
	"fmt"
	"strings"
	"testing"
)

var testCases = []struct {
	rules    string
	messages string
	count    int
}{
	{
		`0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"`,
		`ababbb
bababa
abbbab
aaabbb
aaaabbb`,
		2,
	},
}

func TestEvalRule(t *testing.T) {
	for _, tc := range testCases {
		rules := strings.Split(tc.rules, "\n")
		messages := strings.Split(tc.messages, "\n")

		ruleMap := make(map[string]string)
		ruleMap["4"] = "a"
		ruleMap["5"] = "b"

		rule := evalRule(rules, 0, ruleMap)

		count := countMatches(messages, rule)
		fmt.Println(rule)

		if count != tc.count {
			t.Errorf("expected %d, got %d", tc.count, count)
		}
	}
}

func BenchmarkCount(b *testing.B) {

	for n := 0; n < b.N; n++ {
		messages := getLines("input.txt")
		rules := getLines("sorted_rules.txt")
		ruleMap := make(map[string]string)

		//
		//testing
		// ruleMap["1"] = "a"
		// ruleMap["14"] = "b"

		ruleMap["5"] = "a"
		ruleMap["92"] = "b"

		rule := evalRule(rules, 0, ruleMap)
		countMatches(messages, rule)
	}
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	messages := getLines("input.txt")
	rules := getLines("sorted_rules.txt")
	// messages := getLines("input_test.txt")
	// rules := getLines("sorted_rules_test.txt")

	reString := parseRules(rules)

	count := countMatches(messages, reString)

	fmt.Println("Part one:", count)
}

func countMatches(messages []string, matchString string) int {
	ms := "^" + matchString + "$"
	re := regexp.MustCompile(ms)

	count := 0

	for _, m := range messages {
		if re.MatchString(m) {
			count++
		}
	}

	return count
}

func parseRules(rules []string) string {
	ruleMap := make(map[string]string)

	//
	//testing
	// ruleMap["1"] = "a"
	// ruleMap["14"] = "b"

	ruleMap["5"] = "a"
	ruleMap["92"] = "b"

	rule := evalRule(rules, 0, ruleMap)

	return rule
}

func evalRule(rules []string, idx int, ruleMap map[string]string) string {
	out := ""
	rule := rules[idx]
	rs := strings.Split(rule, ": ")

	ruleId := rs[0]
	if c, ok := ruleMap[ruleId]; ok {
		return c

		// special cases for part 2
		// 8: 42 | 42 8
		// 11: 42 31 | 42 11 31
	} else if ruleId == "8" {
		out += evalRule(rules, 42, ruleMap) + "+"
	} else if ruleId == "11" {
		// recursive regexps would be real nice here :(
		out += "("

		for depth := 1; depth <= 4; depth++ {
			out += evalRule(rules, 42, ruleMap) + "{" + strconv.Itoa(depth) + "}" + evalRule(rules, 31, ruleMap) + "{" + strconv.Itoa(depth) + "}"
			if depth != 4 {
				out += "|"
			}
		}
		out += ")"
	} else {
		rss := strings.Split(rs[1], " | ")

		// left only
		if len(rss) == 1 {
			ruleList := strings.Split(rss[0], " ")

			for _, r := range ruleList {
				out += evalRule(rules, toI(r), ruleMap)
			}
			// left and right
		} else if len(rss) == 2 {
			out += "("

			lRuleList := strings.Split(rss[0], " ")
			rRuleList := strings.Split(rss[1], " ")

			for _, r := range lRuleList {
				out += evalRule(rules, toI(r), ruleMap)
			}

			out += "|"

			for _, r := range rRuleList {
				out += "(" + evalRule(rules, toI(r), ruleMap) + ")"
			}

			out += ")"
		}
	}

	ruleMap[ruleId] = out
	return out
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
		panic(err)
	}

	return i
}

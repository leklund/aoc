package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var reg = regexp.MustCompile(`(\s+|\\n)`)
var required = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func main() {
	file := getString("input.txt")

	rawPassports := strings.Split(file, "\n\n")
	var passports []map[string]string

	for _, p := range rawPassports {
		p = strings.TrimSuffix(p, "\n")
		passports = append(passports, parsePassports(p))
	}

	valid := ValidPassports(passports)
	fmt.Printf("Valid Count: %d \n", len(valid))

}

func parsePassports(p string) map[string]string {
	passport := make(map[string]string)
	fields := reg.Split(p, -1)

	for _, field := range fields {
		kv := strings.Split(field, ":")
		passport[kv[0]] = kv[1]
	}

	return passport
}

func ValidPassports(p []map[string]string) []map[string]string {
	var valids []map[string]string
	for _, pass := range p {
		if valid(pass) {
			valids = append(valids, pass)
		}
	}
	return valids
}

func valid(p map[string]string) bool {
	for _, k := range required {
		if val, ok := p[k]; !ok {
			return false
			// comment out this else for Part one
		} else {
			if !validField(k, val) {
				return false
			}
		}
	}
	return true
}

var fourDigits = regexp.MustCompile(`^\d{4}$`)
var validators = map[string]*regexp.Regexp{
	"byr": fourDigits,
	"iyr": fourDigits,
	"eyr": fourDigits,
	"hgt": regexp.MustCompile(`^(\d+)(cm|in)`),
	"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
	"ecl": regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`),
	"pid": regexp.MustCompile(`^\d{9}$`),
}

func validField(field, value string) bool {
	if v, ok := validators[field]; ok {
		if !v.MatchString(value) {
			return false
		}
	}

	switch field {
	case "byr":
		byr := toI(value)
		if byr >= 1920 && byr <= 2002 {
			return true
		}
	case "iyr":
		yr := toI(value)
		if yr >= 2010 && yr <= 2020 {
			return true
		}
	case "eyr":
		yr := toI(value)
		if yr >= 2020 && yr <= 2030 {
			return true
		}
	case "hgt":
		m := validators[field].FindStringSubmatch(value)
		hgt := toI(m[1])
		units := m[2]
		if units == "cm" {
			if hgt >= 150 && hgt <= 193 {
				return true
			}
		} else if units == "in" {
			if hgt >= 59 && hgt <= 76 {
				return true
			}
		}
	default:
		return true
	}
	return false
}

// boiler plate

func getString(path string) string {
	data, _ := ioutil.ReadFile(path)

	return string(data)
}

func toI(s string) int {
	i, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return i
}

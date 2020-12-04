package main

import "testing"

var testCases = []struct {
	field string
	value string
	valid bool
}{
	{"byr", "2002", true},
	{"byr", "2003", false},
	{"byr", "42", false},
	{"hgt", "60in", true},
	{"hgt", "190cm", true},
	{"hgt", "190in", false},
	{"hgt", "190", false},
}

func TestValidField(t *testing.T) {
	for _, tc := range testCases {
		valid := validField(tc.field, tc.value)

		if valid != tc.valid {
			t.Errorf("expected field %s:%s to be %v, got %v", tc.field, tc.value, tc.valid, valid)
		}
	}
}

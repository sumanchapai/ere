package cmd

import (
	"testing"
)

func TestParseKnock(t *testing.T) {
	type TestCase struct {
		candidate string
		expect    []int
	}
	tests := []TestCase{
		{candidate: "", expect: []int{}},
		{candidate: "1,", expect: []int{1}},
		{candidate: "1, ", expect: []int{1}},
		{candidate: " 1, 2, 5 ", expect: []int{1, 2, 5}},
		{candidate: " 1,2, 11, 39", expect: []int{1, 2, 11, 39}},
	}
	for _, test := range tests {
		got, gotErr := parseKnock(test.candidate)
		if gotErr != nil {
			panic(gotErr)
		}
		expect := test.expect
		if !SlicesEqual(got, expect) {
			t.Errorf("parseKnock(%v) expected %v got %v", test.candidate, expect, got)
		}
	}
}

func SlicesEqual[K comparable](p, q []K) bool {
	if len(p) != len(q) {
		return false
	}
	for i := range p {
		if p[i] != p[i] {
			return false
		}
	}
	return true
}

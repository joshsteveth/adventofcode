package main

import (
	"testing"
)

func TestIsCapital(t *testing.T) {
	testcases := []struct {
		input  string
		output bool
	}{
		{"a", false},
		{"z", false},
		{"A", true},
		{"Z", true},
	}

	for _, tc := range testcases {
		if tc.output != isCapital(tc.input) {
			t.Errorf("unexpected result. wanted %v got %v", tc.output, isCapital(tc.input))
		}
	}
}

func TestExtendWord(t *testing.T) {
	testcases := []struct {
		input  string
		output string
	}{
		{"a", "aA"},
		{"A", "Aa"},
		{"z", "zZ"},
		{"Z", "Zz"},
	}

	for _, tc := range testcases {
		if tc.output != extendWord(tc.input) {
			t.Errorf("unexpected result. wanted %v got %v", tc.output, extendWord(tc.input))
		}
	}
}

func TestToStr(t *testing.T) {
	testcases := []struct {
		input  int
		output string
	}{
		{0, "a"},
		{25, "z"},
	}

	for _, tc := range testcases {
		if tc.output != toStr(tc.input) {
			t.Errorf("unexpected result. wanted %v got %v", tc.output, toStr(tc.input))
		}
	}
}

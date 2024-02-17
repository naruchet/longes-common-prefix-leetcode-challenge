package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	testCase := []struct {
		name     string
		strs     []string
		expected string
	}{
		{name: "expected fl", strs: []string{"flower", "flow", "flight"}, expected: "fl"},
		{name: "expected empty string", strs: []string{"xxx", "yyy", "zzz"}, expected: ""},
		{name: "expected emtpy string", strs: []string{}, expected: ""},
	}

	for _, i := range testCase {
		t.Run(i.name, func(t *testing.T) {
			actual := longestCommonPrefix(i.strs)
			if actual != i.expected {
				t.Errorf("Expected %s but got %s", i.expected, actual)
			}
		})
	}

}

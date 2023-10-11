package levenshtein 

import (
    "testing"
)

func TestDistance(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"", "hello", 5},
		{"hello", "", 5},
		{"hello", "hello", 0},
		{"ab", "aa", 1},
		{"ab", "ba", 2},
		{"ab", "aaa", 2},
		{"bbb", "a", 3},
		{"kitten", "sitting", 3},
		{"distance", "difference", 5},
		{"levenshtein", "frankenstein", 6},
		{"resume and cafe", "resumes and cafes", 2},
		{"a very long string that is meant to exceed", "another very long string that is meant to exceed", 6},
	}
	for i, d := range tests {
		n := Distance(d.a, d.b)
		if n != d.want {
			t.Errorf("Test[%d]: Distance(%q,%q) returned %v, want %v",
				i, d.a, d.b, n, d.want)
		}
	}
}

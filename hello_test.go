package main

import "testing"

func TestIsEvenLength(t *testing.T) {
	cases := []struct {
		in   string
		want bool
	}{
		{"Hello, world", true},
		{"Hello, Mars", false},
		{"", true},
	}
	for _, c := range cases {
		got := isEvenLength(c.in)
		if got != c.want {
			t.Errorf("isEvenLength(%q) == %t, want %t", c.in, got, c.want)
		}
	}
}

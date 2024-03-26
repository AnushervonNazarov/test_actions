package main

import "testing"

var tests = []struct {
	name           string
	a, b, expected int
}{
	{
		"2+2",
		2, 2, 4,
	},
	{
		"3+3",
		3, 3, 6,
	},
}

func TestAdd(t *testing.T) {
	for _, tt := range tests {
		got := Add(tt.a, tt.b)
		if got != tt.expected {
			t.Logf("expected %v but got %v", tt.expected, got)
			t.Fail()
		}
	}
}

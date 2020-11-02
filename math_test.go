package main

import (
	"testing"
)

// TestBetweenZeroAndMax ... ensure that this function always returns a value between [0, input]
func TestBetweenZeroAndMax(t *testing.T) {
	input := 666
	got := BetweenZeroAndMax(input)

	if got < 0 {
		t.Errorf("BetweenZeroAndMax(%d) < 0, want %d >= 0", input, got)
	}

	if got > input {
		t.Errorf("BetweenZeroAndMax(%d) > %d, want %d <= %d", input, got, input, got)
	}
}

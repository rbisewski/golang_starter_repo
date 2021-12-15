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

// TestRound ... ensure that this function always rounds up/down correctly
func TestRound(t *testing.T) {
	input := 1.3
	got := Round(input)

	if got != 1 {
		t.Errorf("Round(%f) != 1, want %d == 1", input, got)
	}

	input = -2.6
	got = Round(input)

	if got != -2 {
		t.Errorf("Round(%f) != -2, want %d == -2", input, got)
	}
}

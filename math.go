package main

import (
	"math"
	"math/rand"
)

// TossCoin ... randomly returns "true" or "false"
func TossCoin() bool {
	return rand.Intn(100) > 50
}

// BetweenZeroAndMax ... Returns a random number between 0 and X
func BetweenZeroAndMax(maximum int) int {
	if maximum < 1 {
		return 0
	}
	return rand.Intn(maximum)
}

// Min ... Get the minimum of a list of int values (i.e. the lowest value)
func Min(a ...int) int {

	// Use bitshifting to get the largest possible value for a
	// golang integer.
	min := int(^uint(0) >> 1)

	for _, i := range a {

		// Smaller? Set the variable then.
		if i < min {
			min = i
		}
	}

	return min
}

// Max ... Get the max of a list of int values (i.e. the highest value)
func Max(a ...int) int {

	// Use bitshifting the smallest possible value for a golang int.
	max := -(int(^uint(0)>>1) - 1)

	for _, i := range a {

		// Bigger? Then use that one.
		if i > max {
			max = i
		}
	}

	return max
}

// Percent ... Return the x-percentage of a given value.
func Percent(percent, of int) int {
	return of * percent / 100
}

// Round a given value or float to an integer.
func Round(x float64) int {

	// Define the precise (e.g. 1 == int)
	prec := float64(1)

	// Set the X^Y power based on the earlier defined precision
	pow := math.Pow(10, prec)

	// Determine the intermediate
	intermed := x * pow

	// Grab the modulo
	_, frac := math.Modf(intermed)

	// Increment the intermediate (i.e. between 0 and 1 is 0.5)
	intermed += .5

	// If the fraction is negative, invert it.
	x = .5
	if frac < 0.0 {
		x = -.5
		intermed--
	}

	// Fraction 0.5 or more? Ceiling then.
	rounder := math.Floor(intermed)
	if frac >= x {
		rounder = math.Ceil(intermed)
	}

	if pow == 0 {
		return 0
	}

	return int(rounder / pow)
}

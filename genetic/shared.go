package genetic

import (
	"math/rand"
	"time"
)

/**
Define a seeded random number generator for the library
*/
var murphy *rand.Rand = rand.New(
	rand.NewSource(
		time.Now().UnixNano()))

/**
getRandomWeightedFloat returns a float in [0.0, 1.0) that is more likely to be low than high
*/
func getRandomWeightedFloat() float64 {
	f1 := murphy.Float64()
	f2 := murphy.Float64()
	return f1 * f2
}

/**
max returns the maximum of two values
*/
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

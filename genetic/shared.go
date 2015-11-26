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

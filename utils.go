package main

import (
	"math/rand"
)

func RandInRange(r RangeInt) int {
	return rand.Intn(r.Max + 1 - r.Min) + r.Min
}

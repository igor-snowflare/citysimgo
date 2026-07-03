package main

import (
	"math/rand"
)

type RangeInt struct {
	Min int
	Max int
}

type WorldDefaults struct {
	population RangeInt
	areas RangeInt
}

func RandInRange(r RangeInt) int {
	return rand.Intn(r.Max + 1 - r.Min) + r.Min
}

type WorldConfig struct {
	population int
	areasCount int
}

func NewWorldDefaults() WorldDefaults {
	return WorldDefaults {
		population: RangeInt{Min: 400, Max: 10000},
		areas: RangeInt{Min: 10, Max:50},
	}
}

func NewWorldConfig(worldDefaults WorldDefaults) WorldConfig {
	return WorldConfig{
		population: RandInRange(worldDefaults.population),
		areasCount: RandInRange(worldDefaults.areas),
	}
}

package main

import (
	"fmt"
)

type Person struct {
	ageGroup string
	gender string
}

type Family struct {
	members []*Person
}

func generatePopulation(worldDefaults WorldDefaults) []*Family {
	var families []*Family

	populationCount := RandInRange(worldDefaults.population)

	for i := range populationCount {
		fmt.Printf("Genereting person number %d", i)
	}

	return families
}

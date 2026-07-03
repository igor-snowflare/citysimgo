package main

import (
	"fmt"
)

func main() {
	worldDefaults := NewWorldDefaults()
	worldConfig := NewWorldConfig(worldDefaults)
	fmt.Printf("Have a config with %d population and %d areas", worldConfig.population, worldConfig.areasCount)
}

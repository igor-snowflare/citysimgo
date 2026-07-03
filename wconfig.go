package main


type RangeInt struct {
	Min int
	Max int
}

type WorldDefaults struct {
	population RangeInt
	areas RangeInt
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

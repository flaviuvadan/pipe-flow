package main

import "fmt"

// main is main... runs main, will use for testing, I guess?
func main() {
	input := map[string][]float64{"a": {1, 2, 3}, "b": {4, 5, 6}}
	for k, v := range input {
		fmt.Printf("k: %v | v: %v\n", k, v)
	}

	var output []float64
	output = make([]float64, 10)
	for i, v := range output {
		fmt.Printf("i: %v | v: %v\n", i, v)
	}
}

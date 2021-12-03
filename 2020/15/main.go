package main

import "fmt"

func main() {
	run(input)
}

func run(input []int) {

	// A map from spoken numbers to the turn on which they were spoken
	var spoken = map[int]int{}

	for turn, i := range input {
		spoken[i] = turn
	}

	lastSpoken := input[len(input)-1]
	for turn := len(input); turn < 30000000; turn++ {
		var speak int
		ot, ok := spoken[lastSpoken]
		if !ok {
			speak = 0
		} else {
			speak = turn - ot - 1
		}
		spoken[lastSpoken] = turn - 1
		lastSpoken = speak
	}

	fmt.Println(lastSpoken)
}

var input = []int{1, 2, 16, 19, 18, 0}

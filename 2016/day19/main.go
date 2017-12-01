package main

import (
	"fmt"
)

var input = 3018458

func main() {
	// Track what's to the left of the current elf, not the number of presents
	elves := make([]int, input)

	for i := range elves {
		elves[i] = i + 1
	}
	elves[input-1] = 0

	current := 0
	numElves := input

	// Now find the opposite elf, and the previous opposite elf
	opposite := current
	prev := current
	for i := 0; i < numElves/2; i++ {
		prev = opposite
		opposite = elves[opposite]
	}

	for numElves > 1 {
		// Remove the opposite elf
		elves[prev] = elves[opposite]
		numElves--

		// Move on from left elf
		current = elves[current]

		// Move on the opposite elf, and remove it. If numElves is odd we do/do not
		// need to move
		if numElves&1 == 0 {
			prev = elves[prev]
		}
		opposite = elves[prev]

		if numElves%1000 == 0 {
			fmt.Printf("%d elves left\r", numElves)
		}
	}

	fmt.Printf("last elf at %d\n", current+1)
}

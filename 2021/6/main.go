package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if err := part1(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := part2(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func part1(input []int) error {
	var states timerStates

	for _, i := range input {
		states[i]++
	}

	for d := 0; d < 80; d++ {
		var next timerStates
		next[6] = states[0]
		next[8] = states[0]
		for i := range next[:8] {
			next[i] += states[i+1]
		}
		states = next
	}
	var total int
	for _, v := range states {
		total += v
	}

	fmt.Println(total)

	return nil
}

func part2(input []int) error {
	start := time.Now()
	var states timerStates

	for _, i := range input {
		states[i]++
	}

	for d := 0; d < 256; d++ {
		var next timerStates
		next[6] = states[0]
		next[8] = states[0]
		for i := range next[:8] {
			next[i] += states[i+1]
		}
		states = next
	}
	var total int
	for _, v := range states {
		total += v
	}

	fmt.Println(total, time.Since(start))

	return nil
}

// Number of fish in each timer state
type timerStates [9]int

var input = []int{3, 5, 1, 2, 5, 4, 1, 5, 1, 2, 5, 5, 1, 3, 1, 5, 1, 3, 2, 1, 5, 1, 1, 1, 2, 3, 1, 3, 1, 2, 1, 1, 5, 1, 5, 4, 5, 5, 3, 3, 1, 5, 1, 1, 5, 5, 1, 3, 5, 5, 3, 2, 2, 4, 1, 5, 3, 4, 2, 5, 4, 1, 2, 2, 5, 1, 1, 2, 4, 4, 1, 3, 1, 3, 1, 1, 2, 2, 1, 1, 5, 1, 1, 4, 4, 5, 5, 1, 2, 1, 4, 1, 1, 4, 4, 3, 4, 2, 2, 3, 3, 2, 1, 3, 3, 2, 1, 1, 1, 2, 1, 4, 2, 2, 1, 5, 5, 3, 4, 5, 5, 2, 5, 2, 2, 5, 3, 3, 1, 2, 4, 2, 1, 5, 1, 1, 2, 3, 5, 5, 1, 1, 5, 5, 1, 4, 5, 3, 5, 2, 3, 2, 4, 3, 1, 4, 2, 5, 1, 3, 2, 1, 1, 3, 4, 2, 1, 1, 1, 1, 2, 1, 4, 3, 1, 3, 1, 2, 4, 1, 2, 4, 3, 2, 3, 5, 5, 3, 3, 1, 2, 3, 4, 5, 2, 4, 5, 1, 1, 1, 4, 5, 3, 5, 3, 5, 1, 1, 5, 1, 5, 3, 1, 2, 3, 4, 1, 1, 4, 1, 2, 4, 1, 5, 4, 1, 5, 4, 2, 1, 5, 2, 1, 3, 5, 5, 4, 5, 5, 1, 1, 4, 1, 2, 3, 5, 3, 3, 1, 1, 1, 4, 3, 1, 1, 4, 1, 5, 3, 5, 1, 4, 2, 5, 1, 1, 4, 4, 4, 2, 5, 1, 2, 5, 2, 1, 3, 1, 5, 1, 2, 1, 1, 5, 2, 4, 2, 1, 3, 5, 5, 4, 1, 1, 1, 5, 5, 2, 1, 1}

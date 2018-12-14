package main

import (
	"bytes"
	"fmt"
)

func main() {

	part1(input)
	part2([]byte{5, 1, 5, 8, 9})
	part2([]byte{0, 1, 2, 4, 5})
	part2([]byte{9, 2, 5, 1, 0})
	part2([]byte{5, 9, 4, 1, 4})
	part2([]byte{3, 2, 7, 9, 0, 1})
}

const input = 327901

func part2(pattern []byte) {
	recipes := []byte{
		3, 7,
	}
	current := [2]int{0, 1}

	for {
		// add current recipes
		sum := recipes[current[0]] + recipes[current[1]]
		// Get the digits of this
		if sum > 9 {
			recipes = append(recipes, sum/10)
			if checkMatch(recipes, pattern) {
				return
			}

			sum = sum % 10
		}
		recipes = append(recipes, sum)

		if checkMatch(recipes, pattern) {
			return
		}

		for elf, offset := range current {
			current[elf] = (offset + int(recipes[offset]) + 1) % len(recipes)
		}

		// print(recipes, current)
	}
}

func checkMatch(recipes, pattern []byte) bool {
	lp := len(pattern)
	if lr := len(recipes); lr > lp {
		if bytes.Compare(recipes[lr-lp:], pattern) == 0 {
			fmt.Printf("%d\n", lr-lp)
			return true
		}
	}

	return false
}

func part1(numRecipes int) {
	recipes := []byte{
		3, 7,
	}
	current := [2]int{0, 1}

	for len(recipes) < numRecipes+10 {
		// add current recipes
		sum := recipes[current[0]] + recipes[current[1]]
		// Get the digits of this
		if sum > 9 {
			recipes = append(recipes, sum/10)
			sum = sum % 10
		}
		recipes = append(recipes, sum)

		for elf, offset := range current {
			current[elf] = (offset + int(recipes[offset]) + 1) % len(recipes)
		}

		// print(recipes, current)
	}
	for _, v := range recipes[numRecipes:] {
		fmt.Printf("%d", v)
	}
	fmt.Println()
}

func print(recipes []byte, current [2]int) {
	// Print the current round
	for i, r := range recipes {
		if i == current[0] {
			fmt.Printf("(%d)", r)
		} else if i == current[1] {
			fmt.Printf("[%d]", r)
		} else {
			fmt.Printf(" %d ", r)
		}
	}
	fmt.Println()
}

package main

import (
	"fmt"
	"os"
)

func nextLine(prev, next []bool) {
	var left, right bool
	ll := len(prev)
	for i := range prev {
		if i < ll-1 {
			right = prev[i+1]
		} else {
			right = false
		}

		next[i] = (left && !right) || (!left && right)

		left = prev[i]
	}
}

func parseLine(input string) []bool {
	output := make([]bool, len(input))
	for i, c := range input {
		if c == '^' {
			output[i] = true
		}
	}
	return output
}

func printLine(input []bool) {
	for _, trap := range input {
		if trap {
			os.Stdout.WriteString("^")
		} else {
			os.Stdout.WriteString(".")
		}
	}
	os.Stdout.WriteString("\n")

}

func countSafe(input []bool) int {
	safe := 0
	for _, trap := range input {
		if !trap {
			safe++
		}
	}
	return safe
}

func main() {
	var input = ".^^^.^.^^^^^..^^^..^..^..^^..^.^.^.^^.^^....^.^...^.^^.^^.^^..^^..^.^..^^^.^^...^...^^....^^.^^^^^^^"
	var rows = 400000

	line := parseLine(input)
	next := make([]bool, len(line))
	safe := countSafe(line)

	for i := 0; i < rows-1; i++ {
		nextLine(line, next)
		line, next = next, line
		safe += countSafe(line)
		// printLine(line)
	}

	fmt.Printf("%d safe\n", safe)
}

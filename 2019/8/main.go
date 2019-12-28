package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	run1(read())
	run2(read())
}

func run1(input string) {
	var dim = 25 * 6

	var minCounts [3]int
	minCounts[0] = dim
	for i := 0; i < len(input); i += dim {
		layer := input[i : i+dim]

		var counts [3]int
		for _, v := range layer {
			counts[v-'0']++
		}

		if counts[0] < minCounts[0] {
			minCounts = counts
		}
	}

	fmt.Println(minCounts[1] * minCounts[2])
}

func run2(input string) {
	var dim = 25 * 6
	final := make([]byte, dim)
	for i := range final {
		final[i] = 2
	}
	for i := 0; i < len(input); i += dim {
		layer := input[i : i+dim]
		for i, v := range layer {
			if final[i] == 2 {
				final[i] = byte(v) - '0'
			}
		}
	}

	// Try to make this visible!
	for i := 0; i < len(final); i += 25 {
		for j, v := range final[i : i+25] {
			if v == 1 {
				final[i+j] = '*'
			} else {
				final[i+j] = ' '
			}
		}
		os.Stdout.Write(final[i : i+25])
		os.Stdout.WriteString("\n")
	}
}

func read() string {
	data, err := ioutil.ReadFile("data")
	if err != nil {
		panic(err)
	}
	return string(data)
}

package main

import (
	"fmt"
	"sort"
)

func main() {
	run(input)
}

func run(input []int) {
	sort.Slice(input, func(i, j int) bool {
		return input[i] < input[j]
	})

	var diff1, diff3 int
	var lastv int
	for _, v := range input {
		switch diff := v - lastv; diff {
		case 3:
			diff3++
		case 1:
			diff1++
		}
		lastv = v
	}
	// Step from the greatest to the device is 3
	fmt.Println((diff3 + 1) * diff1)

	// How many variants do we have? The order will always be the same (as always has to increase). But we can drop out certain steps
	fmt.Println(validRemaining(0, input))
}

type key struct {
	start, len int
}

var cache = map[key]int{}

func validRemaining(start int, remaining []int) int {
	// memoize the results as otherwise we repeat the same calculations many
	// times
	if v, ok := cache[key{start: start, len: len(remaining)}]; ok {
		return v
	}
	if len(remaining) == 0 {
		return 1
	}
	var valid int
	for i, v := range remaining {
		if v-start <= 3 {
			valid += validRemaining(v, remaining[i+1:])
		} else {
			break
		}
	}
	cache[key{start: start, len: len(remaining)}] = valid
	return valid
}

var input = []int{
	46,
	63,
	21,
	115,
	125,
	35,
	89,
	17,
	116,
	90,
	51,
	66,
	111,
	142,
	148,
	60,
	2,
	50,
	82,
	20,
	47,
	24,
	80,
	101,
	103,
	16,
	34,
	72,
	145,
	141,
	124,
	14,
	123,
	27,
	62,
	61,
	95,
	138,
	29,
	7,
	149,
	147,
	104,
	152,
	22,
	81,
	11,
	96,
	97,
	30,
	41,
	98,
	59,
	45,
	88,
	37,
	10,
	114,
	110,
	4,
	56,
	122,
	139,
	117,
	108,
	91,
	36,
	146,
	131,
	109,
	31,
	75,
	70,
	140,
	38,
	121,
	3,
	28,
	118,
	54,
	107,
	84,
	15,
	76,
	71,
	102,
	130,
	132,
	87,
	55,
	129,
	83,
	23,
	42,
	69,
	1,
	77,
	135,
	128,
	94,
}

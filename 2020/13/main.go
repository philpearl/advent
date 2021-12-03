package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	bs := parseBusses(busses)
	run(ts, bs)

	// run3(parseBusses2("7,13,x,x,59,x,31,19"))
	run2(parseBusses2(busses))
}

func run(ts int, busses []int) {
	var at int = math.MaxInt64
	var atBus int
	for _, b := range busses {
		if cand := ((ts-1)/b)*b + b; cand < at {
			at = cand
			atBus = b
		}
	}
	fmt.Println(at, atBus, at-ts, (at-ts)*atBus)
}

type bus struct {
	offset int
	id     int
}

func run2(busses []int) {
	var bs []bus
	for i, b := range busses {
		if b == 0 {
			continue
		}
		bs = append(bs, bus{offset: i, id: b})
		fmt.Println(b, i)
	}

	start := 0
	step := bs[0].id

	for n := 2; n <= len(bs); n++ {
		start, step = findRepeat(start, step, bs[:n])
	}
	fmt.Println(start)
}

// Find where the repeating pattern for a subset of busses happens. Find the
// first instance and how often it happens. Assumes the bus IDs are coprime! (or
// more simply that there is a unique start and pattern repeat)
func findRepeat(start, step int, bs []bus) (ostart, ostep int) {

OUT:
	for t := start; ; t += step {
		for _, b := range bs {
			if (t+b.offset)%b.id != 0 {
				continue OUT
			}
		}
		if ostart == 0 {
			ostart = t
			fmt.Println(ostart)
		} else {
			ostep = t - ostart
			return ostart, ostep
		}
	}
}

func parseBusses(in string) []int {
	var busses []int
	for _, v := range strings.Split(in, ",") {
		if v == "x" {
			continue
		}
		i, _ := strconv.Atoi(v)
		busses = append(busses, i)
	}
	return busses
}

func parseBusses2(in string) []int {
	bs := strings.Split(in, ",")
	busses := make([]int, len(bs))
	for i, v := range bs {
		if v == "x" {
			continue
		}
		iv, _ := strconv.Atoi(v)
		busses[i] = iv
	}
	return busses
}

var ts = 1002578
var busses = `19,x,x,x,x,x,x,x,x,x,x,x,x,37,x,x,x,x,x,751,x,29,x,x,x,x,x,x,x,x,x,x,13,x,x,x,x,x,x,x,x,x,23,x,x,x,x,x,x,x,431,x,x,x,x,x,x,x,x,x,41,x,x,x,x,x,x,17`

/*
19
37, 13
19n = 37m + 13


*/

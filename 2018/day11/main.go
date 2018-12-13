package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println(powerLevel(122, 79, 57))
	fmt.Println(powerLevel(217, 196, 39))
	fmt.Println(powerLevel(101, 153, 71))

	start := time.Now()
	part1(5177)
	part2(5177)
	fmt.Println(time.Since(start))
}

func part2(gridSerial int) {
	// var x, y // x, y start at 1
	var maxX, maxY, maxPower, maxSize int
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {

			var totalPower int
			limit := 301
			if y > x {
				limit -= y
			} else {
				limit -= x
			}
			for size := 1; size <= limit; size++ {
				// As size goes up, we just add the outside line to the previous total. I.e. just the n cells
				// shown below.
				//
				//   p p n
				//   p p n
				//   n n n
				for dx := 0; dx < size; dx++ {
					totalPower += powerLevel(x+dx, y+size-1, gridSerial)
				}
				// Don't count the corner twice
				for dy := 0; dy < size-1; dy++ {
					totalPower += powerLevel(x+size-1, y+dy, gridSerial)
				}
				if totalPower > maxPower {
					maxX, maxY, maxPower, maxSize = x, y, totalPower, size
					// Let's see some progress
					fmt.Printf("%d,%d,%d = %d\n", maxX, maxY, maxSize, maxPower)
				}
			}
		}
	}

	// fmt.Printf("%d,%d,%d = %d\n", maxX, maxY, maxSize, maxPower)
}

func part1(gridSerial int) {
	// var x, y // x, y start at 1
	var maxX, maxY, maxPower int
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {

			var totalPower int
			for dx := 0; dx < 3; dx++ {
				for dy := 0; dy < 3; dy++ {
					totalPower += powerLevel(x+dx, y+dy, gridSerial)
				}
			}
			if totalPower > maxPower {
				maxX, maxY, maxPower = x, y, totalPower
			}
		}
	}

	fmt.Println(maxX, maxY, maxPower)
}

func powerLevel(x, y, gridSerial int) int {
	rackID := x + 10
	power := ((rackID * y) + gridSerial) * rackID

	// Want just the hundreds digit
	return ((power / 100) % 10) - 5
}

package main

import "fmt"

func main() {
	fmt.Printf("part1 severity=%d\n", getSeverity(input, 0))
	fmt.Printf("part2 delay=%d\n", part2(input))
}

type pair struct {
	depth     int
	cycleTime int
}

func part2(input map[int]int) int {
	scanners := make([]pair, 0, len(input))
	for depth, rang := range input {
		scanners = append(scanners,
			pair{
				depth:     depth,
				cycleTime: (rang - 1) * 2,
			})
	}

	workers := 4
	results := make(chan int, workers)
	for start := 0; start < workers; start++ {
		go worker(start, workers, scanners, results)
	}

	// Might not be 100%!
	result := <-results
	return result
}

func worker(start, stride int, scanners []pair, result chan<- int) {
	for delay := start; ; delay += stride {
		if !isHit(scanners, delay) {
			result <- delay
			return
		}
	}
}

func isHit(scanners []pair, delay int) bool {
	for _, scanner := range scanners {
		if ((scanner.depth + delay) % scanner.cycleTime) == 0 {
			return true
		}
	}
	return false
}

func getSeverity(input map[int]int, delay int) int {
	severity := 0
	for pico, rang := range input {
		cycleTime := (rang - 1) * 2
		caught := ((pico+delay)%cycleTime == 0)
		if caught {
			severity += pico * rang
		}
	}
	return severity
}

var input = map[int]int{
	0:  4,
	1:  2,
	2:  3,
	4:  5,
	6:  8,
	8:  4,
	10: 6,
	12: 6,
	14: 6,
	16: 10,
	18: 6,
	20: 12,
	22: 8,
	24: 9,
	26: 8,
	28: 8,
	30: 8,
	32: 12,
	34: 12,
	36: 12,
	38: 8,
	40: 10,
	42: 14,
	44: 12,
	46: 14,
	48: 12,
	50: 12,
	52: 12,
	54: 14,
	56: 14,
	58: 14,
	60: 12,
	62: 14,
	64: 14,
	68: 12,
	70: 14,
	74: 14,
	76: 14,
	78: 14,
	80: 17,
	82: 28,
	84: 18,
	86: 14,
}

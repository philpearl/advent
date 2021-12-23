package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 	input := `5483143223
	// 2745854711
	// 5264556173
	// 6141336146
	// 6357385478
	// 4167524645
	// 2176841721
	// 6882881134
	// 4846848554
	// 5283751526`

	if err := part1(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := part2(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func part1(input string) error {
	a := parseInput(input)

	var total int
	for i := 0; i < 100; i++ {
		a.incrAll()
		flashes := a.flash()
		total += flashes
	}
	fmt.Println(total)
	return nil
}

func part2(input string) error {
	a := parseInput(input)

	for i := 0; ; i++ {
		a.incrAll()
		flashes := a.flash()
		if flashes == a.w*a.h {
			fmt.Println(i + 1)
			break
		}
	}
	return nil
}

type point struct{ x, y int }

func (p point) plus(q point) point {
	return point{x: p.x + q.x, y: p.y + q.y}
}

type arena struct {
	w int
	h int
	d []uint8
}

func (a *arena) incrAll() {
	for i := range a.d {
		a.d[i]++
	}
}

func (a *arena) flash() int {
	flashed := make([]bool, len(a.d))
	var count int

	var current, next []int
	for i, v := range a.d {
		if v > 9 {
			flashed[i] = true
			a.d[i] = 0
			count++
			current = append(current, i)
		}
	}

	for len(current) > 0 {
		for _, i := range current {
			// Things in the queue are > 9, flashed is true, but they haven't
			// propagated yet. Do the propagation
			p := point{x: i % a.w, y: i / a.w}
			for x := -1; x < 2; x++ {
				for y := -1; y < 2; y++ {
					if x == 0 && y == 0 {
						continue
					}
					q := p.plus(point{x: x, y: y})
					if q.x < 0 || q.x >= a.w || q.y < 0 || q.y >= a.h {
						continue
					}
					j := q.x + q.y*a.w
					if flashed[j] {
						continue
					}

					a.d[j]++
					if a.d[j] > 9 {
						flashed[j] = true
						count++
						a.d[j] = 0
						next = append(next, j)
					}
				}
			}
		}
		current, next = next, current
		next = next[:0]
	}
	return count
}

func parseInput(input string) arena {
	lines := strings.Split(input, "\n")
	h := len(lines)
	w := len(lines[0])
	a := arena{
		w: w,
		h: h,
		d: make([]uint8, w*h),
	}

	for y, line := range lines {
		for x, c := range line {
			a.d[x+a.w*y] = uint8(c - '0')
		}
	}
	return a
}

var input = `4743378318
4664212844
2535667884
3273363861
2282432612
2166612134
3776334513
8123852583
8181786685
4362533174`

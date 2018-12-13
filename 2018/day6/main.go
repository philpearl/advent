package main

import (
	"fmt"
	"math"
)

func main() {

	part1(input)
	part2(input)
}

func part2(input []coord) {
	maxX, maxY := bounds(input)
	arena := make([]int, maxX*maxY)

	// mark each point in the arena with the input location that it is nearest to
	for _, xy := range input {
		for a := range arena {
			ac := coord{x: a % maxX, y: a / maxX}
			arena[a] += distance(xy, ac)
		}
	}

	var count int
	for _, c := range arena {
		if c < 10000 {
			count++
		}
	}
	fmt.Println(count)
}

func part1(input []coord) {
	type point struct {
		location int
		dist     int
	}

	maxX, maxY := bounds(input)
	arena := make([]point, maxX*maxY)
	for a := range arena {
		arena[a].dist = math.MaxInt64
	}

	// mark each point in the arena with the input location that it is nearest to
	for i, xy := range input {
		for a := range arena {
			ac := coord{x: a % maxX, y: a / maxX}
			d := distance(xy, ac)
			p := &arena[a]
			if d <= p.dist {
				if d == p.dist {
					p.location = -1
				} else {
					p.dist = d
					p.location = i
				}
			}
		}
	}

	counts := make([]int, len(input))

	for _, a := range arena {
		if a.location >= 0 {
			counts[a.location]++
		}
	}

	// any location that reaches the boundary is invalid.
	for i, a := range arena {
		x := i % maxX
		y := i / maxX
		if x == 0 || x == maxX-1 {
			if a.location >= 0 {
				counts[a.location] = 0
			}
		}
		if y == 0 || y == maxY-1 {
			if a.location >= 0 {
				counts[a.location] = 0
			}
		}
	}

	// fmt.Println(counts)
	var maxCount int
	for _, c := range counts {
		if c > maxCount {
			maxCount = c
		}
	}

	fmt.Println(maxCount)

	// for i := 0; i < len(arena); i += maxX {
	// 	fmt.Println(arena[i : i+maxX])
	// }
}

func distance(a, b coord) int {
	x := (a.x - b.x)
	if x < 0 {
		x = -x
	}
	y := (a.y - b.y)
	if y < 0 {
		y = -y
	}
	return x + y
}

func bounds(input []coord) (maxX, maxY int) {
	for _, c := range input {
		if c.x > maxX {
			maxX = c.x
		}
		if c.y > maxY {
			maxY = c.y
		}
	}
	return maxX + 1, maxY + 1
}

type coord struct {
	x, y int
}

var test = []coord{
	{1, 1},
	{1, 6},
	{8, 3},
	{3, 4},
	{5, 5},
	{8, 9},
}

var input = []coord{
	{152, 292},
	{163, 90},
	{258, 65},
	{123, 147},
	{342, 42},
	{325, 185},
	{69, 45},
	{249, 336},
	{92, 134},
	{230, 241},
	{74, 262},
	{241, 78},
	{299, 58},
	{231, 146},
	{239, 87},
	{44, 157},
	{156, 340},
	{227, 226},
	{212, 318},
	{194, 135},
	{235, 146},
	{171, 197},
	{160, 59},
	{218, 205},
	{323, 102},
	{290, 356},
	{244, 214},
	{174, 250},
	{70, 331},
	{288, 80},
	{268, 128},
	{359, 98},
	{78, 249},
	{221, 48},
	{321, 228},
	{52, 225},
	{151, 302},
	{183, 150},
	{142, 327},
	{172, 56},
	{72, 321},
	{225, 298},
	{265, 300},
	{86, 288},
	{78, 120},
	{146, 345},
	{268, 181},
	{243, 235},
	{262, 268},
	{40, 60},
}

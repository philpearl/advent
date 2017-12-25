package main

import "fmt"

func main() {
	part1(input)
}

func part1(cs []component) {
	b := bridge{}

	s, l := grow(b, cs)
	fmt.Printf("max strength is %d (%d)\n", s, l)
}

type bridge struct {
	end      int
	strength int
	length   int
	used     int64
}

func grow(b bridge, cs []component) (maxStrength, maxLength int) {
	maxStrength = b.strength
	maxLength = b.length
	// Look for possible next growths
	for i, c := range cs {
		if b.used&(1<<uint(i)) != 0 {
			continue
		}
		if c.a != b.end && c.b != b.end {
			continue
		}
		nb := b
		nb.length++
		nb.used |= 1 << uint(i)
		if c.a == b.end {
			nb.end = c.b
		} else {
			nb.end = c.a
		}
		nb.strength += c.a + c.b
		s, l := grow(nb, cs)
		if l >= maxLength {
			if l == maxLength {
				if s > maxStrength {
					maxStrength = s
				}
			} else {
				maxStrength = s
			}
			maxLength = l
		}
	}
	return maxStrength, maxLength
}

type component struct {
	a, b int
}

var input = []component{
	{32, 31},
	{2, 2},
	{0, 43},
	{45, 15},
	{33, 24},
	{20, 20},
	{14, 42},
	{2, 35},
	{50, 27},
	{2, 17},
	{5, 45},
	{3, 14},
	{26, 1},
	{33, 38},
	{29, 6},
	{50, 32},
	{9, 48},
	{36, 34},
	{33, 50},
	{37, 35},
	{12, 12},
	{26, 13},
	{19, 4},
	{5, 5},
	{14, 46},
	{17, 29},
	{45, 43},
	{5, 0},
	{18, 18},
	{41, 22},
	{50, 3},
	{4, 4},
	{17, 1},
	{40, 7},
	{19, 0},
	{33, 7},
	{22, 48},
	{9, 14},
	{50, 43},
	{26, 29},
	{19, 33},
	{46, 31},
	{3, 16},
	{29, 46},
	{16, 0},
	{34, 17},
	{31, 7},
	{5, 27},
	{7, 4},
	{49, 49},
	{14, 21},
	{50, 9},
	{14, 44},
	{29, 29},
	{13, 38},
	{31, 11},
}

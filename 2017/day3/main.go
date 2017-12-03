package main

import "fmt"

var input = 312051

/*
37                     31
38	17  16  15  14  13
	18   5   4   3  12
	19   6   1   2  11
	20   7   8   9  10
	21  22  23  24  25 26

	Bottom right-hand corner (at location (n, n)) is number (2n+1)**2


        4  3  2
        5  0  1
		6  7  8

		(1, 0) -> 1
		(1, 1) -> 2
		(0, 1) -> 3
		(-1, 1) -> 4

	mag of largest coord
*/

type coord struct {
	x int
	y int
}

func (c coord) abs() coord {
	if c.x < 0 {
		c.x = -c.x
	}
	if c.y < 0 {
		c.y = -c.y
	}
	return c
}

func (c coord) max() int {
	if c.x > c.y {
		return c.x
	}
	return c.y
}

func (c coord) toSpiral() int {
	layer := c.abs().max()
	// bottom right of this square has value
	// (2*layer + 1)**2
	maxVal := (2*layer + 1) * (2*layer + 1)
	// Which side are we on?
	switch {
	case c.y == -layer:
		// 9 - (1 - 1)
		// 9 - (1 - 0)
		// 9 - (1 -- 1)
		// 25 - (2 -- 2)
		return maxVal - (layer - c.x) - 1
	case c.x == -layer:
		// 9 - 2 - (layer + c.y)
		return maxVal - (2 * layer) - (layer + c.y) - 1
	case c.y == layer:
		// 9 - 4
		return maxVal - (4 * layer) - (layer + c.x) - 1
	case c.x == layer:
		return maxVal - (6 * layer) - (layer - c.y) - 1
	}
	fmt.Printf("bad coord %%v\n", c)
	return -1
}

func add(a, b coord) coord {
	return coord{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func (c coord) manhat() int {
	c = c.abs()
	return c.x + c.y
}

func rotate90(a coord) coord {
	return coord{
		x: -a.y,
		y: a.x,
	}
}

func rotateMinus90(a coord) coord {
	return coord{
		x: a.y,
		y: -a.x,
	}
}

type grid []int

func (g grid) valueAt(c coord) int {
	i := c.toSpiral()
	if i >= len(g) {
		return 0
	}
	if i < 0 {
		fmt.Printf("tospiral returned %d\n", i)
	}
	return g[i]
}

func main() {

	part1()
	part2()
}

func part2() {

	g := make(grid, 0, 200000)
	g = append(g, 1)
	pos := coord{0, 0}
	dirn := coord{1, 0}
	layer := 0

	for {
		// Move
		pos = add(pos, dirn)

		// Find the value for this position
		v := 0
		for x := -1; x <= 1; x++ {
			for y := -1; y <= 1; y++ {
				v += g.valueAt(add(pos, coord{x, y}))
			}
		}
		g = append(g, v)
		if v > input {
			fmt.Printf("value is %d\n", v)
			return
		}

		// Turn the corner
		if !(pos.x == layer && pos.y == -layer) {
			absPos := pos.abs()
			if absPos.x == absPos.y || absPos.max() > layer {
				dirn = rotate90(dirn)
				layer = absPos.max()
			}
		}
	}
}

func part1() {
	i := 0
	var value int
	for {
		value = (2*i + 1) * (2*i + 1)
		if value >= input {
			break
		}
		i++
	}

	side := 2*i + 1
	pos := coord{x: i, y: i}
	move := coord{x: -1, y: 0}
	moves := 0
	for value != input {
		pos = add(pos, move)
		value--
		moves++
		if moves == side {
			move = rotate90(move)
		}
	}

	fmt.Printf("dist is %d\n", pos.manhat())
}

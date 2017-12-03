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
*/

type coord struct {
	x int
	y int
}

func add(a, b coord) coord {
	return coord{
		x: a.x + b.x,
		y: a.y + b.y,
	}
}

func (c coord) manhat() int {
	if c.x < 0 {
		c.x = -c.x
	}
	if c.y < 0 {
		c.y = -c.y
	}
	return c.x + c.y
}

func rotate(a coord) coord {
	return coord{
		x: -a.y,
		y: a.x,
	}
}

func main() {
	i := 0
	var value int
	for {
		value = (2*i + 1) * (2*i + 1)
		fmt.Printf("(%d,%d) = %d\n", i, i, value)
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
			move = rotate(move)
		}
		fmt.Printf("(%d, %d) %d\n", pos.x, pos.y, value)
	}

	fmt.Printf("dist is %d\n", pos.manhat())
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	if err := part1("input.txt"); err != nil {
		fmt.Println(err)
	}
}

func part1(filename string) error {
	w, err := readFile(filename)
	if err != nil {
		return err
	}

	//w.print()
	for w.evolve() {
		//w.print()
	}

	return nil
}

type cart struct {
	x, y      int
	dx, dy    int
	direction crossChange
}

type track [][]byte

type world struct {
	track track
	carts []cart
}

type crossChange int

const (
	left crossChange = iota
	straight
	right
)

func (w *world) print() {
	lines := make([][]byte, len(w.track))
	for y, row := range w.track {
		line := make([]byte, len(row))
		copy(line, row)
		lines[y] = line
	}
	for _, cart := range w.carts {
		var sym byte
		if cart.dx == 0 {
			if cart.dy == 1 {
				sym = 'v'
			} else {
				sym = '^'
			}
		} else if cart.dx == 1 {
			sym = '>'
		} else {
			sym = '<'
		}
		lines[cart.y][cart.x] = sym
	}

	for _, l := range lines {
		fmt.Println(string(l))
	}
	fmt.Println()
}

func (w *world) evolve() bool {
	// sort carts by y then x
	sort.Slice(w.carts, func(i, j int) bool {
		a, b := w.carts[i], w.carts[j]
		if a.y == b.y {
			return a.x < b.x
		}
		return a.y < b.y
	})
	for i := range w.carts {
		c := &w.carts[i]

		cell := w.track[c.y][c.x]
		switch cell {
		case '|':
		case '-':
		case '/':
			// Hopelessly wrong - reflection!
			// v => <  0,1 => 1,0
			// > => ^  1,0 => 0,-1
			// < => v  -1,0 => 0, 1
			// ^ => >  0,-1 => 1, 0
			if c.dx == 0 {
				if c.dy == 1 {
					//  => v
					c.dx, c.dy = -1, 0
				} else {
					c.dx, c.dy = 1, 0
				}
			} else {
				if c.dx == 1 {
					//  => v
					c.dx, c.dy = 0, -1
				} else {
					c.dx, c.dy = 0, 1
				}
			}
		case '\\':
			if c.dx == 0 {
				if c.dy == 1 {
					//  => v
					c.dx, c.dy = +1, 0
				} else {
					c.dx, c.dy = -1, 0
				}
			} else {
				if c.dx == 1 {
					//  => v
					c.dx, c.dy = 0, 1
				} else {
					c.dx, c.dy = 0, -1
				}
			}

		case '+':
			switch c.direction {
			case left:
				// v 0, 1 => 1, 0 >
				// ^ 0, -1 => -1, 0 <
				// > 1, 0 => 0, -1 ^
				// < -1, 0 => 0, 1 v
				c.dx, c.dy = c.dy, -c.dx
			case straight:
			case right:
				c.dx, c.dy = -c.dy, c.dx

			}
			c.direction++
			if c.direction > right {
				c.direction = left
			}
		}
		c.x += c.dx
		c.y += c.dy

		// Check for collision
		for j, c2 := range w.carts {
			if i == j {
				continue
			}
			if c2.x == c.x && c2.y == c.y {
				fmt.Printf("collision at %d,%d\n", c.x, c.y)
				// remove cards i & j
				carts := w.carts[:0]
				for k, cart := range w.carts {
					if k != i && k != j {
						carts = append(carts, cart)
					}
				}
				w.carts = carts

				if len(w.carts) == 1 {
					fmt.Printf("last card at %d,%d", w.carts[0].x, w.carts[0].y)
					return false
				}
				return true
			}
		}
	}
	return true
}

func readFile(filename string) (w world, err error) {
	f, err := os.Open(filename)
	if err != nil {
		return w, err
	}

	var carts []cart
	var track track
	var lineLength int
	s := bufio.NewScanner(f)
	for s.Scan() {
		line := s.Bytes()
		if len(line) != lineLength {
			if lineLength == 0 {
				lineLength = len(line)
			} else {
				fmt.Printf("Line length %d does not match expected %d\n", len(line), lineLength)
			}
		}
		cells := make([]byte, lineLength)
		for i, c := range line {
			switch c {
			case '<':
				c = '-'
				carts = append(carts, cart{
					x:  i,
					y:  len(track),
					dx: -1,
					dy: 0,
				})
			case '>':
				c = '-'
				carts = append(carts, cart{
					x:  i,
					y:  len(track),
					dx: 1,
					dy: 0,
				})
			case '^':
				c = '|'
				carts = append(carts, cart{
					x:  i,
					y:  len(track),
					dx: 0,
					dy: -1,
				})
			case 'v':
				c = '|'
				carts = append(carts, cart{
					x:  i,
					y:  len(track),
					dx: 0,
					dy: 1,
				})
			}
			cells[i] = c
			// TODO some kind of matrix / vector??
		}
		track = append(track, cells)
	}

	w.track = track
	w.carts = carts

	return w, s.Err()
}

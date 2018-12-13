package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {

	if err := run("input.txt"); err != nil {
		fmt.Println(err)
	}
}

func run(filename string) error {
	w, err := readFile(filename)
	if err != nil {
		return err
	}

	// w.print()
	for w.evolve() {
		// w.print()
	}

	return nil
}

type cart struct {
	x, y      int
	direction byte
	cross     crossChange
	dead      bool
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
		if cart.dead {
			continue
		}
		lines[cart.y][cart.x] = cart.direction
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
		if c.dead {
			continue
		}

		cell := w.track[c.y][c.x]
		switch cell {
		case '|':
		case '-':
		case '/':
			switch c.direction {
			case '>':
				c.direction = '^'
			case '<':
				c.direction = 'v'
			case '^':
				c.direction = '>'
			case 'v':
				c.direction = '<'
			}
		case '\\':
			switch c.direction {
			case '>':
				c.direction = 'v'
			case '<':
				c.direction = '^'
			case '^':
				c.direction = '<'
			case 'v':
				c.direction = '>'
			}
		case '+':
			switch c.cross {
			case left:
				switch c.direction {
				case '>':
					c.direction = '^'
				case '<':
					c.direction = 'v'
				case '^':
					c.direction = '<'
				case 'v':
					c.direction = '>'
				}
			case straight:
			case right:
				switch c.direction {
				case '>':
					c.direction = 'v'
				case '<':
					c.direction = '^'
				case '^':
					c.direction = '>'
				case 'v':
					c.direction = '<'
				}
			}
			c.cross++
			if c.cross > right {
				c.cross = left
			}
		}

		switch c.direction {
		case '>':
			c.x++
		case '<':
			c.x--
		case '^':
			c.y--
		case 'v':
			c.y++
		}

		// Check for collision
		for j, c2 := range w.carts {
			if i == j || c2.dead {
				continue
			}
			if c2.x == c.x && c2.y == c.y {
				fmt.Printf("collision at %d,%d\n", c.x, c.y)
				c.dead = true
				w.carts[j].dead = true
			}
		}
	}

	var last *cart
	count := 0
	for cc := range w.carts {
		cart := &w.carts[cc]
		if !cart.dead {
			count++
			last = cart
		}
	}
	if count == 1 {
		fmt.Printf("last cart at %d,%d\n", last.x, last.y)
		return false
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
			case '<', '>', '^', 'v':
				carts = append(carts, cart{
					x:         i,
					y:         len(track),
					direction: c,
				})
				if c == '>' || c == '<' {
					c = '-'
				} else {
					c = '|'
				}
			}
			cells[i] = c
		}
		track = append(track, cells)
	}

	w.track = track
	w.carts = carts

	return w, s.Err()
}

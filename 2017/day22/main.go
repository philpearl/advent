package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// part1(10000000, `..#
	// #..
	// ...`)
	part1(10000000, input)
}

func part1(bursts int, input string) {

	s := state{
		d: direction{x: 0, y: 1},
		w: parseInput(input),
	}

	for i := 0; i < bursts; i++ {
		// s.print()
		s.step()
	}
	// s.print()

	fmt.Printf("%d steps caused infection\n", s.infectionCount)
}

func (s *state) step() {
	status := s.w.status(s.p)

	switch status {
	case clean:
		s.d.left()
	case weakened:
	case infected:
		s.d.right()
	case flagged:
		s.d.right()
		s.d.right()
	}
	if s.w.infect(status, s.p) {
		s.infectionCount++
	}
	s.p.move(s.d)
}

type status uint8

const (
	clean status = iota
	weakened
	infected
	flagged
)

type position struct{ x, y int }

type world map[position]status

func (w world) numInfected() int {
	return len(w)
}

func (w world) status(p position) status {
	if status, ok := w[p]; ok {
		return status
	}
	return clean
}

func (w world) infect(s status, p position) bool {
	switch s {
	case clean:
		w[p] = weakened
	case weakened:
		w[p] = infected
		return true
	case infected:
		w[p] = flagged
	case flagged:
		delete(w, p)
	}
	return false
}

type direction struct{ x, y int }

type state struct {
	w              world
	d              direction
	p              position
	infectionCount int
}

func (s *state) print() {
	for y := 4; y > -4; y-- {
		for x := -4; x < 5; x++ {
			b := s.status(position{x: x, y: y})
			os.Stdout.Write(b[:])
		}
		os.Stdout.WriteString("\n")
	}
	os.Stdout.WriteString("\n")
}

func (s *state) status(p position) [3]byte {
	var b = [3]byte{' ', '.', ' '}
	if p == s.p {
		b[0] = '['
		b[2] = ']'
	}
	status := s.w.status(p)
	switch status {
	case weakened:
		b[1] = 'W'
	case infected:
		b[1] = '#'
	case flagged:
		b[1] = 'F'
	}
	return b
}

func (p *position) move(d direction) {
	p.x += d.x
	p.y += d.y
}

func (d *direction) left() {
	d.x, d.y = -d.y, d.x
}

func (d *direction) right() {
	d.x, d.y = d.y, -d.x
}

func parseInput(in string) (w world) {
	w = make(world)
	lines := strings.Split(in, "\n")
	height := len(lines)
	width := len(lines[0])
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				w.infect(weakened, position{
					x: x - width/2,
					y: height/2 - y,
				})
			}
		}
	}
	return w
}

var input = `#.....##.####.#.#########
.###..#..#..####.##....#.
..#########...###...####.
.##.#.##..#.#..#.#....###
...##....###..#.#..#.###.
###..#...######.####.#.#.
#..###..###..###.###.##..
.#.#.###.#.#...####..#...
##........##.####..##...#
.#.##..#.#....##.##.##..#
###......#..##.####.###.#
....#..###..#######.#...#
#####.....#.##.#..#..####
.#.###.#.###..##.#..####.
..#..##.###...#######....
.#.##.#.#.#.#...###.#.#..
##.###.#.#.###.#......#..
###..##.#...#....#..####.
.#.#.....#..#....##..#..#
#####.#.##..#...##..#....
##..#.#.#.####.#.##...##.
..#..#.#.####...#........
###.###.##.#..#.##.....#.
.##..##.##...#..#..#.#..#
#...####.#.##...#..#.#.##`

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	part1(readInput("test.txt"))
	part1(readInput("day19input.txt"))
}

func part1(d diagram) {
	start := d.findStart()
	s := state{
		position:  start,
		direction: vec{x: 0, y: 1},
	}
	letters, steps := s.follow(d)
	fmt.Printf("letters are %s\n", letters)
	fmt.Printf("steps are %d\n", steps)
}

type state struct {
	position  vec
	direction vec
	letters   []byte
}

func (s *state) follow(d diagram) (string, int) {
	steps := 0
	for {
		c := d.symbol(s.position)
		if !symbolOKForMovement(c) {
			// Either something is wrong or we've made it to the end
			break
		}
		switch {
		case c == '+':
			// Need to change direction
			s.direction.rotate()
			next := s.position.add(s.direction)
			c = d.symbol(next)
			// fmt.Printf("turning, consider %c\n", c)
			if !symbolOKForDirection(c, s.direction) {
				// fmt.Printf("reverse\n")
				s.direction.reverse()
			}
		case c >= 'A' && c <= 'Z':
			// Pick up letter and continue
			s.letters = append(s.letters, c)
		}
		s.position = s.position.add(s.direction)
		steps++
	}

	return string(s.letters), steps
}

func symbolOKForMovement(c byte) bool {
	return (c == '|') ||
		(c == '-') ||
		(c == '+') ||
		(c >= 'A' && c <= 'Z')
}

func symbolOKForDirection(c byte, direction vec) bool {
	return (c == '|' && direction.x == 0) ||
		(c == '-' && direction.y == 0) ||
		(c == '+') ||
		(c >= 'A' && c <= 'Z')
}

type vec struct {
	x, y int
}

func (v *vec) add(m vec) vec {
	return vec{x: v.x + m.x, y: v.y + m.y}
}

func (v *vec) rotate() {
	v.x, v.y = v.y, -v.x
}

func (v *vec) reverse() {
	v.x = -v.x
	v.y = -v.y
}

type diagram []string

func (d diagram) symbol(p vec) byte {
	if p.y >= len(d) {
		return 0
	}
	r := d[p.y]
	if p.x >= len(r) {
		return 0
	}
	return r[p.x]
}

func (d diagram) findStart() vec {
	for i, c := range d[0] {
		if c == '|' {
			return vec{x: i, y: 0}
		}
	}
	return vec{}
}

func readInput(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type vector struct {
	x int
	y int
}

func (v *vector) manhatan() int {
	return v.x + v.y
}

type direction int

const (
	north direction = iota
	west
	south
	east
)

type path struct {
	direction direction
	vector    vector
	visited   map[vector]struct{}
}

func newPath(initial direction) *path {
	return &path{
		direction: initial,
		visited:   map[vector]struct{}{},
	}
}

type turn rune

const (
	L turn = 'L'
	R turn = 'R'
)

func (p *path) step(turn turn, steps int) bool {
	switch turn {
	case L:
		p.direction++
	case R:
		p.direction--
	}
	if p.direction < north {
		p.direction += 4
	}
	if p.direction > east {
		p.direction -= 4
	}

	for i := 0; i < steps; i++ {
		switch p.direction {
		case north:
			p.vector.y++
		case west:
			p.vector.x--
		case south:
			p.vector.y--
		case east:
			p.vector.x++
		}
		if p.haveVisited() {
			return true
		}
	}
	return false
}

func (p *path) haveVisited() bool {
	_, ok := p.visited[p.vector]
	if ok {
		return true
	}
	p.visited[p.vector] = struct{}{}
	return false
}

func (p *path) print(w io.Writer) {
	fmt.Fprintf(w, "(%d, %d)", p.vector.x, p.vector.y)
	switch p.direction {
	case north:
		io.WriteString(w, "^")
	case west:
		io.WriteString(w, "<")
	case south:
		io.WriteString(w, "v")
	case east:
		io.WriteString(w, ">")
	}
	io.WriteString(w, "\n")
}

func (p *path) distance() int {
	return p.vector.manhatan()
}

func parseStep(instruction string) (tur turn, steps int) {
	tur = turn(rune(instruction[0]))
	stepss, _ := strconv.ParseInt(instruction[1:], 10, 64)
	steps = int(stepss)
	return
}

func main() {

	input := `L5, R1, R4, L5, L4, R3, R1, L1, R4, R5, L1, L3, R4, L2, L4, R2, L4, L1, R3, R1, R1, L1, R1, L5, R5, R2, L5, R2, R1, L2, L4, L4, R191, R2, R5, R1, L1, L2, R5, L2, L3, R4, L1, L1, R1, R50, L1, R1, R76, R5, R4, R2, L5, L3, L5, R2, R1, L1, R2, L3, R4, R2, L1, L1, R4, L1, L1, R185, R1, L5, L4, L5, L3, R2, R3, R1, L5, R1, L3, L2, L2, R5, L1, L1, L3, R1, R4, L2, L1, L1, L3, L4, R5, L2, R3, R5, R1, L4, R5, L3, R3, R3, R1, R1, R5, R2, L2, R5, L5, L4, R4, R3, R5, R1, L3, R1, L2, L2, R3, R4, L1, R4, L1, R4, R3, L1, L4, L1, L5, L2, R2, L1, R1, L5, L3, R4, L1, R5, L5, L5, L1, L3, R1, R5, L2, L4, L5, L1, L1, L2, R5, R5, L4, R3, L2, L1, L3, L4, L5, L5, L2, R4, R3, L5, R4, R2, R1, L5`

	w := os.Stdout
	p := newPath(north)

	for _, instruction := range strings.Split(input, ", ") {
		turn, steps := parseStep(instruction)
		if p.step(turn, steps) {
			fmt.Printf("Final distance is %d\n", p.distance())
			break
		}
		p.print(w)
	}
}

package main

import (
	"container/list"
	"fmt"

	"github.com/steakknife/hamming"
)

func isWall(x, y int) bool {
	magic := x*x + 3*x + 2*x*y + y + y*y + 1364
	// count bits
	bits := hamming.CountBitsUint64(uint64(magic))
	return bits&1 == 1
}

type coord struct{ x, y, depth int }
type coordWithDepth struct {
	coord
	depth int
}

type problem struct {
	l       *list.List
	visited map[coord]struct{}
}

func newProblem() *problem {
	return &problem{
		l:       list.New(),
		visited: make(map[coord]struct{}),
	}
}

func (p *problem) addMove(c coordWithDepth) {
	if c.x < 0 || c.y < 0 {
		return
	}
	if isWall(c.x, c.y) {
		return
	}

	// if p.hasVisited(c) {
	// 	return
	// }
	p.l.PushBack(c)
}

func (p *problem) hasVisited(c coord) bool {
	if _, visited := p.visited[c]; visited {
		return true
	}
	p.visited[c] = struct{}{}
	return false
}

func (p *problem) visitCount() int {
	return len(p.visited)
}

func (p *problem) addPossibleMoves(x, y, depth int) {
	p.addMove(coordWithDepth{coord: coord{x: x, y: y + 1}, depth: depth})
	p.addMove(coordWithDepth{coord: coord{x: x, y: y - 1}, depth: depth})
	p.addMove(coordWithDepth{coord: coord{x: x + 1, y: y}, depth: depth})
	p.addMove(coordWithDepth{coord: coord{x: x - 1, y: y}, depth: depth})
}

func (p *problem) nextCoord() (coordWithDepth, bool) {
	e := p.l.Front()
	if e == nil {
		return coordWithDepth{}, true
	}
	p.l.Remove(e)
	return e.Value.(coordWithDepth), false
}

func main() {
	prob := newProblem()

	prob.addMove(coordWithDepth{coord: coord{x: 1, y: 1}})

	for {
		c, empty := prob.nextCoord()
		if empty {
			break
		}
		if c.depth == 51 {
			break
		}
		if prob.hasVisited(c.coord) {
			continue
		}
		// if c.x == 31 && c.y == 39 {
		// 	fmt.Printf("Reached at depth %d\n", c.depth)
		// 	break
		// }
		prob.addPossibleMoves(c.x, c.y, c.depth+1)
	}

	fmt.Println(prob.visitCount())
}

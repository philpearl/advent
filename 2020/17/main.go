package main

import (
	"fmt"
	"strings"
)

func main() {
	run(input)
}

func run(input string) {
	w := new()
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				w.active[coord{x: x, y: y}] = struct{}{}
			}
		}
	}

	for i := 0; i < 6; i++ {
		w.step()
	}

	fmt.Println(w.count())
}

type coord struct {
	x, y, z, w int
}

type world struct {
	active map[coord]struct{}
	next   map[coord]struct{}
}

func new() world {
	return world{
		active: make(map[coord]struct{}),
		next:   make(map[coord]struct{}),
	}
}

func (w world) count() int {
	return len(w.active)
}

func (w *world) step() {

	// if an active cell has 2 or 3 active neighbours it stays active
	for p := range w.active {
		c := w.activeNeighbours(p)
		if c == 2 || c == 3 {
			w.next[p] = struct{}{}
		}
	}
	// We also want to look at inactive cells. But we don't track those! Only
	// neighbours of active cells are interesting. This is probably inefficient as
	// we look at many cells multiple times!
	for p := range w.active {
		var q coord
		for _, dx := range deltas {
			q.x = p.x + dx
			for _, dy := range deltas {
				q.y = p.y + dy
				for _, dz := range deltas {
					q.z = p.z + dz
					for _, dw := range deltas {
						if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
							continue
						}
						q.w = p.w + dw
						if _, active := w.active[q]; active {
							continue
						}
						if w.activeNeighbours(q) == 3 {
							w.next[q] = struct{}{}
						}
					}
				}
			}
		}
	}

	w.next, w.active = w.active, w.next
	for k := range w.next {
		delete(w.next, k)
	}
}

var deltas = []int{-1, 0, 1}

func (w world) activeNeighbours(p coord) int {
	var count int
	var q coord
	for _, dx := range deltas {
		q.x = p.x + dx
		for _, dy := range deltas {
			q.y = p.y + dy
			for _, dz := range deltas {
				q.z = p.z + dz
				for _, dw := range deltas {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					q.w = p.w + dw
					if _, active := w.active[q]; active {
						count++
					}
				}
			}
		}
	}
	return count
}

var input = `######.#
#.###.#.
###.....
#.####..
##.#.###
.######.
###.####
######.#`

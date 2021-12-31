package main

import (
	"fmt"
	"os"
)

func main() {
	// oops did both parts in part1
	if err := part1(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := part2(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func part1(input string) error {
	w := parseInput(input)
	var hits, maxHitY int
	for x := 1; x < 4000; x++ {
		for y := -400; y < 4000; y++ {
			w.pos = point{}
			w.velocity.x = x
			w.velocity.y = y
			var max int
			for !w.willNeverHit() {
				w.step()
				if w.pos.y > max {
					max = w.pos.y
				}
				// fmt.Println(w.pos)
				if w.onTarget() {
					fmt.Println("hit target!", x, y)
					hits++
					if max > maxHitY {
						maxHitY = max
					}
					break
				}
			}
		}
	}
	fmt.Println(hits, maxHitY)
	return nil
}

func part2(input string) error {
	// World is extended
	w := parseInput(input)

	_ = w
	return nil
}

type point struct{ x, y int }

func (p *point) add(q point) {
	p.x += q.x
	p.y += q.y
}

type world struct {
	bl       point
	tr       point
	pos      point
	velocity point
}

func (w *world) step() {
	w.pos.add(w.velocity)

	if w.velocity.x > 0 {
		w.velocity.x--
	} else if w.velocity.x < 0 {
		w.velocity.x++
	}
	w.velocity.y--
}

func (w *world) onTarget() bool {
	return w.pos.x >= w.bl.x && w.pos.x <= w.tr.x &&
		w.pos.y >= w.bl.y && w.pos.y <= w.tr.y
}

func (w *world) willNeverHit() bool {
	if w.velocity.x == 0 && (w.pos.x < w.bl.x || w.pos.x > w.tr.x) {
		return true
	}
	return (w.pos.y < w.bl.y) || (w.pos.x > w.tr.x)
}

func parseInput(input string) (w world) {
	n, err := fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &w.bl.x, &w.tr.x, &w.bl.y, &w.tr.y)
	if err != nil {
		panic(err)
	}
	if n != 4 {
		panic("should be 4")
	}
	return w
}

var input = `target area: x=81..129, y=-150..-108`

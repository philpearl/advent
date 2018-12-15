package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"sort"
)

func main() {
	// part1("testmove.txt")
	// part1("testmove2.txt")
	// part1("test1.txt")
	// part1("test2.txt")
	// part1("test3.txt")
	// part1("test4.txt")
	// part1("test5.txt")
	// part1("test6.txt")
	// part1("test7.txt")
	part1("inputred.txt")
}

func part1(filename string) {
	w, err := readInput(filename)
	if err != nil {
		fmt.Println(err)
		return
	}

	rounds := 0
	w.print()
	for w.runRound() {
		fmt.Println()
		w.print()
		rounds++
	}

	fmt.Println()
	w.print()

	hp := w.totalHP()
	fmt.Printf("%d rounds %d points = %d\n", rounds, hp, rounds*hp)
}

func (w *world) totalHP() int {
	var total int
	for _, u := range w.units {
		if u.hitPoints > 0 {
			total += int(u.hitPoints)
		}
	}
	return total
}

func (w *world) runRound() (cont bool) {
	// Order units by book-order location
	w.sortUnits()

	for i := range w.units {
		u := w.units[i]
		if u.hitPoints <= 0 {
			continue
		}

		// Any enemies left?
		tt := u.targetType()
		finished := true
		for _, v := range w.units {
			if v.hitPoints > 0 && v.species == tt {
				finished = false
				break
			}
		}
		if finished {
			return false
		}

		// find shortest paths to a cell adjacent to an enemy, then get our next move
		// fmt.Printf("next step for %v\n", u.location)
		moveTo, found := w.findNextStep(u)

		if found {
			// Move
			// fmt.Printf("move %v to %v\n", u.location, moveTo)
			w.move(u, moveTo)
		} else {
			// fmt.Printf("%v does not move\n", u.location)
		}

		// Now try to attack
		w.attack(u)
	}

	return true
}

type cell struct {
	content byte
	unit    *unit
}

type location struct {
	x, y int
}

func distance(a, b location) int {
	x := a.x - b.x
	y := a.y - b.y
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}

type unit struct {
	location
	attackPower int
	hitPoints   int
	species     byte
}

type world struct {
	units []*unit

	arena []cell
	w     int
	h     int
}

func (w *world) print() {
	for y := 0; y < w.h; y++ {
		for x := 0; x < w.w; x++ {
			cell := w.arena[x+y*w.w]
			fmt.Printf("%c", cell.content)
		}
		fmt.Printf(" ")
		for x := 0; x < w.w; x++ {
			cell := w.arena[x+y*w.w]
			if cell.unit != nil {
				fmt.Printf(" %c(%d)", cell.unit.species, cell.unit.hitPoints)
			}
		}

		fmt.Printf("\n")
	}
}

func (w *world) sortUnits() {
	sort.Slice(w.units, func(i, j int) bool {
		a, b := w.units[i], w.units[j]
		return a.y < b.y || (a.y == b.y && a.x < b.x)
	})
}

func (w *world) attack(u *unit) {
	targetType := u.targetType()
	var target *unit
	for _, candidate := range w.candidates(u.location) {
		if w.check(candidate, targetType) == checkTarget {
			// Find the unit at this location
			cell := &w.arena[candidate.x+w.w*candidate.y]
			if target == nil ||
				cell.unit.hitPoints < target.hitPoints ||
				(cell.unit.hitPoints == target.hitPoints &&
					(cell.unit.y < target.y ||
						(cell.unit.y == target.y && cell.unit.x < target.x))) {
				target = cell.unit
			}
		}
	}
	if target == nil {
		return
	}

	target.hitPoints -= u.attackPower
	if target.hitPoints > 0 {
		return
	}

	cell := &w.arena[target.x+w.w*target.y]
	cell.content = '.'
	cell.unit = nil
}

func (w *world) move(from *unit, to location) {
	if distance(from.location, to) != 1 {
		panic(fmt.Sprintf("Illegal move for %v to %v\n", *from, to))
	}
	cell := &w.arena[from.x+w.w*from.y]
	cell.content = '.'
	cell.unit = nil

	from.location = to
	cell = &w.arena[from.x+w.w*from.y]
	if cell.content != '.' {
		panic(fmt.Sprintf("moving to %v, content is %c", to, cell.content))
	}
	cell.content = from.species
	cell.unit = from
}

func (u *unit) targetType() byte {
	if u.species == 'E' {
		return 'G'
	}
	return 'E'
}

func (w *world) findNextStep(from *unit) (location, bool) {
	var depth int
	var current, next, ends []location
	parents := make(map[location]location)
	targetType := from.targetType()

	current = append(current, from.location)
	parents[from.location] = from.location

	for len(ends) == 0 && len(current) > 0 {
		for _, currentLoc := range current {
			for _, candidate := range w.candidates(currentLoc) {
				if _, ok := parents[candidate]; ok {
					continue
				}
				switch w.check(candidate, targetType) {
				case checkInvalid:
					// can't go here
				case checkValid:
					// can go here
					parents[candidate] = currentLoc
					next = append(next, candidate)
				case checkTarget:
					// have reached a target. Need to evaluate the rest of this generation then stop
					// Note that the locations we're moving to is not this candidate, but the currentLoc that
					// is adjacent to a target
					parents[candidate] = currentLoc
					ends = append(ends, currentLoc)
				}
			}
		}

		current, next = next, current
		next = next[:0]
		depth++
	}

	// fmt.Printf("%d Paths found %v\n", len(ends), ends)

	// Calculate every shortest path to find candidate next steps
	firstStep := location{x: 127, y: 127}
	var found bool
	for _, e := range ends {
		if e == from.location {
			return e, false
		}
		// Calculate path
		for {
			n := parents[e]
			if n == from.location {
				if e.y < firstStep.y || (e.y == firstStep.y && e.x == firstStep.x) {
					firstStep = e
					found = true
				}
				break
			}
			e = n
		}

	}
	return firstStep, found
}

type checkResult int

const (
	checkInvalid checkResult = iota
	checkValid
	checkTarget
)

func (w *world) check(l location, targetType byte) checkResult {
	if l.x < 0 || l.y < 0 || l.x >= w.w || l.y >= w.h {
		return checkInvalid
	}
	cell := w.arena[l.x+l.y*w.w]
	if cell.content == '#' {
		return checkInvalid
	}
	if cell.content == '.' {
		return checkValid
	}
	if cell.content == targetType {
		if cell.unit == nil || cell.unit.species != targetType {
			panic(fmt.Sprintf("check cell unit doesn't match. location %v. target %c content %c - %c", l, targetType, cell.content, cell.unit.species))
		}
		return checkTarget
	}
	return checkInvalid
}

func (w *world) candidates(l location) []location {
	return []location{
		{x: l.x, y: l.y - 1},
		{x: l.x - 1, y: l.y},
		{x: l.x + 1, y: l.y},
		{x: l.x, y: l.y + 1},
	}
}

func readInput(filename string) (w world, err error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return w, err
	}

	lines := bytes.Split(data, []byte{'\n'})
	w.w, w.h = len(lines[0]), len(lines)

	w.arena = make([]cell, int(w.w)*int(w.h))

	for y, line := range lines {
		for x, c := range line {
			cell := &w.arena[x+y*int(w.w)]
			cell.content = c
			if c == 'G' || c == 'E' {
				w.units = append(w.units, &unit{
					location: location{
						x: x,
						y: y,
					},
					attackPower: 3,
					hitPoints:   200,
					species:     c,
				})
				cell.unit = w.units[len(w.units)-1]
			}
		}
	}

	return w, nil
}

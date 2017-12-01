package main

import (
	"container/list"
	"fmt"
	"strings"
)

type element int

const (
	STRONTIUM element = 1 << iota
	PLUTONIUM
	THULIUM
	RUTHENIUM
	CURIUM
	ELERIUM
	DILITHIUM
)

var elements = []element{STRONTIUM, PLUTONIUM, THULIUM, RUTHENIUM, CURIUM, ELERIUM, DILITHIUM}

func (e element) String() string {
	names := []string{}
	for _, elt := range elements {
		if e&elt == elt {
			switch elt {
			case STRONTIUM:
				names = append(names, "strontium")
			case PLUTONIUM:
				names = append(names, "plutonium")
			case THULIUM:
				names = append(names, "thulium")
			case RUTHENIUM:
				names = append(names, "ruthenium")
			case CURIUM:
				names = append(names, "curium")
			case ELERIUM:
				names = append(names, "elerium")
			case DILITHIUM:
				names = append(names, "dilithium")
			}
		}
	}
	return "[" + strings.Join(names, ", ") + "]"
}

type floor struct {
	micro element
	gen   element
}

func (f *floor) String() string {
	return fmt.Sprintf("micro: %s, gen: %s", f.micro, f.gen)
}

type building struct {
	floors   [4]floor
	elevator int
}

func (b *building) String() string {
	s := ""
	for i, f := range b.floors {
		if b.elevator == i {
			s += "E "
		} else {
			s += "  "
		}
		s += f.String() + "\n"
	}
	return s
}

type move struct {
	building   *building
	micro      element
	gen        element
	up         bool
	generation int
}

func (f *floor) valid() bool {
	// If any generators are present, then all present microchips must have
	// a matching generator
	if f.gen == 0 {
		return true
	}
	return f.micro&f.gen == f.micro
}

func applyMove(m move) (building, bool) {
	b := *m.building

	f := &(b.floors[b.elevator])
	f.gen &= ^m.gen
	f.micro &= ^m.micro
	if !f.valid() {
		return b, false
	}

	if m.up {
		b.elevator++
	} else {
		b.elevator--
	}
	if b.elevator > 3 || b.elevator < 0 {
		return b, false
	}

	f = &(b.floors[b.elevator])
	f.gen |= m.gen
	f.micro |= m.micro
	if !f.valid() {
		return b, false
	}

	return b, true
}

func (b *building) isDone() bool {
	for i := 0; i < 3; i++ {
		f := &b.floors[i]
		if f.gen != 0 || f.micro != 0 {
			return false
		}
	}
	return true
}

func (b *building) findMoves(moves *list.List, generation int) {
	f := &b.floors[b.elevator]

	// Possible moves are 1 or 2 available micro
	for i, elt := range elements {
		if elt&f.micro == elt {
			// 1 possible move is this elt alone
			moves.PushBack(move{micro: elt, building: b, generation: generation})
			moves.PushBack(move{micro: elt, building: b, up: true, generation: generation})

			// Look to see if there's a subsequent element we can have as well for
			// a second move. Only need to look forward to avoid duplicates
			for _, elt2 := range elements[i:] {
				if elt2&f.micro == elt2 {
					moves.PushBack(move{micro: elt | elt2, building: b, generation: generation})
					moves.PushBack(move{micro: elt | elt2, building: b, up: true, generation: generation})
				}
			}

			// A 3rd possible move this this micro and it's gen
			if elt&f.gen == elt {
				moves.PushBack(move{micro: elt, gen: elt, building: b, generation: generation})
				moves.PushBack(move{micro: elt, gen: elt, building: b, up: true, generation: generation})
			}
		}
	}

	// Or 1 or 2 available generators
	for i, elt := range elements {
		if elt&f.gen == elt {
			// 1 possible move is this elt alone
			moves.PushBack(move{gen: elt, building: b, generation: generation})
			moves.PushBack(move{gen: elt, building: b, up: true, generation: generation})
			for _, elt2 := range elements[i:] {
				if elt2&f.gen == elt2 {
					moves.PushBack(move{gen: elt | elt2, building: b, generation: generation})
					moves.PushBack(move{gen: elt | elt2, building: b, up: true, generation: generation})
				}
			}
		}
	}
}

func main() {
	b := building{
		floors: [4]floor{
			{gen: STRONTIUM | PLUTONIUM | ELERIUM | DILITHIUM, micro: STRONTIUM | PLUTONIUM | ELERIUM | DILITHIUM},
			{gen: THULIUM | RUTHENIUM | CURIUM, micro: RUTHENIUM | CURIUM},
			{micro: THULIUM},
			{},
		},
		elevator: 0,
	}

	fmt.Println(b.String())

	statesSeen := make(map[building]struct{})
	statesSeen[b] = struct{}{}

	moves := list.New()
	b.findMoves(moves, 1)
	for e := moves.Front(); e != nil; e = moves.Front() {
		moves.Remove(e)
		m := e.Value.(move)

		newB, valid := applyMove(m)
		if valid {
			// fmt.Println(newB.String())
			if newB.isDone() {
				fmt.Printf("Done by generation %d\n", m.generation)
				fmt.Println(newB.String())
				break
			}

			if _, seen := statesSeen[newB]; !seen {
				statesSeen[newB] = struct{}{}
				newB.findMoves(moves, m.generation+1)
			}

		}
	}
}

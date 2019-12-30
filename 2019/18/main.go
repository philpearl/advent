package main

import (
	"fmt"
	"math/bits"
	"strings"
)

func main() {
	run(input)

	run2(`#############
#g#f.D#..h#l#
#F###e#E###.#
#dCba@#@BcIJ#
#############
#nK.L@#@G...#
#M###N#H###.#
#o#m..#i#jk.#
#############`, 15)
	run2(input2, 26)
}

func run2(input string, numKeys int) {
	var depth int
	var w world
	w.arena = input
	w.width = strings.IndexRune(w.arena, '\n') + 1

	var s state2
	var index int
	for i := range s.robots {
		index = strings.IndexRune(w.arena[index:], '@') + index
		s.robots[i] = pos{x: int8(index % w.width), y: int8(index / w.width)}
		index++
	}

	fmt.Printf("robots at %v\n", s.robots)

	visited := make(map[state2]struct{})

	var q, nq []state2
	q = append(q, s)

	for len(q) != 0 {
		for _, s := range q {
			// Find all possible next steps for all robots
		ROBOTS:
			for ri, rpos := range s.robots {
				// if this robot is on a key or @ then we only move it if all
				// others are on a key or @. This cuts down the search space
				// without skipping anything that matters.
				//
				// This isn't terribly efficient to do this check on every loop
				if c := w.at(rpos); c == '@' || ('a' <= c && c <= 'z') {
					for _, p := range s.robots {
						if c := w.at(p); c != '@' && !('a' <= c && c <= 'z') {
							continue ROBOTS
						}
					}
				}
				for _, v := range []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
					keys := s.keys
					pos := rpos.add(v)
					c := w.at(pos)
					switch {
					case c == '#':
						continue
					case c >= 'A' && c <= 'Z':
						needed := uint32(1) << (c - 'A')
						if keys&needed == 0 {
							continue
						}
					case c >= 'a' && c <= 'z':
						keys |= uint32(1) << (c - 'a')
						// Are we done?
						if bits.OnesCount32(keys) >= numKeys {
							// Steps is one more than this as we haven't taken this
							// step yet
							fmt.Printf("all keys at depth %d\n", depth)
							return
						}
					}

					candidate := s
					candidate.robots[ri] = pos
					candidate.keys = keys
					if _, ok := visited[candidate]; ok {
						continue
					}
					visited[candidate] = struct{}{}

					nq = append(nq, candidate)
				}
			}
		}
		q, nq = nq, q
		nq = nq[:0]
		depth++
	}
}

func run(input string) {
	var depth int
	var w world
	w.arena = input
	w.width = strings.IndexRune(w.arena, '\n') + 1
	visited := make(map[state]struct{})

	currIndex := strings.IndexRune(w.arena, '@')

	var q, nq []state

	q = append(q, state{
		pos:  pos{x: int8(currIndex % w.width), y: int8(currIndex / w.width)},
		keys: 0,
	})

	for len(q) != 0 {
		for _, s := range q {
			// Find all possible next steps
			for _, v := range []pos{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
				candidate := state{
					keys: s.keys,
					pos:  s.pos.add(v),
				}
				c := w.at(candidate.pos)
				switch {
				case c == '#':
					continue
				case c >= 'A' && c <= 'Z':
					needed := uint32(1) << (c - 'A')
					if candidate.keys&needed == 0 {
						continue
					}
				case c >= 'a' && c <= 'z':
					candidate.keys |= uint32(1) << (c - 'a')
					// Are we done?
					if bits.OnesCount32(candidate.keys) >= 26 {
						// Steps is one more than this as we haven't taken this
						// step yet
						fmt.Printf("all keys at depth %d\n", depth)
						return
					}
				}

				if _, ok := visited[candidate]; ok {
					continue
				}
				visited[candidate] = struct{}{}
				nq = append(nq, candidate)
			}
		}
		q, nq = nq, q
		nq = nq[:0]
		depth++
	}
}

type pos struct{ x, y int8 }

func (p pos) add(p2 pos) pos {
	return pos{x: p2.x + p.x, y: p2.y + p.y}
}

type state struct {
	pos
	keys uint32
}

type state2 struct {
	robots [4]pos
	keys   uint32
}

type world struct {
	width int
	arena string
}

func (w *world) at(p pos) byte {
	return w.arena[int(p.y)*w.width+int(p.x)]
}

var input = `#################################################################################
#.........#.............#.....#.....R.#.#.....#...............#.#..v............#
#######.#T#.#######.###.#####.#.###.#.#.#.###.#.#.###.#######.#.#.#####.#######.#
#.....#.#.#.#.......#.#.F...#...#...#...#...#.#.#...#.#.....#.#...#...#...#.....#
#.###.#.#.#.#.#######.#####.###.#.#####.#.#.###.###.###.###.#.#.#####.###.###.###
#.#.#...#.#.#..a..#.......#...#.#.#...#.#.#.......#.......#.#.#.#.....#.#...#...#
#L#.#####.#.#####.#.###.#.###.###.#.###.#.#########.#######.#.#.###.#.#.###.###.#
#.#.......#...#...#...#.#...#...#.#.#...#.....#...#.#.....#.#.#.....#.....#...#.#
#.###########.#.#####.#.###.###.#.#.#.#####.###.#.###.###.#.#.#####.#########.#.#
#.............#.#...#.#.#...#...#.#.#...#.#.#...#.....#...#.#...#...#.........#.#
#.#########.###.###.#.#.#####.###.#.###.#.#.#Z#########.###.###.#.###.#########.#
#...#.....#...#...#...#.......#...#.....#...#...#...#.#...#...#.#...#...#.......#
###.#.###.#######.#.###########.#.#.###########.#.#.#.###.#.#.#.#######.#######.#
#.#.#.#...#.......#.....#.......#.#.....#.......#.#.M.#...#.#.#.#.....#.#.....#.#
#.#.#.#.###.###########.#.#######.#######.#######.#####.#####.#.#.###.#.#.###.###
#.#.#.#.....#...#...G.#.#.....#.#.......#...B.#.#.....#.#..s..#...#.#.#.#.#.#...#
#.#.#.#######.###.#.###.#####.#.#######.#.###.#.#.###.#.#.#########.#.#.#.#.###.#
#...#.#...#.....#.#.#...#.......#.....#.#.#i..#...#.#.....#.....#.....#.....#...#
#.###.#.#.#####.#.#.#.###########.###.#.###.#####.#.#######.###X#.###########.#.#
#.#.#...#.#..c..#.#.....#.......#h#.#...#...#...#.....#..x#...#.#.#.........#.#.#
#.#.#####.#.#####.#####.#.#####.#.#.###W#.###.#Y#######D#.#.###.#.#####.###.#Q###
#...#...#.#.#.....#b..#...#.#...#.#...#.#...#.#....n..#.#...#.#.#.#.....#...#...#
###.#.#.#.#.#.#####.#######.#.###.#.#.#.#.#.#.#######.#.#####.#.#.#.#####.#####.#
#.#.#.#.#.#.......#.....#...#.....#.#.#.#.#.#.#.#.....#...#...#.#.#.#.#...#.....#
#.#.#.#.#.#.###########.###.#######.#.#.###.#.#.#N#######.#E###.#.#.#.#.###.###.#
#...#.#...#.#.#.....#...#.........#.#.#.#pI.#...#......d..#....e#...#.#.....#.#.#
#####.#####.#.#.###.#.###.#########.#.#.#.#####.#####################.#######.#.#
#.....#...#.#.#...#...#.......#.....#.#.#.....#y#.....................#.........#
#.#####.###.#.###.#####.#####.#.#####.#.#####.#.#.###########.###.###.#.#########
#...#.......#...#...#.#.#.#...#.#.....#.#...#.#.#...#.#.....#.#...#...#.........#
#.#.#.#########.###.#.#.#.#.#.#.#.###.#.###.#.#.###.#.#.###.###.###.###########.#
#.#.#.#...........#.#...#.#.#.#.#...#.#.#...#.P.#.....#.#.......#.........#...#.#
#.#.#.#.#####.#####.#.###.#.###.###.#.#.#.###########.#.#.#####.#########.#.#.#C#
#.#.#...#.#...#...#.#.....#...#.#.#.#.#.#.#.........#.#.#.#...#.#.....#.#...#.#.#
#.#.#####.#.###.#.#.#####.###.#.#.#.###.#.#.###.###.###.###.#.###.###.#.#.#####.#
#.#.#.......#.#.#.#.#.....#.#.#.#.#...#.#.#...#...#.....#...#.#...#.#.#.#.#.....#
###.#.#######.#.#.#.#.#####.#.#.#.###.#.#.###.###.###.###.###.#.###.#.#.#.#.#####
#...#.........#.#...#.#.....#...#...#.#.#.....#.#.#...#...#...#...#.#.#..o#...#.#
#.#############.#####.#####.#######.#.#.#.#####.#.#####.###.#####.#.#.#######.#.#
#...........U...#...................#...........#.........#.........#...........#
#######################################.@.#######################################
#.....#.O.#.....#.......#...#.....#...........#.....#.............#...#...#..k..#
#####.#.#.#.#.#.###.###.#.###.#.###.###.#.###.###.#.###.#####.###.#.#.#.#.#.#.#.#
#...#...#...#.#...#...#.#.#...#.....#...#...#...#.#...#.#.....#...#.#.#.#...#.#.#
#.#.#.#######.###.###.#.#.#.#########.#.###.###.#.###.#.#.#.#####.#.#.#######.#.#
#.#.#...#.....#.#...#.#...#...#...#...#.#...#...#.#...#.#.#.#...#.#.#..u#.....#.#
#.#.#####.#####.###.#.#######.###.#.###.#.#####.#.#.###.#.###.#.###.###.#.#####.#
#.#.......#.........#.......#...#.#.#...#.....#...#...#.#.#...#.....#.#.#.#.....#
#.#########.###########.###.###.#.#.#.#######.#######H###.#.#########.#.#.#######
#.....#...#...#.......#...#.......#.#.#.#.....#.....#.#...#.#.........#.#.#.....#
#####.#.#.###.#####.#.###.#########.#.#.#.#####.#.###.#.#.#.###.###.#.#.#.#.###.#
#.....#.#...#....r..#.#.....#...#...#..w#.#.....#.......#.#...#.#...#.#.#.....#.#
#.#####.#############.#######.#.#.#####.#.###.###############.#.#.###.#.#######.#
#.#.....#.....#...............#...#.#...#...#.......#.......#.#.#...#.#...#.....#
#.###.###.#.###.###################.#.###.#.#########.#####.#.#####.#####.#.###.#
#..q#.....#.....#...#.......#...#.....#.#.#...........#.#...#.....#.....#.#.#.#.#
###.#.###########.#.#.#####.#.###.#####.#.#############.#.#.#####.#.#.###.#.#.#.#
#.#.#.......#.....#...#...#.#...#.......#.#...#...#.....#.#.....#z#.#...#.#.#.#.#
#.#.#######.#.#########.#.#.#.#.#######.#.#.#.#.#.#####.#.###.###.#####.#.#.#.#.#
#.#.....#...#.....#.#...#...#.#.......#.#.#.#...#.......#.#...#...#.....#.#.#...#
#.#####.#########.#.#.#########.#######.#.#.#######.#####.#.###.###.###.#.#.###.#
#.....#.#...#.....#.#.#.K.......#.......#.#.#.....#.#.....#.#.#.#...#...#.#...#.#
#.#.###.#.#.#.#####.#####.#.#####.#######.#.###.#.#.#.#######.#.#.###.###.#.#.#.#
#.#...#...#.#.#.....#.....#.#.....#.....#.#...#.#...#.#.....#...#...#...#.#.#.#.#
#.###.#####.#.#.###.#.###.###.#####.#.#.#.###.#####.#.#.###.#####.#.###.#.###.#.#
#...#.#...#...#...#...#...#...#...#.#.#.#...#.#...#.#.....#.....#.#...#.#.J...#.#
###.#.#.#.#######.#####.###.###.#.###.#.###.#.#.#.#########.###.#.###.#########.#
#.#.#...#.........#...#...#.#...#...#.#.#.#.#...#...#.....#.#...#...#.#.........#
#.#.#######.#.#######.###.#.#.#####.#.#.#.#.#.#####.#.###.#.#.###.###.#.#########
#.#.#.....#.#.#.....#...#.#.#.#...#...#.#...#.#...#.#.#.#...#.#...#...#.#.......#
#.#.#.#####V#.#.###.#.###.#.#.#.#.#####.#.#####.#.#.#.#.#####.#.###.###.#.###.#.#
#...#...#...#j#...#.#...#.#...#.#.....#.#...#...#...#.#.......#...#.....#...#.#.#
#.#####.#.#######.#.###.#.#####.###.###.###.#.#####.#.###.#######.###########.#.#
#.......#.......#.#...#.#.#.......#.#...#...#...#...#...#.#...#...#...#.....#.#l#
#######.#####.#.#.###.#.#.#.#####.#.#.###.###.#.#.#####.###.#.#.###.#.#.#.#.#.#.#
#.....#.....#.#.#...#.#.#.#...#...#.#...#...#.#.#.....#.#...#...#...#...#.#...#.#
#.###.#####.#.###.###.#.#.###.#.#######.###.#.#.#######.#.#######.#######.#####.#
#.#..m..#...#.....#...#.....#.#.......#f#...#.#...#.....#.#...........#...#..t#.#
#.#.#####.#########.#########.#######.#.#.#######.#.#####.#############.###.###.#
#.#.......S.......#.....A...........#...#g........#.....................#.......#
#################################################################################
`

var input2 = `#################################################################################
#.........#.............#.....#.....R.#.#.....#...............#.#..v............#
#######.#T#.#######.###.#####.#.###.#.#.#.###.#.#.###.#######.#.#.#####.#######.#
#.....#.#.#.#.......#.#.F...#...#...#...#...#.#.#...#.#.....#.#...#...#...#.....#
#.###.#.#.#.#.#######.#####.###.#.#####.#.#.###.###.###.###.#.#.#####.###.###.###
#.#.#...#.#.#..a..#.......#...#.#.#...#.#.#.......#.......#.#.#.#.....#.#...#...#
#L#.#####.#.#####.#.###.#.###.###.#.###.#.#########.#######.#.#.###.#.#.###.###.#
#.#.......#...#...#...#.#...#...#.#.#...#.....#...#.#.....#.#.#.....#.....#...#.#
#.###########.#.#####.#.###.###.#.#.#.#####.###.#.###.###.#.#.#####.#########.#.#
#.............#.#...#.#.#...#...#.#.#...#.#.#...#.....#...#.#...#...#.........#.#
#.#########.###.###.#.#.#####.###.#.###.#.#.#Z#########.###.###.#.###.#########.#
#...#.....#...#...#...#.......#...#.....#...#...#...#.#...#...#.#...#...#.......#
###.#.###.#######.#.###########.#.#.###########.#.#.#.###.#.#.#.#######.#######.#
#.#.#.#...#.......#.....#.......#.#.....#.......#.#.M.#...#.#.#.#.....#.#.....#.#
#.#.#.#.###.###########.#.#######.#######.#######.#####.#####.#.#.###.#.#.###.###
#.#.#.#.....#...#...G.#.#.....#.#.......#...B.#.#.....#.#..s..#...#.#.#.#.#.#...#
#.#.#.#######.###.#.###.#####.#.#######.#.###.#.#.###.#.#.#########.#.#.#.#.###.#
#...#.#...#.....#.#.#...#.......#.....#.#.#i..#...#.#.....#.....#.....#.....#...#
#.###.#.#.#####.#.#.#.###########.###.#.###.#####.#.#######.###X#.###########.#.#
#.#.#...#.#..c..#.#.....#.......#h#.#...#...#...#.....#..x#...#.#.#.........#.#.#
#.#.#####.#.#####.#####.#.#####.#.#.###W#.###.#Y#######D#.#.###.#.#####.###.#Q###
#...#...#.#.#.....#b..#...#.#...#.#...#.#...#.#....n..#.#...#.#.#.#.....#...#...#
###.#.#.#.#.#.#####.#######.#.###.#.#.#.#.#.#.#######.#.#####.#.#.#.#####.#####.#
#.#.#.#.#.#.......#.....#...#.....#.#.#.#.#.#.#.#.....#...#...#.#.#.#.#...#.....#
#.#.#.#.#.#.###########.###.#######.#.#.###.#.#.#N#######.#E###.#.#.#.#.###.###.#
#...#.#...#.#.#.....#...#.........#.#.#.#pI.#...#......d..#....e#...#.#.....#.#.#
#####.#####.#.#.###.#.###.#########.#.#.#.#####.#####################.#######.#.#
#.....#...#.#.#...#...#.......#.....#.#.#.....#y#.....................#.........#
#.#####.###.#.###.#####.#####.#.#####.#.#####.#.#.###########.###.###.#.#########
#...#.......#...#...#.#.#.#...#.#.....#.#...#.#.#...#.#.....#.#...#...#.........#
#.#.#.#########.###.#.#.#.#.#.#.#.###.#.###.#.#.###.#.#.###.###.###.###########.#
#.#.#.#...........#.#...#.#.#.#.#...#.#.#...#.P.#.....#.#.......#.........#...#.#
#.#.#.#.#####.#####.#.###.#.###.###.#.#.#.###########.#.#.#####.#########.#.#.#C#
#.#.#...#.#...#...#.#.....#...#.#.#.#.#.#.#.........#.#.#.#...#.#.....#.#...#.#.#
#.#.#####.#.###.#.#.#####.###.#.#.#.###.#.#.###.###.###.###.#.###.###.#.#.#####.#
#.#.#.......#.#.#.#.#.....#.#.#.#.#...#.#.#...#...#.....#...#.#...#.#.#.#.#.....#
###.#.#######.#.#.#.#.#####.#.#.#.###.#.#.###.###.###.###.###.#.###.#.#.#.#.#####
#...#.........#.#...#.#.....#...#...#.#.#.....#.#.#...#...#...#...#.#.#..o#...#.#
#.#############.#####.#####.#######.#.#.#.#####.#.#####.###.#####.#.#.#######.#.#
#...........U...#...................#..@#@......#.........#.........#...........#
#################################################################################
#.....#.O.#.....#.......#...#.....#....@#@....#.....#.............#...#...#..k..#
#####.#.#.#.#.#.###.###.#.###.#.###.###.#.###.###.#.###.#####.###.#.#.#.#.#.#.#.#
#...#...#...#.#...#...#.#.#...#.....#...#...#...#.#...#.#.....#...#.#.#.#...#.#.#
#.#.#.#######.###.###.#.#.#.#########.#.###.###.#.###.#.#.#.#####.#.#.#######.#.#
#.#.#...#.....#.#...#.#...#...#...#...#.#...#...#.#...#.#.#.#...#.#.#..u#.....#.#
#.#.#####.#####.###.#.#######.###.#.###.#.#####.#.#.###.#.###.#.###.###.#.#####.#
#.#.......#.........#.......#...#.#.#...#.....#...#...#.#.#...#.....#.#.#.#.....#
#.#########.###########.###.###.#.#.#.#######.#######H###.#.#########.#.#.#######
#.....#...#...#.......#...#.......#.#.#.#.....#.....#.#...#.#.........#.#.#.....#
#####.#.#.###.#####.#.###.#########.#.#.#.#####.#.###.#.#.#.###.###.#.#.#.#.###.#
#.....#.#...#....r..#.#.....#...#...#..w#.#.....#.......#.#...#.#...#.#.#.....#.#
#.#####.#############.#######.#.#.#####.#.###.###############.#.#.###.#.#######.#
#.#.....#.....#...............#...#.#...#...#.......#.......#.#.#...#.#...#.....#
#.###.###.#.###.###################.#.###.#.#########.#####.#.#####.#####.#.###.#
#..q#.....#.....#...#.......#...#.....#.#.#...........#.#...#.....#.....#.#.#.#.#
###.#.###########.#.#.#####.#.###.#####.#.#############.#.#.#####.#.#.###.#.#.#.#
#.#.#.......#.....#...#...#.#...#.......#.#...#...#.....#.#.....#z#.#...#.#.#.#.#
#.#.#######.#.#########.#.#.#.#.#######.#.#.#.#.#.#####.#.###.###.#####.#.#.#.#.#
#.#.....#...#.....#.#...#...#.#.......#.#.#.#...#.......#.#...#...#.....#.#.#...#
#.#####.#########.#.#.#########.#######.#.#.#######.#####.#.###.###.###.#.#.###.#
#.....#.#...#.....#.#.#.K.......#.......#.#.#.....#.#.....#.#.#.#...#...#.#...#.#
#.#.###.#.#.#.#####.#####.#.#####.#######.#.###.#.#.#.#######.#.#.###.###.#.#.#.#
#.#...#...#.#.#.....#.....#.#.....#.....#.#...#.#...#.#.....#...#...#...#.#.#.#.#
#.###.#####.#.#.###.#.###.###.#####.#.#.#.###.#####.#.#.###.#####.#.###.#.###.#.#
#...#.#...#...#...#...#...#...#...#.#.#.#...#.#...#.#.....#.....#.#...#.#.J...#.#
###.#.#.#.#######.#####.###.###.#.###.#.###.#.#.#.#########.###.#.###.#########.#
#.#.#...#.........#...#...#.#...#...#.#.#.#.#...#...#.....#.#...#...#.#.........#
#.#.#######.#.#######.###.#.#.#####.#.#.#.#.#.#####.#.###.#.#.###.###.#.#########
#.#.#.....#.#.#.....#...#.#.#.#...#...#.#...#.#...#.#.#.#...#.#...#...#.#.......#
#.#.#.#####V#.#.###.#.###.#.#.#.#.#####.#.#####.#.#.#.#.#####.#.###.###.#.###.#.#
#...#...#...#j#...#.#...#.#...#.#.....#.#...#...#...#.#.......#...#.....#...#.#.#
#.#####.#.#######.#.###.#.#####.###.###.###.#.#####.#.###.#######.###########.#.#
#.......#.......#.#...#.#.#.......#.#...#...#...#...#...#.#...#...#...#.....#.#l#
#######.#####.#.#.###.#.#.#.#####.#.#.###.###.#.#.#####.###.#.#.###.#.#.#.#.#.#.#
#.....#.....#.#.#...#.#.#.#...#...#.#...#...#.#.#.....#.#...#...#...#...#.#...#.#
#.###.#####.#.###.###.#.#.###.#.#######.###.#.#.#######.#.#######.#######.#####.#
#.#..m..#...#.....#...#.....#.#.......#f#...#.#...#.....#.#...........#...#..t#.#
#.#.#####.#########.#########.#######.#.#.#######.#.#####.#############.###.###.#
#.#.......S.......#.....A...........#...#g........#.....................#.......#
#################################################################################
`

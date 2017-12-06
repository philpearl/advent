package main

import (
	"fmt"
)

func main() {

	cycles := 0
	states := make(map[banks]int)
	memory := input
	for {
		states[memory] = cycles
		i := memory.mostBlocks()
		memory.redistribute(i)
		cycles++
		if firstSeen, seen := states[memory]; seen {
			fmt.Printf("Seen before after %d cycles, first seen %d, repeat %d\n", cycles, firstSeen, cycles-firstSeen)
			break
		}
	}
}

type banks [16]byte

func (b *banks) redistribute(i int) {
	blocks := (*b)[i]
	(*b)[i] = 0
	for blocks > 0 {
		i++
		if i >= len(*b) {
			i = 0
		}
		(*b)[i]++
		blocks--
	}
}

func (b *banks) mostBlocks() int {
	var maxC byte
	var maxI int
	for i, c := range *b {
		if c > maxC {
			maxI = i
			maxC = c
		}
	}
	return maxI
}

var input = banks{10, 3, 15, 10, 5, 15, 5, 15, 9, 2, 5, 8, 5, 2, 3, 6}

package main

import (
	"bytes"
	"fmt"
)

func main() {

	part1(initialState, stateTransitions, 20)
	part1(initialState, stateTransitions, 1000)
}

func part1(initial []byte, transitions []transition, rounds int) {

	arena := make([]byte, len(initial)+(rounds*4))
	for i := range arena {
		arena[i] = '.'
	}
	zero := rounds * 2
	copy(arena[zero:], initial)

	narena := make([]byte, len(arena))

	for r := 0; r < rounds; r++ {
		narena[0] = '.'
		narena[1] = '.'
		for i := range arena[:len(arena)-5] {
			for _, t := range transitions {
				if bytes.Compare(t.initial, arena[i:i+5]) == 0 {
					narena[i+2] = t.result
					break
				}
			}
		}
		arena, narena = narena, arena
		// fmt.Println(string(arena))
		if bytes.Compare(arena, narena) == 0 {
			fmt.Println("steady state")
			break
		}

		var sum int
		for i, v := range arena {
			if v == '#' {
				sum += i - zero
			}
		}
		fmt.Printf("round %d = %d\n", r+1, sum)

		fmt.Printf("50000000000 prediction %d\n", (50*1e9-(r+1))*32+sum)

	}

	// fmt.Println(string(arena))
}

var initialState = []byte("####....#...######.###.#...##....#.###.#.###.......###.##..##........##..#.#.#..##.##...####.#..##.#")

type transition struct {
	initial []byte
	result  byte
}

var stateTransitions = []transition{
	{initial: []byte("..#.."), result: '.'},
	{initial: []byte("#.#.#"), result: '#'},
	{initial: []byte("#.###"), result: '#'},
	{initial: []byte(".##.."), result: '.'},
	{initial: []byte("#.#.."), result: '#'},
	{initial: []byte(".#.#."), result: '#'},
	{initial: []byte(".###."), result: '#'},
	{initial: []byte(".####"), result: '#'},
	{initial: []byte("##..."), result: '#'},
	{initial: []byte("#.##."), result: '#'},
	{initial: []byte("#..##"), result: '#'},
	{initial: []byte("....#"), result: '.'},
	{initial: []byte("###.#"), result: '.'},
	{initial: []byte("#####"), result: '#'},
	{initial: []byte("....."), result: '.'},
	{initial: []byte("..#.#"), result: '.'},
	{initial: []byte(".#..."), result: '#'},
	{initial: []byte("##.#."), result: '.'},
	{initial: []byte(".#.##"), result: '#'},
	{initial: []byte("..##."), result: '.'},
	{initial: []byte("#...#"), result: '.'},
	{initial: []byte("##.##"), result: '#'},
	{initial: []byte("...#."), result: '.'},
	{initial: []byte("#..#."), result: '.'},
	{initial: []byte("..###"), result: '.'},
	{initial: []byte(".##.#"), result: '.'},
	{initial: []byte("#...."), result: '.'},
	{initial: []byte(".#..#"), result: '#'},
	{initial: []byte("####."), result: '.'},
	{initial: []byte("...##"), result: '#'},
	{initial: []byte("##..#"), result: '.'},
	{initial: []byte("###.."), result: '.'},
}

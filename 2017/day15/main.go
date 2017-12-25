package main

import "fmt"

func main() {

	fmt.Printf("count %d\n", part1(65, 8921, 40000000))
	fmt.Printf("count %d\n", part1(591, 393, 40000000))

	fmt.Printf("count %d\n", part2(65, 8921, 5000000))
	fmt.Printf("count %d\n", part2(591, 393, 5000000))

}

func part2(starta, startb int, rounds int) int {
	ga := generator{
		current:  starta,
		factor:   16807,
		criteria: 4,
	}
	gb := generator{
		current:  startb,
		factor:   48271,
		criteria: 8,
	}

	count := 0
	for i := 0; i < rounds; i++ {
		for ga.next() {
		}
		for gb.next() {
		}
		if ga.current&0xFFFF == gb.current&0xFFFF {
			count++
		}
	}
	return count
}

func part1(starta, startb int, rounds int) int {

	ga := generator{
		current:  starta,
		factor:   16807,
		criteria: 1,
	}
	gb := generator{
		current:  startb,
		factor:   48271,
		criteria: 1,
	}

	count := 0
	for i := 0; i < rounds; i++ {
		ga.next()
		gb.next()

		if ga.current&0xFFFF == gb.current&0xFFFF {
			count++
		}
	}
	return count
}

type generator struct {
	current  int
	factor   int
	criteria int
}

func (g *generator) next() bool {
	g.current = (g.current * g.factor) % 2147483647
	return (g.current % g.criteria) != 0
}

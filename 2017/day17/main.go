package main

import "fmt"

func main() {

	afterLast, afterZero := part1(343, 2017)
	fmt.Printf("next value is %d, %d\n", afterLast, afterZero)

	fmt.Printf("after zero %d\n", part2(343, 50000000))

	// afterLast, afterZero = part1(343, 50000000)
	// fmt.Printf("next value is %d, %d\n", afterLast, afterZero)
}

func part1(steps, max int) (node, node) {
	n := newNodes(max)

	for counter := node(1); counter <= node(max); counter++ {
		for j := 0; j < steps; j++ {
			n.fwd()
		}

		n.insert(counter)
	}

	return n.ns[n.current], n.ns[0]
}

func part2(steps, max int) int {
	len := 1
	current := 0
	afterZero := 0
	for i := 1; i <= max; i++ {
		current += steps
		current = current % len
		if current == 0 {
			afterZero = i
		}
		len++
		current++
	}
	return afterZero
}

type node int32

type nodes struct {
	current node
	ns      []node
}

func newNodes(max int) *nodes {
	n := &nodes{
		ns: make([]node, max+1),
	}
	for i := range n.ns {
		n.ns[i] = node(i)
	}

	return n
}

func (n *nodes) insert(v node) {
	n.ns[v] = n.ns[n.current]
	n.ns[n.current] = v
	n.current = v
}

func (n *nodes) fwd() {
	n.current = n.ns[n.current]
}

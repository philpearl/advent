package main

import (
	"fmt"
)

func main() {
	run([]byte("389125467"))
	run(input)

	run2([]byte("389125467"))
	run2(input)
}

/*
	A linked list where each number is represented by a slot in an array, and
    the value in the slot is the next number in the sequence

    2 5 8 6 4 7 3 9 1 next number
  0 1 2 3 4 5 6 7 8 9 index

*/
func run(input []byte) {
	remove := make([]int, 3)

	nextLabel := make([]int, len(input))

	first := int(input[0] - '1')
	last := first
	for _, l := range input[1:] {
		next := int(l - '1')
		nextLabel[last] = next
		last = next
	}
	nextLabel[last] = first

	cursor := first
	for i := 0; i < 100; i++ {
		end := cursor
		for i := range remove {
			end = nextLabel[end]
			remove[i] = end
		}

		dest := cursor - 1
	NEXT:
		if dest < 0 {
			dest = len(nextLabel) - 1
		}
		for _, v := range remove {
			if dest == v {
				dest--
				goto NEXT
			}
		}

		nextLabel[cursor] = nextLabel[remove[2]]

		nextLabel[dest], nextLabel[remove[2]] = remove[0], nextLabel[dest]
		cursor = nextLabel[cursor]
	}

	for i := nextLabel[0]; i > 0; i = nextLabel[i] {
		fmt.Printf("%c", i+'1')
	}

	fmt.Println("")
}

func run2(input []byte) {
	remove := make([]int, 3)

	nextLabel := make([]int, 1e6)

	first := int(input[0] - '1')
	last := first
	for _, l := range input[1:] {
		next := int(l - '1')
		nextLabel[last] = next
		last = next
	}
	nextLabel[last] = len(input)
	for i := len(input); i < len(nextLabel); i++ {
		nextLabel[i] = i + 1
	}
	nextLabel[1e6-1] = first

	cursor := first
	for i := 0; i < 1e7; i++ {
		end := cursor
		for i := range remove {
			end = nextLabel[end]
			remove[i] = end
		}

		dest := cursor - 1
	NEXT:
		if dest < 0 {
			dest = len(nextLabel) - 1
		}
		for _, v := range remove {
			if dest == v {
				dest--
				goto NEXT
			}
		}

		nextLabel[cursor] = nextLabel[remove[2]]

		nextLabel[dest], nextLabel[remove[2]] = remove[0], nextLabel[dest]
		cursor = nextLabel[cursor]
	}

	a := nextLabel[0]
	b := nextLabel[a]
	fmt.Println(a, b, (a+1)*(b+1))
}

var input = []byte(`215694783`)

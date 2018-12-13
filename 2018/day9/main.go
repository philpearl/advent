package main

import "fmt"

func main() {
	run(477, 70851)
	run(477, 7085100)
}

func run(numPlayers, lastMarble int) {

	players := make([]int, numPlayers)
	marbles := make(marbles, lastMarble+1)
	currentPlayer := -1

	currentMarble := &marbles[0]
	for marbleNum := 1; marbleNum <= lastMarble; marbleNum++ {
		currentPlayer++
		if currentPlayer == numPlayers {
			currentPlayer = 0
		}

		// printMarbles(currentMarble)

		if marbleNum%23 == 0 {
			players[currentPlayer] += int(marbleNum)
			for i := 0; i < 6; i++ {
				currentMarble = &marbles[currentMarble.prev]
			}
			players[currentPlayer] += int(currentMarble.prev)
			marbles.remove(currentMarble.prev)
			continue
		}

		marbles.insertAfter(currentMarble.next, int32(marbleNum))
		currentMarble = &marbles[marbleNum]
	}

	var highScore int
	for _, score := range players {
		if score > highScore {
			highScore = score
		}
	}

	fmt.Println(highScore)
}

func (mm marbles) printMarbles(m int32) {
	fmt.Printf("%d ", m)
	for c := mm[m].next; c != m; c = mm[c].next {
		fmt.Printf("%d ", c)
	}
	fmt.Println()
}

type marble struct {
	next, prev int32
}

type marbles []marble

func (mm marbles) insertAfter(current, newm int32) {
	cm := &mm[current]
	nm := &mm[newm]
	nm.next = cm.next
	nm.prev = current
	mm[cm.next].prev = newm
	cm.next = newm
}

func (mm marbles) remove(mv int32) {
	m := &mm[mv]
	mm[m.prev].next = m.next
	mm[m.next].prev = m.prev
}

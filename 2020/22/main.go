package main

import (
	"fmt"
	"log"
	"net/http"

	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	run(player1, player2)
	run2([]byte{43, 19}, []byte{2, 29, 14})
	run2([]byte{9, 2, 6, 3, 1}, []byte{5, 8, 4, 7, 10})
	run2(player1, player2)
}

func run(player1, player2 []byte) {
	for len(player1) > 0 && len(player2) > 0 {
		c1, c2 := player1[0], player2[0]
		player1, player2 = player1[1:], player2[1:]

		if c1 > c2 {
			player1 = append(player1, c1, c2)
		}
		if c2 > c1 {
			player2 = append(player2, c2, c1)
		}
	}

	var score int
	for i, c := range player1 {
		score += int(c) * (len(player1) - i)
	}
	for i, c := range player2 {
		score += int(c) * (len(player2) - i)
	}
	fmt.Println(score)
}

func run2(player1, player2 []byte) {
	p1, s := play(player1, player2)
	fmt.Println(p1, s)
}

func stateToKey(player1, player2 []byte) (key [64]byte) {
	n := copy(key[:], player1)
	key[n] = 0
	copy(key[n+1:], player2)
	return key
}

type result struct {
	player1Wins bool
	score       int
}

func play(player1, player2 []byte) (player1Wins bool, score int) {
	var p1, p2 []byte
	seen := make(map[[64]byte]struct{})
	for len(player1) > 0 && len(player2) > 0 {

		// first check for repeated state. Player 1 wins in this case. Just repeated within this function
		key := stateToKey(player1, player2)
		if _, ok := seen[key]; ok {
			player2 = nil
			break
		}
		seen[key] = struct{}{}

		c1, c2 := player1[0], player2[0]
		copy(player1, player1[1:])
		player1 = player1[:len(player1)-1]
		copy(player2, player2[1:])
		player2 = player2[:len(player2)-1]

		if int(c1) <= len(player1) && int(c2) <= len(player2) {
			// play a sub-game
			p1 = append(p1[:0], player1[:c1]...)
			p2 = append(p2[:0], player2[:c2]...)
			p1Wins, _ := play(p1, p2)
			if p1Wins {
				player1 = append(player1, c1, c2)
			} else {
				player2 = append(player2, c2, c1)
			}
		} else {
			if c1 > c2 {
				player1 = append(player1, c1, c2)
			}
			if c2 > c1 {
				player2 = append(player2, c2, c1)
			}
		}
	}

	for i, c := range player1 {
		score += int(c) * (len(player1) - i)
	}
	for i, c := range player2 {
		score += int(c) * (len(player2) - i)
	}

	return len(player1) > 0, score
}

var player1 = []byte{
	21,
	48,
	44,
	31,
	29,
	5,
	23,
	11,
	12,
	27,
	49,
	22,
	18,
	7,
	15,
	20,
	2,
	45,
	14,
	17,
	40,
	35,
	6,
	24,
	41,
}

var player2 = []byte{
	47,
	1,
	10,
	16,
	28,
	37,
	8,
	26,
	46,
	25,
	3,
	9,
	34,
	50,
	32,
	36,
	43,
	4,
	42,
	33,
	19,
	13,
	38,
	39,
	30,
}

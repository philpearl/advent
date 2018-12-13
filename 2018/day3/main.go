package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	claims, err := readInput("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	part1(claims)
	part2(claims)
}

func part2(claims []claim) {
	maxX, maxY := findArea(claims)
	area := make([]int16, maxX*maxY)

	overlaps := make([]bool, len(claims))

	for claimNo, claim := range claims {
		for i := 0; i < claim.w; i++ {
			for j := 0; j < claim.h; j++ {
				spot := (claim.x+i)*maxY + (claim.y + j)

				current := area[spot]
				if current != 0 {
					overlaps[current-1] = true
					overlaps[claimNo] = true
				} else {
					area[spot] = int16(claimNo + 1)
				}
			}
		}
	}

	for i, overlapping := range overlaps {
		if !overlapping {
			fmt.Println(i + 1)
		}
	}
}

func findArea(claims []claim) (x, y int) {
	var maxX, maxY int
	for _, claim := range claims {
		x := claim.x + claim.w
		y := claim.y + claim.h
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
	}

	return maxX, maxY
}

func part1(claims []claim) {

	maxX, maxY := findArea(claims)
	area := make([]byte, maxX*maxY)

	for _, claim := range claims {
		for i := 0; i < claim.w; i++ {
			for j := 0; j < claim.h; j++ {
				area[(claim.x+i)*maxY+(claim.y+j)]++
			}
		}
	}
	var overlaps int
	for _, c := range area {
		if c > 1 {
			overlaps++
		}
	}

	fmt.Println(overlaps)
}

type claim struct {
	x, y int
	w, h int
}

func readInput(filename string) ([]claim, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)
	var claims []claim
	for s.Scan() {
		// #1 @ 56,249: 24x16
		line := s.Bytes()
		parts := bytes.Split(line, []byte(" @ "))
		// don't seem to need the line number
		parts = bytes.Split(parts[1], []byte(": "))
		// parts[0] is x,y
		xy := bytes.Split(parts[0], []byte{','})

		x, _ := strconv.Atoi(string(xy[0]))
		y, _ := strconv.Atoi(string(xy[1]))
		wh := bytes.Split(parts[1], []byte{'x'})
		w, _ := strconv.Atoi(string(wh[0]))
		h, _ := strconv.Atoi(string(wh[1]))

		claims = append(claims, claim{x: x, y: y, w: w, h: h})
	}
	return claims, s.Err()
}

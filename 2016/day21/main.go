package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

/*
rotate based on position of letter d
move position 1 to position 6
swap position 3 with position 6
rotate right 5 steps
rotate left 3 steps
reverse positions 0 through 3
swap letter c with letter f
*/

func atoi(val string) int {
	i, _ := strconv.Atoi(val)
	return i
}

type problem struct {
	data []byte
}

func (p *problem) parseLine(line string) {
	fields := strings.Fields(line)

	switch fields[0] {
	case "move":
		from := atoi(fields[2])
		to := atoi(fields[5])
		p.move(to, from)
	case "reverse":
		start := atoi(fields[2])
		end := atoi(fields[4])
		p.reverse(start, end)
	case "rotate":
		switch fields[1] {
		case "based":
			letter := fields[6]
			p.rotateBased(letter[0])
		case "right":
			amount := atoi(fields[2])
			p.rotate(-amount)
		case "left":
			amount := atoi(fields[2])
			p.rotate(amount)
		}
	case "swap":
		switch fields[1] {
		case "position":
			first := atoi(fields[2])
			second := atoi(fields[5])
			p.data[first], p.data[second] = p.data[second], p.data[first]

		case "letter":
			first := fields[2][0]
			second := fields[5][0]
			for i, v := range p.data {
				if v == first {
					p.data[i] = second
				} else if v == second {
					p.data[i] = first
				}
			}
		}
	}
}

func (p *problem) move(from, to int) {
	// take value out of position from and put in postion to
	val := p.data[from]

	if from < to {
		// shuffle data down, leave a space to put val in at to
		for i := from; i < to; i++ {
			p.data[i] = p.data[i+1]
		}
	} else {
		// from > to
		// shuffle data up, leave space to put val in at to
		for i := from; i > to; i-- {
			p.data[i] = p.data[i-1]
		}
	}

	p.data[to] = val
}

func (p *problem) reverse(start, end int) {
	data := make([]byte, end-start+1)
	copy(data, p.data[start:end+1])

	for i, c := range data {
		p.data[end-i] = c
	}
}

func (p *problem) rotateBased(letter byte) {
	// This is the tricky one to reverse!
	//
	// abbbbbbb
	// newIndex = 2*oldIndex+1
	// oldIndex = (newIndex-1) / 2
	//
	// Forward map moves positions as follows
	// 0 -> 1
	// 1 -> 3
	// 2 -> 5
	// 3 -> 7
	// 4 -> 10 = 2
	// 5 -> 12 = 4
	// 6 -> 14 = 6
	// 7 -> 16 = 0

	rotationMap := map[int]int{
		1: -1,
		3: -2,
		5: -3,
		7: -4,
		2: 2,
		4: 1,
		6: 0,
		0: -1,
	}
	index := bytes.IndexByte(p.data, letter)
	p.rotate(rotationMap[index])
}

func (p *problem) rotate(count int) {
	for count > 0 {
		p.rotateOnce()
		count--
	}
	for count < 0 {
		p.rotateOnceBack()
		count++
	}
}

func (p *problem) rotateOnce() {
	last := p.data[len(p.data)-1]
	for i, c := range p.data {
		p.data[i] = last
		last = c
	}
}

func (p *problem) rotateOnceBack() {
	last := p.data[0]
	ll := len(p.data)
	for i := ll - 1; i >= 0; i-- {
		p.data[i], last = last, p.data[i]
	}
}

func main() {

	lines := strings.Split(input, "\n")

	p := problem{data: []byte("fbgdceah")}

	ll := len(lines)
	for i := ll - 1; i >= 0; i-- {
		line := lines[i]
		p.parseLine(line)

		fmt.Printf("%s\t%s\n", line, string(p.data))
	}

	fmt.Printf("result is %s\n", string(p.data))
}

var input = `rotate based on position of letter d
move position 1 to position 6
swap position 3 with position 6
rotate based on position of letter c
swap position 0 with position 1
rotate right 5 steps
rotate left 3 steps
rotate based on position of letter b
swap position 0 with position 2
rotate based on position of letter g
rotate left 0 steps
reverse positions 0 through 3
rotate based on position of letter a
rotate based on position of letter h
rotate based on position of letter a
rotate based on position of letter g
rotate left 5 steps
move position 3 to position 7
rotate right 5 steps
rotate based on position of letter f
rotate right 7 steps
rotate based on position of letter a
rotate right 6 steps
rotate based on position of letter a
swap letter c with letter f
reverse positions 2 through 6
rotate left 1 step
reverse positions 3 through 5
rotate based on position of letter f
swap position 6 with position 5
swap letter h with letter e
move position 1 to position 3
swap letter c with letter h
reverse positions 4 through 7
swap letter f with letter h
rotate based on position of letter f
rotate based on position of letter g
reverse positions 3 through 4
rotate left 7 steps
swap letter h with letter a
rotate based on position of letter e
rotate based on position of letter f
rotate based on position of letter g
move position 5 to position 0
rotate based on position of letter c
reverse positions 3 through 6
rotate right 4 steps
move position 1 to position 2
reverse positions 3 through 6
swap letter g with letter a
rotate based on position of letter d
rotate based on position of letter a
swap position 0 with position 7
rotate left 7 steps
rotate right 2 steps
rotate right 6 steps
rotate based on position of letter b
rotate right 2 steps
swap position 7 with position 4
rotate left 4 steps
rotate left 3 steps
swap position 2 with position 7
move position 5 to position 4
rotate right 3 steps
rotate based on position of letter g
move position 1 to position 2
swap position 7 with position 0
move position 4 to position 6
move position 3 to position 0
rotate based on position of letter f
swap letter g with letter d
swap position 1 with position 5
reverse positions 0 through 2
swap position 7 with position 3
rotate based on position of letter g
swap letter c with letter a
rotate based on position of letter g
reverse positions 3 through 5
move position 6 to position 3
swap letter b with letter e
reverse positions 5 through 6
move position 6 to position 7
swap letter a with letter e
swap position 6 with position 2
move position 4 to position 5
rotate left 5 steps
swap letter a with letter d
swap letter e with letter g
swap position 3 with position 7
reverse positions 0 through 5
swap position 5 with position 7
swap position 1 with position 7
swap position 1 with position 7
rotate right 7 steps
swap letter f with letter a
reverse positions 0 through 7
rotate based on position of letter d
reverse positions 2 through 4
swap position 7 with position 1
swap letter a with letter h`

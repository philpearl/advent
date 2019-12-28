package main

import (
	"fmt"
)

func main() {
	run(input)
}

func run(input []int) {
	// Positions that aren't in the map are black
	arena := make(map[pos]colour)
	var count int
	var p pos
	d := pos{0, -1}

	var cmp amp
	cmp.input = 1 // start on white for part 2
	cmp.program = cp(input)
	cmp.output = func(o int) {
		if count&1 == 0 {
			// Paint with this colour
			arena[p] = colour(o)

			// Input always matches colour of this square
			cmp.input = o
		} else {
			// Turn then move
			if o == 0 {
				// left 90
				d.x, d.y = d.y, -d.x
			} else {
				// right 90
				d.x, d.y = -d.y, d.x
			}
			// move forward
			p.x += d.x
			p.y += d.y

			// Set the new input colour to match the square we're on
			c, _ := arena[p]
			cmp.input = int(c)
		}
		count++
	}

	cmp.run()

	var maxX, maxY int
	for p := range arena {
		if p.x > maxX {
			maxX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
	}
	fmt.Println(len(arena))

	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			c, _ := arena[pos{x, y}]
			if c == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

type pos struct{ x, y int }

type colour int8

const (
	black colour = 0
	white colour = 1
)

type amp struct {
	program []int
	input   int
	output  func(o int)
}

func cp(program []int) []int {
	n := make([]int, len(program)*10)
	copy(n, program)
	return n
}

func (a *amp) run() int {
	var base int
	pv := func(pc, num, opcode int) *int {
		v := &a.program[pc+num]
		for i := 0; i <= num; i++ {
			opcode /= 10
		}
		if opcode%10 == 2 {
			v = &a.program[*v+base]
		} else if opcode%10 == 0 {
			v = &a.program[*v]
		}
		return v
	}
	for i := 0; i < len(a.program); {
		opcode := a.program[i]
		switch opcode % 100 {
		case 1:
			*pv(i, 3, opcode) = *pv(i, 1, opcode) + *pv(i, 2, opcode)
			i += 4
		case 2:
			*pv(i, 3, opcode) = *pv(i, 1, opcode) * *pv(i, 2, opcode)
			i += 4
		case 3:
			// Save input here
			*pv(i, 1, opcode) = a.input
			i += 2
		case 4:
			// Send this to output
			a.output(*pv(i, 1, opcode))
			i += 2
		case 5:
			if *pv(i, 1, opcode) != 0 {
				i = *pv(i, 2, opcode)
			} else {
				i += 3
			}
		case 6:
			if *pv(i, 1, opcode) == 0 {
				i = *pv(i, 2, opcode)
			} else {
				i += 3
			}
		case 7:
			if *pv(i, 1, opcode) < *pv(i, 2, opcode) {
				*pv(i, 3, opcode) = 1
			} else {
				*pv(i, 3, opcode) = 0
			}
			i += 4
		case 8:
			if *pv(i, 1, opcode) == *pv(i, 2, opcode) {
				*pv(i, 3, opcode) = 1
			} else {
				*pv(i, 3, opcode) = 0
			}
			i += 4
		case 9:
			base += *pv(i, 1, opcode)
			i += 2

		case 99:
			return a.program[0]
		default:
			panic(fmt.Sprintf("bad opcode %d", opcode))
		}
	}
	panic("off the end")
}

var input = []int{
	3, 8, 1005, 8, 318, 1106, 0, 11, 0, 0, 0, 104, 1, 104, 0, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 101, 0, 8, 29, 1, 107, 12, 10, 2, 1003, 8, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1002, 8, 1, 59, 1, 108, 18, 10, 2, 6, 7, 10, 2, 1006, 3, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1002, 8, 1, 93, 1, 1102, 11, 10, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 101, 0, 8, 118, 2, 1102, 10, 10, 3, 8, 102, -1, 8, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 145, 1006, 0, 17, 1006, 0, 67, 3, 8, 1002, 8, -1, 10, 101, 1, 10, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 101, 0, 8, 173, 2, 1109, 4, 10, 1006, 0, 20, 3, 8, 102, -1, 8, 10, 1001, 10, 1, 10, 4, 10, 108, 0, 8, 10, 4, 10, 102, 1, 8, 201, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 0, 10, 4, 10, 1002, 8, 1, 224, 1006, 0, 6, 1, 1008, 17, 10, 2, 101, 5, 10, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 108, 1, 8, 10, 4, 10, 1001, 8, 0, 256, 2, 1107, 7, 10, 1, 2, 4, 10, 2, 2, 12, 10, 1006, 0, 82, 3, 8, 1002, 8, -1, 10, 1001, 10, 1, 10, 4, 10, 1008, 8, 1, 10, 4, 10, 1002, 8, 1, 294, 2, 1107, 2, 10, 101, 1, 9, 9, 1007, 9, 988, 10, 1005, 10, 15, 99, 109, 640, 104, 0, 104, 1, 21102, 1, 837548352256, 1, 21102, 335, 1, 0, 1105, 1, 439, 21102, 1, 47677543180, 1, 21102, 346, 1, 0, 1106, 0, 439, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 1, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 1, 21102, 1, 235190374592, 1, 21101, 393, 0, 0, 1105, 1, 439, 21102, 3451060455, 1, 1, 21102, 404, 1, 0, 1105, 1, 439, 3, 10, 104, 0, 104, 0, 3, 10, 104, 0, 104, 0, 21102, 837896909668, 1, 1, 21102, 1, 427, 0, 1105, 1, 439, 21102, 1, 709580555020, 1, 21102, 438, 1, 0, 1105, 1, 439, 99, 109, 2, 21201, -1, 0, 1, 21102, 1, 40, 2, 21102, 1, 470, 3, 21102, 460, 1, 0, 1106, 0, 503, 109, -2, 2105, 1, 0, 0, 1, 0, 0, 1, 109, 2, 3, 10, 204, -1, 1001, 465, 466, 481, 4, 0, 1001, 465, 1, 465, 108, 4, 465, 10, 1006, 10, 497, 1101, 0, 0, 465, 109, -2, 2105, 1, 0, 0, 109, 4, 1201, -1, 0, 502, 1207, -3, 0, 10, 1006, 10, 520, 21101, 0, 0, -3, 21202, -3, 1, 1, 22101, 0, -2, 2, 21101, 1, 0, 3, 21101, 0, 539, 0, 1106, 0, 544, 109, -4, 2105, 1, 0, 109, 5, 1207, -3, 1, 10, 1006, 10, 567, 2207, -4, -2, 10, 1006, 10, 567, 21202, -4, 1, -4, 1105, 1, 635, 22101, 0, -4, 1, 21201, -3, -1, 2, 21202, -2, 2, 3, 21101, 0, 586, 0, 1105, 1, 544, 22102, 1, 1, -4, 21102, 1, 1, -1, 2207, -4, -2, 10, 1006, 10, 605, 21102, 1, 0, -1, 22202, -2, -1, -2, 2107, 0, -3, 10, 1006, 10, 627, 21202, -1, 1, 1, 21101, 627, 0, 0, 105, 1, 502, 21202, -2, -1, -2, 22201, -4, -2, -4, 109, -5, 2105, 1, 0,
}
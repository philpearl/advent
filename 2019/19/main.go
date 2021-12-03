package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	run1(input)
}

type pos struct{ x, y int }

func (p pos) add(p2 pos) pos {
	return pos{x: p2.x + p.x, y: p2.y + p.y}
}

func (p pos) left() pos {
	// 1, 0 => 0, -1
	return pos{x: p.y, y: -p.x}
}

func (p pos) right() pos {
	return pos{x: -p.y, y: p.x}
}

func run1(input []int) {

	width, height := 100, 100

	var comp amp
	var x, y, count int
	comp.program = cp(input)
	comp.output = func(o int) {
		switch o {
		case 0:
			io.WriteString(os.Stdout, " ")
		case 1:
			io.WriteString(os.Stdout, "#")
			count++
		default:
			fmt.Printf("unexpected output %d\n", o)
		}

		x++
		if x == width {
			io.WriteString(os.Stdout, "\n")
			x = 0
			y++
		}
		comp.inputs = append(comp.inputs, x, y)
	}

	comp.inputs = append(comp.inputs, x, y)
	for i := 0; i < width*height; i++ {
		comp.program = cp(input)
		comp.run()
	}

	fmt.Println(count)
}

type amp struct {
	program []int
	inputs  []int
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
			*pv(i, 1, opcode) = int(a.inputs[0])
			a.inputs = a.inputs[1:]
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
	109, 424, 203, 1, 21101, 11, 0, 0, 1106, 0, 282, 21102, 18, 1, 0, 1106, 0, 259, 1201, 1, 0, 221, 203, 1, 21102, 31, 1, 0, 1106, 0, 282, 21101, 0, 38, 0, 1105, 1, 259, 21002, 23, 1, 2, 21202, 1, 1, 3, 21102, 1, 1, 1, 21102, 1, 57, 0, 1105, 1, 303, 2101, 0, 1, 222, 21002, 221, 1, 3, 21002, 221, 1, 2, 21101, 0, 259, 1, 21101, 0, 80, 0, 1105, 1, 225, 21101, 169, 0, 2, 21101, 0, 91, 0, 1106, 0, 303, 1202, 1, 1, 223, 20101, 0, 222, 4, 21101, 259, 0, 3, 21102, 225, 1, 2, 21102, 225, 1, 1, 21101, 0, 118, 0, 1106, 0, 225, 20102, 1, 222, 3, 21101, 94, 0, 2, 21101, 0, 133, 0, 1106, 0, 303, 21202, 1, -1, 1, 22001, 223, 1, 1, 21102, 148, 1, 0, 1105, 1, 259, 2102, 1, 1, 223, 21001, 221, 0, 4, 21002, 222, 1, 3, 21101, 0, 22, 2, 1001, 132, -2, 224, 1002, 224, 2, 224, 1001, 224, 3, 224, 1002, 132, -1, 132, 1, 224, 132, 224, 21001, 224, 1, 1, 21101, 0, 195, 0, 106, 0, 108, 20207, 1, 223, 2, 21002, 23, 1, 1, 21102, 1, -1, 3, 21102, 214, 1, 0, 1105, 1, 303, 22101, 1, 1, 1, 204, 1, 99, 0, 0, 0, 0, 109, 5, 1202, -4, 1, 249, 21201, -3, 0, 1, 21202, -2, 1, 2, 22101, 0, -1, 3, 21101, 0, 250, 0, 1106, 0, 225, 21202, 1, 1, -4, 109, -5, 2106, 0, 0, 109, 3, 22107, 0, -2, -1, 21202, -1, 2, -1, 21201, -1, -1, -1, 22202, -1, -2, -2, 109, -3, 2105, 1, 0, 109, 3, 21207, -2, 0, -1, 1206, -1, 294, 104, 0, 99, 21202, -2, 1, -2, 109, -3, 2105, 1, 0, 109, 5, 22207, -3, -4, -1, 1206, -1, 346, 22201, -4, -3, -4, 21202, -3, -1, -1, 22201, -4, -1, 2, 21202, 2, -1, -1, 22201, -4, -1, 1, 22101, 0, -2, 3, 21102, 343, 1, 0, 1105, 1, 303, 1106, 0, 415, 22207, -2, -3, -1, 1206, -1, 387, 22201, -3, -2, -3, 21202, -2, -1, -1, 22201, -3, -1, 3, 21202, 3, -1, -1, 22201, -3, -1, 2, 21201, -4, 0, 1, 21101, 0, 384, 0, 1105, 1, 303, 1106, 0, 415, 21202, -4, -1, -4, 22201, -4, -3, -4, 22202, -3, -2, -2, 22202, -2, -4, -4, 22202, -3, -2, -3, 21202, -4, -1, -2, 22201, -3, -2, 1, 21201, 1, 0, -4, 109, -5, 2106, 0, 0,
}

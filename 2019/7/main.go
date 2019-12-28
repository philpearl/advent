package main

import (
	"fmt"
	"sync"
)

func main() {
	run1(input)
	run2(input)
}

func run2(input []int) {

	var max int
	phases := []int{5, 6, 7, 8, 9}
	perm(phases, func(perm []int) {
		var amps [5]amp
		var in chan int
		for i := range amps {
			amps[i].program = cp(input)
			amps[i].inputs = in
			amps[i].outputs = make(chan int, 2)
			in = amps[i].outputs
		}
		amps[0].inputs = amps[4].outputs

		for i, v := range perm {
			amps[i].inputs <- v
		}

		amps[0].inputs <- 0

		var wg sync.WaitGroup
		for i := range amps {
			i := i
			wg.Add(1)
			go func() {
				defer wg.Done()
				amps[i].run()
			}()
		}

		wg.Wait()

		val := <-amps[4].outputs
		if val > max {
			max = val
		}
	}, 0)
	fmt.Println(max)
}

func run1(input []int) {

	var max int
	phases := []int{0, 1, 2, 3, 4}
	perm(phases, func(perm []int) {
		var amps [5]amp
		in := make(chan int, 2)

		for i := range amps {
			amps[i].program = cp(input)
			amps[i].inputs = in
			amps[i].outputs = make(chan int, 2)
			in = amps[i].outputs
		}

		for i, v := range perm {
			amps[i].inputs <- v
		}

		amps[0].inputs <- 0
		for i := range amps {
			amps[i].run()
		}
		val := <-amps[4].outputs
		if val > max {
			max = val
		}
	}, 0)
	fmt.Println(max)
}

// Permute the values at index i to len(a)-1.
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

type amp struct {
	program []int
	inputs  chan int
	outputs chan int
}

func cp(program []int) []int {
	n := make([]int, len(program))
	copy(n, program)
	return n
}

func (a amp) run() int {
	for i := 0; i < len(a.program); {
		opcode := a.program[i]
		switch opcode % 100 {
		case 1:
			p0 := a.program[i+1]
			if (opcode/100)%10 == 0 {
				p0 = a.program[p0]
			}
			p1 := a.program[i+2]
			if (opcode/1000)%10 == 0 {
				p1 = a.program[p1]
			}
			if (opcode/10000)%10 == 0 {
				a.program[a.program[i+3]] = p0 + p1
			} else {
				a.program[i+3] = p0 + p1
			}
			i += 4
		case 2:
			p0 := a.program[i+1]
			if (opcode/100)%10 == 0 {
				p0 = a.program[p0]
			}
			p1 := a.program[i+2]
			if (opcode/1000)%10 == 0 {
				p1 = a.program[p1]
			}
			if (opcode/10000)%10 == 0 {
				a.program[a.program[i+3]] = p0 * p1
			} else {
				a.program[i+3] = p0 * p1
			}
			i += 4
		case 3:
			p0 := &a.program[i+1]
			if (opcode/100)%10 == 0 {
				p0 = &a.program[*p0]
			}
			// Save input here
			*p0 = <-a.inputs
			i += 2
		case 4:
			p0 := a.program[i+1]
			if (opcode/100)%10 == 0 {
				p0 = a.program[p0]
			}
			// Send this to output
			a.outputs <- p0
			i += 2
		case 5:
			p0 := a.program[i+1]
			if (opcode/100)%10 == 0 {
				p0 = a.program[p0]
			}
			p1 := a.program[i+2]
			if (opcode/1000)%10 == 0 {
				p1 = a.program[p1]
			}
			if p0 != 0 {
				i = p1
			} else {
				i += 3
			}
		case 6:
			p0 := a.program[i+1]
			if (opcode/100)%10 == 0 {
				p0 = a.program[p0]
			}
			p1 := a.program[i+2]
			if (opcode/1000)%10 == 0 {
				p1 = a.program[p1]
			}
			if p0 == 0 {
				i = p1
			} else {
				i += 3
			}
		case 7:
			p0 := a.program[i+1]
			if (opcode/100)%10 == 0 {
				p0 = a.program[p0]
			}
			p1 := a.program[i+2]
			if (opcode/1000)%10 == 0 {
				p1 = a.program[p1]
			}
			p2 := &a.program[i+3]
			if (opcode/10000)%10 == 0 {
				p2 = &a.program[*p2]
			}
			if p0 < p1 {
				*p2 = 1
			} else {
				*p2 = 0
			}
			i += 4
		case 8:
			p0 := a.program[i+1]
			if (opcode/100)%10 == 0 {
				p0 = a.program[p0]
			}
			p1 := a.program[i+2]
			if (opcode/1000)%10 == 0 {
				p1 = a.program[p1]
			}
			p2 := &a.program[i+3]
			if (opcode/10000)%10 == 0 {
				p2 = &a.program[*p2]
			}
			if p0 == p1 {
				*p2 = 1
			} else {
				*p2 = 0
			}
			i += 4

		case 99:
			return a.program[0]
		default:
			panic(fmt.Sprintf("bad opcode %d", opcode))
		}
	}
	panic("off the end")
}

var input = []int{
	3, 8, 1001, 8, 10, 8, 105, 1, 0, 0, 21, 34, 59, 76, 101, 114, 195, 276, 357, 438, 99999, 3, 9, 1001, 9, 4, 9, 1002, 9, 4, 9, 4, 9, 99, 3, 9, 102, 4, 9, 9, 101, 2, 9, 9, 102, 4, 9, 9, 1001, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 4, 9, 9, 102, 5, 9, 9, 101, 5, 9, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 1001, 9, 4, 9, 102, 4, 9, 9, 1001, 9, 4, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 1002, 9, 3, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 99, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 99, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1001, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 102, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 3, 9, 101, 1, 9, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1002, 9, 2, 9, 4, 9, 3, 9, 1001, 9, 1, 9, 4, 9, 3, 9, 101, 2, 9, 9, 4, 9, 99,
}

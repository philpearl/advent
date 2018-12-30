package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	part1("input.txt")
	// part2("input.txt")
	part2hack()
}

// Hmm, analysing the program, it looks like we're trying to find the sum of the factors of 10551378
func part2hack() {
	var r0 int

	for sc := 1; sc <= 10551378; sc++ {
		if 10551378%sc == 0 {
			r0 += sc
		}
	}

	fmt.Println(r0)
}

func part2(filename string) {
	program, ip := readProgram(filename)
	var regs registers
	regs[0] = 1
	lp := len(program)
	for {
		offset := regs[ip]
		if offset < 0 || offset > lp {
			break
		}
		i := program[offset]
		i.op(&regs, i.a, i.b, i.c)
		regs[ip]++
		fmt.Println(regs)
	}
	fmt.Println(regs[0])
}

func part1(filename string) {
	program, ip := readProgram(filename)
	var regs registers
	for {
		offset := regs[ip]
		if offset < 0 || offset > len(program) {
			break
		}
		i := program[offset]
		i.op(&regs, i.a, i.b, i.c)
		regs[ip]++
	}
	fmt.Println(regs[0])
}

type test struct {
	before registers
	after  registers
	instruction
}

type instruction struct {
	op      op
	a, b, c int
}

type registers [6]int

type op func(r *registers, a, b, c int)

var ops = map[string]op{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtri": gtri,
	"gtir": gtir,
	"gtrr": gtrr,
	"eqri": eqri,
	"eqir": eqir,
	"eqrr": eqrr,
}

func addr(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a] + (*r)[b]
}

func addi(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a] + b
}

func mulr(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a] * (*r)[b]
}

func muli(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a] * b
}

func banr(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a] & (*r)[b]
}

func bani(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a] & b
}

func borr(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a] | (*r)[b]
}

func bori(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a] | b
}

func setr(r *registers, a, b, c int) {
	(*r)[c] = (*r)[a]
}

func seti(r *registers, a, b, c int) {
	(*r)[c] = a
}

func gtir(r *registers, a, b, c int) {
	if a > (*r)[b] {
		(*r)[c] = 1
	} else {
		(*r)[c] = 0
	}
}

func gtri(r *registers, a, b, c int) {
	if (*r)[a] > b {
		(*r)[c] = 1
	} else {
		(*r)[c] = 0
	}
}

func gtrr(r *registers, a, b, c int) {
	if (*r)[a] > (*r)[b] {
		(*r)[c] = 1
	} else {
		(*r)[c] = 0
	}
}

func eqir(r *registers, a, b, c int) {
	if a == (*r)[b] {
		(*r)[c] = 1
	} else {
		(*r)[c] = 0
	}
}

func eqri(r *registers, a, b, c int) {
	if (*r)[a] == b {
		(*r)[c] = 1
	} else {
		(*r)[c] = 0
	}
}

func eqrr(r *registers, a, b, c int) {
	if (*r)[a] == (*r)[b] {
		(*r)[c] = 1
	} else {
		(*r)[c] = 0
	}
}

func readProgram(filename string) (program []instruction, ip int) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := fmt.Fscanf(f, "#ip %d\n", &ip); err != nil {
		log.Fatal(err)
	}

	line := 1
	for {
		var i instruction
		var opcode string
		line++
		_, err = fmt.Fscanf(f, "%s %d %d %d\n", &opcode, &i.a, &i.b, &i.c)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err, line)
		}
		i.op = ops[opcode]
		program = append(program, i)
	}

	return program, ip
}

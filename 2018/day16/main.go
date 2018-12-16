package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	part1("input.txt")
	part2("input.txt")
}

func part2(filename string) {
	tests := readInput(filename)
	matchedOpCodes := make(map[int]op)
	matchedOps := make(map[int]int)

	for len(matchedOps) < len(ops) {
		for _, tt := range tests {
			if _, matched := matchedOpCodes[tt.opcode]; matched {
				continue // already found this one
			}
			matches := findMatchingOps(tt)
			remaining := matches[:0]
			for _, m := range matches {
				if _, already := matchedOps[m]; !already {
					remaining = append(remaining, m)
				}
			}

			if len(remaining) == 0 {
				panic(fmt.Sprintf("no matches for %d!", tt.opcode))
			}
			if len(remaining) == 1 {
				fmt.Printf("found match for %d\n", tt.opcode)
				matchedOps[remaining[0]] = tt.opcode
				matchedOpCodes[tt.opcode] = ops[remaining[0]]
			}
		}
	}
	fmt.Printf("all ops matched\n")
	program := readProgram("program.txt")
	var regs registers
	for _, i := range program {
		matchedOpCodes[i.opcode](&regs, i.a, i.b, i.c)
	}
	fmt.Println(regs[0])
}

func part1(filename string) {
	tests := readInput(filename)

	var count int
	for _, tt := range tests {
		if countMatchingOps(tt) >= 3 {
			count++
		}
	}
	fmt.Println(count)
}

type test struct {
	before registers
	after  registers
	instruction
}

type instruction struct {
	opcode  int
	a, b, c int
}

func findMatchingOps(tt test) (matches []int) {
	for i, op := range ops {
		if worksForOp(tt, op) {
			matches = append(matches, i)
		}
	}

	return matches
}

func countMatchingOps(tt test) int {
	var count int
	for _, op := range ops {
		if worksForOp(tt, op) {
			count++
		}
	}
	return count
}

func worksForOp(tt test, op op) bool {
	r := tt.before
	op(&r, tt.a, tt.b, tt.c)
	return r == tt.after
}

type registers [4]int

type op func(r *registers, a, b, c int)

var ops = []op{
	addr,
	addi,
	mulr,
	muli,
	banr,
	bani,
	borr,
	bori,
	setr,
	seti,
	gtri,
	gtir,
	gtrr,
	eqri,
	eqir,
	eqrr,
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

func readInput(filename string) []test {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var tests []test
	line := 0
	for {
		var tt test
		line++
		_, err := fmt.Fscanf(f, "Before: [%d, %d, %d, %d]\n", &tt.before[0], &tt.before[1], &tt.before[2], &tt.before[3])
		if err != nil {
			if err == io.ErrUnexpectedEOF {
				break
			}
			log.Fatal(err, line)
		}
		line++
		_, err = fmt.Fscanf(f, "%d %d %d %d\n", &tt.opcode, &tt.a, &tt.b, &tt.c)
		if err != nil {
			log.Fatal(err, line)
		}

		line++
		_, err = fmt.Fscanf(f, "After: [%d, %d, %d, %d]\n\n", &tt.after[0], &tt.after[1], &tt.after[2], &tt.after[3])
		if err != nil {
			log.Fatal(err, line)
		}
		line++

		tests = append(tests, tt)
	}

	return tests
}

func readProgram(filename string) (program []instruction) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	line := 0
	for {
		var i instruction
		line++
		_, err = fmt.Fscanf(f, "%d %d %d %d\n", &i.opcode, &i.a, &i.b, &i.c)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err, line)
		}
		program = append(program, i)
	}

	return program
}

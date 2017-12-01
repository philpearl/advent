package main

import (
	"fmt"
	"strconv"
	"strings"
)

type register int

const (
	A register = iota
	B
	C
	D
	I
	N
)

type operand struct {
	number   int
	register register
}

func parseOperand(op string) operand {
	switch op[0] {
	case 'a':
		return operand{register: A}
	case 'b':
		return operand{register: B}
	case 'c':
		return operand{register: C}
	case 'd':
		return operand{register: D}
	default:
		n, _ := strconv.Atoi(op)
		return operand{register: N, number: n}
	}
}

func (o *operand) val(m *machine) int {
	if o.register == N {
		return o.number
	} else {
		return m.registers[o.register]
	}
}

type copy struct {
	from operand
	to   operand
}

func (c copy) execute(m *machine) {
	val := c.from.val(m)

	m.registers[c.to.register] = val
}

type inc struct {
	operand operand
}

func (i inc) execute(m *machine) {
	m.registers[i.operand.register]++
}

type dec struct {
	operand operand
}

func (d dec) execute(m *machine) {
	m.registers[d.operand.register]--
}

type jnz struct {
	test   operand
	offset operand
}

func (j jnz) execute(m *machine) {
	val := j.test.val(m)
	if val != 0 {
		offset := j.offset.val(m)
		m.ip--
		m.ip += offset
	}
}

type instruction interface {
	execute(m *machine)
}

type machine struct {
	registers    [4]int
	ip           int
	instructions []instruction
}

func (m *machine) run() {
	ll := len(m.instructions)
	for m.ip < ll {
		inst := m.instructions[m.ip]
		inst.execute(m)
		m.ip++
	}
}

func parseInstruction(line string) instruction {
	fields := strings.Fields(line)

	switch fields[0] {
	case "cpy":
		return copy{from: parseOperand(fields[1]), to: parseOperand(fields[2])}
	case "inc":
		return inc{operand: parseOperand(fields[1])}
	case "dec":
		return dec{operand: parseOperand(fields[1])}
	case "jnz":
		return jnz{test: parseOperand(fields[1]), offset: parseOperand(fields[2])}
	}
	return nil
}

func main() {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}

	m := machine{instructions: instructions}
	m.registers[C] = 1

	m.run()

	fmt.Println(m.registers[A])
}

var input = `cpy 1 a
cpy 1 b
cpy 26 d
jnz c 2
jnz 1 5
cpy 7 c
inc d
dec c
jnz c -2
cpy a c
inc a
dec b
jnz b -2
cpy c b
dec d
jnz d -6
cpy 17 c
cpy 18 d
inc a
dec d
jnz d -2
dec c
jnz c -5`

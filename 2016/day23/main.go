package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
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

func (o *operand) String() string {
	switch o.register {
	case A:
		return "a"
	case B:
		return "b"
	case C:
		return "c"
	case D:
		return "d"
	case I:
		return "i"
	case N:
		return strconv.Itoa(o.number)
	}
	return "X"
}

type copy struct {
	from operand
	to   operand
}

func (c copy) execute(m *machine) {
	val := c.from.val(m)

	m.registers[c.to.register] = val
}

func (c copy) String() string {
	return "cpy " + c.from.String() + " " + c.to.String()
}

type inc struct {
	operand operand
}

func (i inc) execute(m *machine) {
	m.registers[i.operand.register]++
}

func (i inc) String() string {
	return "inc " + i.operand.String()
}

type dec struct {
	operand operand
}

func (d dec) execute(m *machine) {
	m.registers[d.operand.register]--
}

func (d dec) String() string {
	return "dec " + d.operand.String()
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

func (j jnz) String() string {
	return "jnz " + j.test.String() + " " + j.offset.String()
}

type tgl struct {
	offset operand
}

func (t tgl) execute(m *machine) {
	offset := t.offset.val(m) + m.ip
	if offset < 0 || offset >= len(m.instructions) {
		return
	}
	instruction := m.instructions[offset]
	switch instruction := instruction.(type) {
	case copy:
		m.instructions[offset] = jnz{test: instruction.from, offset: instruction.to}
	case inc:
		m.instructions[offset] = dec{operand: instruction.operand}
	case dec:
		m.instructions[offset] = inc{operand: instruction.operand}
	case jnz:
		m.instructions[offset] = copy{from: instruction.test, to: instruction.offset}
	case tgl:
		m.instructions[offset] = inc{operand: instruction.offset}
	}
}

func (t tgl) String() string {
	return "tgl " + t.offset.String()
}

type instruction interface {
	execute(m *machine)
	String() string
}

type machine struct {
	registers    [4]int
	ip           int
	instructions []instruction
}

func (m *machine) run() {
	start := time.Now()
	ll := len(m.instructions)
	count := 0
	for m.ip < ll {
		inst := m.instructions[m.ip]
		// fmt.Println(inst)
		inst.execute(m)
		m.ip++
		count++
	}
	duration := time.Since(start)
	fmt.Printf("%d instructions in %s. %s per instruction\n", count, duration, duration/time.Duration(count))
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
	case "tgl":
		return tgl{offset: parseOperand(fields[1])}
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
	m.registers[A] = 12

	m.run()

	fmt.Println(m.registers[A])
}

var input = `cpy a b
dec b
cpy a d
cpy 0 a
cpy b c
inc a
dec c
jnz c -2
dec d
jnz d -5
dec b
cpy b c
cpy c d
dec d
inc c
jnz d -2
tgl c
cpy -16 c
jnz 1 c
cpy 73 c
jnz 71 d
inc a
inc d
jnz d -2
inc c
jnz c -5`

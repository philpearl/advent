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

type out struct {
	operand operand
}

func (o out) execute(m *machine) {
	val := o.operand.val(m)

	if !(val == 0 || val == 1) {
		fmt.Printf("\nbad out value %d\n", val)
		m.good = false
	}

	if m.outCount == 0 && val != 0 {
		fmt.Printf("\ndid not start at zero\n")
		m.good = false
	}

	if val == m.clock && m.outCount != 0 {
		fmt.Printf("\nclock did not change (after %d clocks)\n", m.outCount)
		m.good = false
	}

	m.clock = val
	m.outCount++
	fmt.Print(val)
}

func (o out) String() string {
	return "out " + o.operand.String()
}

type instruction interface {
	execute(m *machine)
	String() string
}

type machine struct {
	registers    [4]int
	ip           int
	instructions []instruction
	clock        int
	good         bool
	outCount     int
}

func (m *machine) run(max int) {
	start := time.Now()
	ll := len(m.instructions)
	count := 0
	for m.ip < ll && (max == 0 || count < max) && m.good {
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
	case "out":
		return out{operand: parseOperand(fields[1])}
	}
	return nil
}

func main() {
	lines := strings.Split(input, "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}

	for a := 0; a < 10000; a++ {
		m := machine{instructions: instructions, good: true}
		m.registers[A] = a

		fmt.Printf("a=%d\n", a)

		m.run(100000)

		if m.good {
			fmt.Printf("did ok with a = %d\n", a)
			return
		}
	}

}

var input = `cpy a d
cpy 4 c
cpy 643 b
inc d
dec b
jnz b -2
dec c
jnz c -5
cpy d a
jnz 0 0
cpy a b
cpy 0 a
cpy 2 c
jnz b 2
jnz 1 6
dec b
dec c
jnz c -4
inc a
jnz 1 -7
cpy 2 b
jnz c 2
jnz 1 4
dec b
dec c
jnz 1 -4
jnz 0 0
out b
jnz a -19
jnz 1 -21`

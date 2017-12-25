package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	part1(input)
	fmt.Printf("my %d", part2x())
}

func part2(program string) {
	m := newMachine()
	m.registers[a] = 1
	prog := parseProgram(program)
	m.run(prog)
	fmt.Printf("h is %d\n", m.registers[h])
}

func part1(program string) {
	m := newMachine()
	prog := parseProgram(program)
	fmt.Printf("%d mul executed\n", m.run(prog))
}

type register int

func (r register) String() string {
	return string([]byte{'a' + byte(r)})
}

const (
	a register = iota
	b
	c
	d
	e
	f
	g
	h
	reg_ip
	reg_n
)

type operand struct {
	val int
	reg register
}

func (o operand) value(m *machine) int {
	if o.reg == reg_n {
		return o.val
	}
	return m.registers[o.reg]
}

type set struct {
	r register
	o operand
}

func (s set) exec(m *machine) bool {
	m.registers[s.r] = s.o.value(m)
	return false
}

type add struct {
	r register
	o operand
}

func (s add) exec(m *machine) bool {
	m.registers[s.r] += s.o.value(m)
	return false
}

type sub struct {
	r register
	o operand
}

func (s sub) exec(m *machine) bool {
	m.registers[s.r] -= s.o.value(m)
	return false
}

type mul struct {
	r register
	o operand
}

func (s mul) exec(m *machine) bool {
	m.registers[s.r] *= s.o.value(m)
	return false
}

type jnz struct {
	o1 operand
	o2 operand
}

func (s jnz) exec(m *machine) bool {
	if s.o1.value(m) != 0 {
		m.registers[reg_ip] += s.o2.value(m) - 1
	}
	return false
}

type instruction interface {
	exec(m *machine) (blocked bool)
}

type machine struct {
	registers []int
}

func newMachine() *machine {
	return &machine{
		registers: make([]int, reg_ip+1),
	}
}

func parseOp(text string) operand {
	if text[0] >= 'a' && text[0] <= 'z' {
		return operand{
			reg: parseReg(text),
		}
	}
	val, _ := strconv.Atoi(text)
	return operand{
		reg: reg_n,
		val: val,
	}
}

func parseReg(text string) register {
	return register(text[0]-'a') + a
}

func parseInstruction(text string) instruction {
	parts := strings.Fields(text)
	switch parts[0] {
	case "set":
		return set{r: parseReg(parts[1]), o: parseOp(parts[2])}
	case "add":
		return add{r: parseReg(parts[1]), o: parseOp(parts[2])}
	case "sub":
		return sub{r: parseReg(parts[1]), o: parseOp(parts[2])}
	case "mul":
		return mul{r: parseReg(parts[1]), o: parseOp(parts[2])}
	case "jnz":
		return jnz{o1: parseOp(parts[1]), o2: parseOp(parts[2])}
	}
	panic("can't deal with " + text)
}

func parseProgram(text string) []instruction {
	lines := strings.Split(text, "\n")
	instructions := make([]instruction, len(lines))
	for i, line := range lines {
		instructions[i] = parseInstruction(line)
	}
	return instructions
}

func (m *machine) run(instructions []instruction) int {
	mulExecuted := 0
	for {
		ip := m.registers[reg_ip]
		if ip >= len(instructions) || ip < 0 {
			break
		}
		instruction := instructions[ip]
		// fmt.Printf("instruction %#v\n", instruction)
		// fmt.Printf("regs %#v\n", m.registers)
		m.registers[reg_ip]++
		instruction.exec(m)
		if _, ok := instruction.(mul); ok {
			mulExecuted++
		}
	}
	return mulExecuted
}

var input = `set b 79
set c b
jnz a 2
jnz 1 5
mul b 100
sub b -100000
set c b
sub c -17000
set f 1
set d 2
set e 2
set g d
mul g e
sub g b
jnz g 2
set f 0
sub e -1
set g e
sub g b
jnz g -8
sub d -1
set g d
sub g b
jnz g -13
jnz f 2
sub h -1
set g b
sub g c
jnz g 2
jnz 1 3
sub b -17
jnz 1 -23`

func part2x() int {
	h := 0
	for b := 107900; b <= 124900; b += 17 {
		f := 1
		// h is number of non-prime bs
		for d := 2; d < b; d++ {
			if b%d == 0 {
				f = 0
				break
			}
		}

		if f == 1 {
			fmt.Printf("b is %d\n", b)
		}

		if f == 0 {
			h++
		}
	}

	return h
}

// set f 1 // a=1 b=107900 c=124900 f=1
// set d 2 // a=1 b=107900 c=124900 d=2 f=1
// +set e 2 // a=1 b=107900 c=124900 d=2 e=2 f=1

// ++set g d
// mul g e
// sub g b
// jnz g 2  if b = d*e then f = 0
// set f 0
// sub e -1   e++
// set g e
// sub g b
// ++jnz g -8 if e != b  back to start

// sub d -1   d++
// set g d
// sub g b
// +jnz g -13  if d != b back to start
// jnz f 2
// sub h -1  if f == 0 h++
// set g b
// sub g c
// jnz g 2  if b == c then exit
// jnz 1 3
// sub b -17  b += 17
// jnz 1 -23  back to start

// g = d
// g = d * e
// if d * e - b == 0
// f = 0
// e++
// g = e - b

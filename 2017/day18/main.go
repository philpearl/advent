package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {

	// part1(input)
	part2(`snd 1
		snd 2
		snd p
		rcv a
		rcv b
		rcv c
		rcv d`)
	part2(input)
}

func part2(program string) {
	prog := parseProgram(program)
	m0 := newMachine(27)
	m1 := newMachine(27)
	m0.partner = m1
	m1.partner = m0
	m0.registers['p'-'a'+a] = 0
	m1.registers['p'-'a'+a] = 1

	// run each program in turn until it blocks. If a program remains blocked without
	// executing any instructions we're stuck
	total := 0
	for {
		c := m0.run(prog)
		total += c
		if c == 0 {
			fmt.Printf("no instructions run on m0\n")
			break
		}
		c = m1.run(prog)
		total += c
		if c == 0 {
			fmt.Printf("no instructions run on m1\n")
			break
		}

		if total%10000 == 0 {
			fmt.Printf("%d instructions run\n", total)
		}
	}

	fmt.Printf("m1 sent %d times after %d instructions\n", m1.sends, total)
}

func part1(program string) {
	m := newMachine(27)
	prog := parseProgram(program)
	m.run(prog)
}

type register int

func (r register) String() string {
	return string([]byte{'a' + byte(r)})
}

const (
	a      register = 0
	reg_ip register = 26
	reg_n  register = 27
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

type snd struct {
	o operand
}

func (s snd) exec(m *machine) bool {
	if m.qInUse {
		return true
	}

	m.send(s.o.value(m))
	m.sends++
	return false
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

type mul struct {
	r register
	o operand
}

func (s mul) exec(m *machine) bool {
	m.registers[s.r] *= s.o.value(m)
	return false
}

type mod struct {
	r register
	o operand
}

func (s mod) exec(m *machine) bool {
	m.registers[s.r] %= s.o.value(m)
	return false
}

type rcv struct {
	r register
}

func (s rcv) exec(m *machine) bool {

	if ok, v := m.partner.pop(); ok {
		m.registers[s.r] = v
		m.waitingToRcv = false
		return false
	}
	m.waitingToRcv = true
	return true
}

type jgz struct {
	o1 operand
	o2 operand
}

func (s jgz) exec(m *machine) bool {
	if s.o1.value(m) > 0 {
		m.registers[reg_ip] += s.o2.value(m) - 1
	}
	return false
}

type instruction interface {
	exec(m *machine) (blocked bool)
}

type machine struct {
	partner      *machine
	qInUse       bool
	q            []int
	sends        int
	waitingToRcv bool
	registers    []int
}

func (m *machine) send(v int) {
	m.q = append(m.q, v)
}

func (m *machine) pop() (bool, int) {
	if len(m.q) == 0 {
		return false, 0
	}
	v := m.q[0]
	m.q = m.q[1:]
	return true, v
}

func newMachine(regs int) *machine {
	return &machine{
		registers: make([]int, regs),
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
	case "snd":
		return snd{o: parseOp(parts[1])}
	case "set":
		return set{r: parseReg(parts[1]), o: parseOp(parts[2])}
	case "add":
		return add{r: parseReg(parts[1]), o: parseOp(parts[2])}
	case "mul":
		return mul{r: parseReg(parts[1]), o: parseOp(parts[2])}
	case "mod":
		return mod{r: parseReg(parts[1]), o: parseOp(parts[2])}
	case "rcv":
		return rcv{r: parseReg(parts[1])}
	case "jgz":
		return jgz{o1: parseOp(parts[1]), o2: parseOp(parts[2])}
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
	executed := 0
	for {
		instruction := instructions[m.registers[reg_ip]]
		// fmt.Printf("instruction %#v\n", instruction)
		m.registers[reg_ip]++
		if blocked := instruction.exec(m); blocked {
			m.registers[reg_ip]--
			return executed
		}
		executed++
		// fmt.Printf("regs: %#v\n", m.registers)
	}
}

var input = `set i 31
set a 1
mul p 17
jgz p p
mul a 2
add i -1
jgz i -2
add a -1
set i 127
set p 680
mul p 8505
mod p a
mul p 129749
add p 12345
mod p a
set b p
mod b 10000
snd b
add i -1
jgz i -9
jgz a 3
rcv b
jgz b -1
set f 0
set i 126
rcv a
rcv b
set p a
mul p -1
add p b
jgz p 4
snd a
set a b
jgz 1 3
snd b
set f 1
add i -1
jgz i -11
snd a
jgz f -16
jgz a -19`

package main

import "fmt"

func main() {
	part1()
}

func part1() {
	m := new()
	for i := 0; i < steps; i++ {
		m.step()
	}
	fmt.Printf("chksum is %d\n", m.checksum())
}

type machine struct {
	state   state
	tape    []uint8
	current int
}

func new() *machine {
	return &machine{
		state:   A,
		tape:    make([]uint8, 1000000),
		current: 500000,
	}
}

func (m *machine) step() {
	a := states[m.state][m.tape[m.current]]
	m.state = a.state
	m.tape[m.current] = a.val
	if a.move == left {
		m.current--
	} else {
		m.current++
	}
}

func (m *machine) checksum() int {
	count := 0
	for _, v := range m.tape {
		if v == 1 {
			count++
		}
	}
	return count
}

type state uint8

const (
	A state = iota
	B
	C
	D
	E
	F
)

type dirn uint8

const (
	left dirn = iota
	right
)

type action struct {
	val   uint8
	move  dirn
	state state
}

type actions [2]action

var steps = 12173597

// Begin in state A.
// Perform a diagnostic checksum after 12173597 steps.

var states = []actions{
	{
		{val: 1, move: right, state: B},
		{val: 0, move: left, state: C},
	},
	{
		{val: 1, move: left, state: A},
		{val: 1, move: right, state: D},
	},
	{
		{val: 1, move: right, state: A},
		{val: 0, move: left, state: E},
	},
	{
		{val: 1, move: right, state: A},
		{val: 0, move: right, state: B},
	},
	{
		{val: 1, move: left, state: F},
		{val: 1, move: left, state: C},
	},
	{
		{val: 1, move: right, state: D},
		{val: 1, move: right, state: A},
	},
}

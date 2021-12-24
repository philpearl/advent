package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	if err := part1(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	if err := part2(input); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func part1(input string) error {
	w := parseInput(input)
	w.run(10)
	return nil
}

func part2(input string) error {
	w := parseInput(input)
	w.run(40)
	return nil
}

func (w *world) run(depth int) {
	fmt.Println(w.polymer)

	lastC := w.polymer[0]
	var counts counts
	for i := range w.polymer[1:] {
		c := w.polymer[i+1]
		counts.add(w.countsFor(lastC, c, depth))
		lastC = c
	}
	// decrement counts for internal letters
	for i := range w.polymer[1 : len(w.polymer)-1] {
		c := w.polymer[1+i]
		counts[c-'A']--
	}

	counts.print()

	max, min := counts.maxMin()
	fmt.Println(min, max, max-min)
}

func (w *world) countsFor(a, b byte, depth int) (counts counts) {
	if depth == 0 {
		counts[a-'A'] = 1
		counts[b-'A'] += 1
		return counts
	}

	p := params{a: a, b: b, depth: depth}
	if c, ok := w.memo[p]; ok {
		return c
	}

	inc, ok := w.insertions[[2]byte{a, b}]
	if !ok {
		// It turns out there's an insertion for every combination
		panic("no insertion")
	}
	c1 := w.countsFor(a, inc, depth-1)
	c2 := w.countsFor(inc, b, depth-1)
	c1.add(c2)
	c1[inc-'A']--

	w.memo[p] = c1
	return c1
}

type counts [26]int

func (c counts) maxMin() (max, min int) {
	min = math.MaxInt
	for _, c := range c {
		if c == 0 {
			continue
		}
		if c < min {
			min = c
		}
		if c > max {
			max = c
		}
	}
	return max, min
}

func (c *counts) add(c2 counts) {
	for i := range c {
		c[i] += c2[i]
	}
}

func (c counts) print() {
	for i, c := range c {
		if c == 0 {
			continue
		}
		fmt.Printf("%c %d\n", i+'A', c)
	}
}

type params struct {
	a, b  byte
	depth int
}

type world struct {
	polymer string

	insertions map[[2]byte]byte

	memo map[params][26]int
}

func parseInput(input string) (w world) {
	lines := strings.Split(input, "\n")

	w.polymer = lines[0]
	w.insertions = make(map[[2]byte]byte)
	w.memo = make(map[params][26]int)

	for _, line := range lines[2:] {
		var a, b, c byte
		if n, err := fmt.Sscanf(line, "%c%c -> %c", &a, &b, &c); err != nil {
			panic(err)
		} else if n != 3 {
			panic("should be 3")
		}
		w.insertions[[2]byte{a, b}] = c
	}
	return w
}

var input = `OOBFPNOPBHKCCVHOBCSO

NS -> H
NN -> P
FF -> O
HF -> C
KN -> V
PO -> B
PS -> B
FB -> N
ON -> F
OK -> F
OO -> K
KS -> F
FN -> F
KC -> H
NC -> N
NB -> C
KH -> S
SV -> B
BC -> S
KB -> B
SC -> S
KP -> H
FS -> K
NK -> K
OC -> H
NH -> C
PH -> F
OS -> V
BB -> C
CC -> F
CF -> H
CP -> V
VB -> N
VC -> F
PK -> V
NV -> N
FO -> S
CK -> O
BH -> K
PN -> B
PP -> S
NF -> B
SF -> K
VF -> H
HS -> F
NP -> F
SH -> V
SK -> K
PC -> V
BO -> H
HN -> P
BK -> O
BP -> S
OP -> N
SP -> N
KK -> C
HB -> H
OF -> C
VH -> C
HO -> N
FK -> V
NO -> H
KF -> S
KO -> V
PF -> K
HV -> C
SO -> F
SS -> F
VN -> K
HH -> B
OB -> S
CH -> B
FH -> B
CO -> V
HK -> F
VK -> K
CN -> V
SB -> K
PV -> O
PB -> F
VV -> P
CS -> N
CB -> C
BS -> F
HC -> B
SN -> P
VP -> P
OV -> P
BV -> P
FC -> N
KV -> S
CV -> F
BN -> S
BF -> C
OH -> F
VO -> B
FP -> S
FV -> V
VS -> N
HP -> B`

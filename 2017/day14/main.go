package main

import "fmt"
import "math/bits"

func main() {

	fmt.Printf("%d bits set\n", part1("flqrgnkx"))
	fmt.Printf("%d bits set\n", part1(input))

	fmt.Printf("%d groups\n", part2("flqrgnkx"))
	fmt.Printf("%d groups\n", part2(input))
}

func part2(input string) int {
	rows := [128]hash{}
	for i := range rows {
		rows[i] = knothash(fmt.Sprintf("%s-%d", input, i))
	}

	uf := newuf(128 * 128)

	for x, r := range rows {
		for y := 0; y < 128; y++ {
			if r.isSet(y) {
				pos := node(x + y*128)
				n1 := uf.find(pos)
				if y < 127 && r.isSet(y+1) {
					// x,y and x,y+1 are connected
					n2 := uf.find(pos + 128)
					uf.union(n1, n2)
				}
				n1 = uf.find(pos)
				if x < 127 && rows[x+1].isSet(y) {
					// x,y and x+1,y are connected
					n2 := uf.find(pos + 1)
					uf.union(n1, n2)
				}
			}
		}
	}

	return uf.countGroups()
}

type hash [16]uint8

func (h hash) isSet(index int) bool {
	byt := index / 8
	bit := 7 - uint(index%8)
	return (h[byt] & (1 << bit)) != 0
}

func part1(input string) int {
	count := 0
	for i := 0; i < 128; i++ {
		h := knothash(fmt.Sprintf("%s-%d", input, i))
		for _, x := range h {
			count += bits.OnesCount8(x)
		}
	}
	return count
}

var input = "ugkiagan"

func knothash(input string) hash {
	l := new()
	extents := append([]byte(input), 17, 31, 73, 47, 23)
	for round := 0; round < 64; round++ {
		for _, extent := range extents {
			l.reverse(extent)
			l.update(extent)
		}
	}

	return densify(l.data)
}

type list struct {
	data     [256]uint8
	current  uint8
	skipSize uint8
}

func new() *list {
	l := &list{}
	for i := range l.data {
		l.data[i] = uint8(i)
	}
	return l
}

func densify(in [256]uint8) hash {
	d := hash{}
	for i := 0; i < 16; i++ {
		var x uint8
		for j := 0; j < 16; j++ {
			x ^= in[i*16+j]
		}
		d[i] = x
	}
	return d
}

func (l *list) reverse(extent uint8) {
	// Current 0
	// extent 3
	// l = 0, r = 5 + 3 - i % 5
	for i := uint8(0); i < (extent / 2); i++ {
		left := (l.current + i)
		right := (l.current + extent - i - 1)
		l.data[right], l.data[left] = l.data[left], l.data[right]
	}
}

func (l *list) update(extent uint8) {
	l.current += extent + l.skipSize
	l.skipSize++
}

type node int16

type unionfind []node

func newuf(size int16) unionfind {
	u := make(unionfind, size+1)
	for i := range u {
		u[i] = 0
	}
	return u
}

func (u unionfind) countGroups() int {
	count := 0
	for _, n := range u {
		if n < 0 {
			count++
		}
	}
	return count
}

func (u unionfind) find(a node) node {
	// fmt.Printf("find %d\n", a)
	a++
	if u[a] == 0 {
		u[a] = -1
	}
	for {
		n := u[a]
		if n < 0 {
			return a
		}
		if n == 0 {
			panic(fmt.Sprintf("u[%d]=0", a))
		}
		a = n
	}
}

func (u unionfind) union(a, b node) {
	if a <= 0 || b <= 0 {
		panic("a,b small")
	}
	if a == b {
		return
	}
	// fmt.Printf("join %d & %d with values %d %d", a, b, u[a], u[b])
	if u[b] < u[a] {
		u[b] += u[a]
		u[a] = b
	} else {
		// na is bigger or they're the same
		u[a] += u[b]
		u[b] = a
	}
}

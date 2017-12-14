package main

import "fmt"

func main() {
	fmt.Printf("part1 result %d\n", part1(256, input))
	fmt.Printf("part2 is %s\n", part2(256, input2))
}

func part2(length int, input string) string {
	l := new(length)
	extents := append([]byte(input), 17, 31, 73, 47, 23)
	for round := 0; round < 64; round++ {
		for _, extent := range extents {
			l.reverse(extent)
			l.update(extent)
		}
	}

	hash := densify(l.data[0:256])
	return fmt.Sprintf("%x", hash[:])
}

type list struct {
	data     []uint8
	current  uint8
	skipSize uint8
}

func part1(length int, lengths []uint8) int {
	l := new(length)
	for _, extent := range lengths {
		fmt.Printf("%#v current %d, extent %d\n", l.data, l.current, extent)
		l.reverse(extent)
		l.update(extent)
		fmt.Printf("%#v current %d, extent %d\n", l.data, l.current, extent)
	}

	return int(l.data[0]) * int(l.data[1])
}

func new(length int) *list {
	data := make([]uint8, length)
	for i := range data {
		data[i] = uint8(i)
	}
	return &list{
		data: data,
	}
}

func densify(in []uint8) [16]uint8 {
	d := [16]uint8{}
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
		if ll := len(l.data); ll < 256 {
			left = left % uint8(ll)
			right = right % uint8(ll)
		}
		l.data[right], l.data[left] = l.data[left], l.data[right]
	}
}

func (l *list) update(extent uint8) {
	l.current += extent + l.skipSize
	if ll := len(l.data); ll < 256 {
		l.current %= uint8(ll)
	}

	l.skipSize++
}

var input = []uint8{
	227, 169, 3, 166, 246, 201, 0, 47, 1, 255, 2, 254, 96, 3, 97, 144,
}
var input2 = "227,169,3,166,246,201,0,47,1,255,2,254,96,3,97,144"

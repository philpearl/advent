package main

import (
	"fmt"
	"strings"
)

func main() {
	part1(input2, input3)
}

func part1(in1, in2 string) {
	s2m := parseInput2(in1)
	s3m := parseInput3(in2)

	image := []byte(".#...####")
	side := 3
	for i := 0; i < 18; i++ {
		if side%2 == 0 {
			image = s2m.grow(image, side)
			side = side * 3 / 2
		} else {
			image = s3m.grow(image, side)
			side = side * 4 / 3
		}
	}

	count := 0
	for _, p := range image {
		if p == '#' {
			count++
		}
	}
	fmt.Printf("%d pixels on\n", count)
}

func (m square2map) print() {
	for k, v := range m {
		k.print()
		v.print()
		fmt.Printf("\n")
	}
}

func (m square2map) grow(image []byte, side int) []byte {
	oside := side * 3 / 2
	out := make([]byte, oside*oside)
	for x := 0; x < side; x += 2 {
		for y := 0; y < side; y += 2 {
			s := square2{
				image[x+y*side], image[x+1+y*side],
				image[x+(y+1)*side], image[x+1+(y+1)*side],
			}

			o, ok := m[s]
			if !ok {
				fmt.Printf("no match for %s\n", string(s[:]))
			}
			ox := x * 3 / 2
			oy := y * 3 / 2
			copy(out[ox+oy*oside:ox+3+oy*oside], o[:3])
			copy(out[ox+(oy+1)*oside:ox+3+(oy+1)*oside], o[3:6])
			copy(out[ox+(oy+2)*oside:ox+3+(oy+2)*oside], o[6:9])
		}
	}

	return out
}

func (m square3map) grow(image []byte, side int) []byte {
	oside := side * 4 / 3
	out := make([]byte, oside*oside)
	for x := 0; x < side; x += 3 {
		for y := 0; y < side; y += 3 {
			s := square3{
				image[x+y*side], image[x+1+y*side], image[x+2+y*side],
				image[x+(y+1)*side], image[x+1+(y+1)*side], image[x+2+(y+1)*side],
				image[x+(y+2)*side], image[x+1+(y+2)*side], image[x+2+(y+2)*side],
			}

			o, ok := m[s]
			if !ok {
				fmt.Printf("no match for %s\n", string(s[:]))
			}

			ox := x * 4 / 3
			oy := y * 4 / 3
			copy(out[ox+oy*oside:ox+4+oy*oside], o[:4])
			copy(out[ox+(oy+1)*oside:ox+4+(oy+1)*oside], o[4:8])
			copy(out[ox+(oy+2)*oside:ox+4+(oy+2)*oside], o[8:12])
			copy(out[ox+(oy+3)*oside:ox+4+(oy+3)*oside], o[12:16])
		}
	}

	return out
}

type square2 [4]byte
type square3 [9]byte
type square4 [16]byte

func (s square2) rotate() square2 {
	// 0 1  2 0
	// 2 3  3 1
	return square2{
		s[2], s[0], s[3], s[1],
	}
}

func (s square2) flip() square2 {
	return square2{
		s[1], s[0], s[3], s[2],
	}
}

func (s square2) print() {
	fmt.Println(string(s[:2]))
	fmt.Println(string(s[2:4]))
}

func (s square3) rotate() square3 {
	// 0 1 2      6 3 0
	// 3 4 5  =>  7 4 1
	// 6 7 8      8 5 2
	return square3{
		s[6], s[3], s[0],
		s[7], s[4], s[1],
		s[8], s[5], s[2],
	}
}

func (s square3) flip() square3 {
	// 0 1 2      2 1 0
	// 3 4 5  =>  5 4 3
	// 6 7 8      8 7 6
	return square3{
		s[2], s[1], s[0],
		s[5], s[4], s[3],
		s[8], s[7], s[6],
	}
}

func (s square3) print() {
	fmt.Println(string(s[:3]))
	fmt.Println(string(s[3:6]))
	fmt.Println(string(s[6:9]))
}

func (s square4) print() {
	fmt.Println(string(s[:4]))
	fmt.Println(string(s[4:8]))
	fmt.Println(string(s[8:12]))
	fmt.Println(string(s[12:16]))
}

type square2map map[square2]square3
type square3map map[square3]square4

func parseInput2(in string) square2map {
	lines := strings.Split(in, "\n")
	m := make(square2map, len(lines))

	for _, l := range lines {
		s2 := parseSquare2(l)
		s3 := parseSquare3(l[9:])
		for i := 0; i < 4; i++ {
			m[s2] = s3
			s2 = s2.rotate()
		}
		s2 = s2.flip()
		for i := 0; i < 4; i++ {
			m[s2] = s3
			s2 = s2.rotate()
		}
	}
	return m
}

func parseInput3(in string) square3map {
	lines := strings.Split(in, "\n")
	m := make(square3map, len(lines))

	for _, l := range lines {
		s3 := parseSquare3(l)
		s4 := parseSquare4(l[15:])
		for i := 0; i < 4; i++ {
			m[s3] = s4
			s3 = s3.rotate()
		}
		s3 = s3.flip()
		for i := 0; i < 4; i++ {
			m[s3] = s4
			s3 = s3.rotate()
		}
	}
	return m
}

func parseSquare2(in string) square2 {
	return square2{in[0], in[1], in[3], in[4]}
}

func parseSquare3(in string) square3 {
	return square3{in[0], in[1], in[2], in[4], in[5], in[6], in[8], in[9], in[10]}
}

func parseSquare4(in string) square4 {
	return square4{in[0], in[1], in[2], in[3],
		in[5], in[6], in[7], in[8],
		in[10], in[11], in[12], in[13],
		in[15], in[16], in[17], in[18],
	}
}

var input2 = `../.. => .##/.##/###
#./.. => .../#.#/###
##/.. => .##/.../.#.
.#/#. => ###/.#./##.
##/#. => .#./#../#.#
##/## => .##/#.#/###`
var input3 = `.../.../... => ####/.##./####/.#..
#../.../... => ..../..##/#.../.##.
.#./.../... => #.#./##.#/#.../#.#.
##./.../... => .#../.##./#.../....
#.#/.../... => ###./..##/..##/##.#
###/.../... => .###/#.##/..../....
.#./#../... => ##.#/#..#/.##./...#
##./#../... => ..../#..#/#.#./...#
..#/#../... => #.##/.#../.#.#/###.
#.#/#../... => ##../.#.#/...#/...#
.##/#../... => ##.#/.##./..#./##.#
###/#../... => ...#/####/..#./#...
.../.#./... => ##.#/#.#./..##/.##.
#../.#./... => .#.#/#.##/.##./....
.#./.#./... => #..#/#.../.##./....
##./.#./... => ###./###./..##/#..#
#.#/.#./... => .###/...#/###./###.
###/.#./... => ...#/..##/..#./#.##
.#./##./... => .##./.#../...#/..#.
##./##./... => .###/..#./.###/###.
..#/##./... => .#.#/..#./..#./...#
#.#/##./... => .#.#/##../#.../.##.
.##/##./... => .##./...#/#.##/###.
###/##./... => ...#/###./####/#.##
.../#.#/... => #.#./#.../#.#./..#.
#../#.#/... => ###./##../..#./.#..
.#./#.#/... => #.../..##/#..#/#.#.
##./#.#/... => #.#./.##./#..#/##.#
#.#/#.#/... => #.##/.#.#/#..#/.#.#
###/#.#/... => #.../##.#/###./....
.../###/... => ..##/...#/##.#/###.
#../###/... => .#.#/...#/#.##/.#..
.#./###/... => ####/#.../..#./.#.#
##./###/... => ..../####/#.##/#..#
#.#/###/... => ####/..#./####/.#.#
###/###/... => ..##/..../...#/.#..
..#/.../#.. => .###/..##/.#.#/.##.
#.#/.../#.. => #.##/#..#/.#.#/##.#
.##/.../#.. => #.##/####/.#.#/..#.
###/.../#.. => ##../##.#/..../##..
.##/#../#.. => ...#/####/..##/.##.
###/#../#.. => ..#./...#/#.../##.#
..#/.#./#.. => #..#/##.#/..##/#..#
#.#/.#./#.. => ..../.###/#..#/..##
.##/.#./#.. => ..#./...#/..##/...#
###/.#./#.. => ...#/..../##.#/....
.##/##./#.. => .#../..##/...#/.#.#
###/##./#.. => .###/#.#./####/#.#.
#../..#/#.. => .###/##.#/##../##..
.#./..#/#.. => ##../.#../###./##.#
##./..#/#.. => #..#/####/####/..##
#.#/..#/#.. => ..##/..../###./..##
.##/..#/#.. => ..##/.#.#/.#../.#..
###/..#/#.. => ...#/.###/.###/.#.#
#../#.#/#.. => ##../##../##.#/.##.
.#./#.#/#.. => ...#/.##./.#.#/#...
##./#.#/#.. => .##./.#../.#../#...
..#/#.#/#.. => ..##/##.#/####/###.
#.#/#.#/#.. => ..../.###/#.../#..#
.##/#.#/#.. => ..#./#.#./.#../...#
###/#.#/#.. => ##.#/#.../##.#/.##.
#../.##/#.. => ..../#.../..#./####
.#./.##/#.. => #..#/.#../#.#./..##
##./.##/#.. => .###/..##/###./....
#.#/.##/#.. => .###/.##./.###/#.##
.##/.##/#.. => #.##/###./.##./...#
###/.##/#.. => ...#/#.##/.##./#.#.
#../###/#.. => #..#/.###/.###/#.#.
.#./###/#.. => ..#./#.#./..../...#
##./###/#.. => ..##/##../#..#/....
..#/###/#.. => ..##/.#../.#../###.
#.#/###/#.. => ..#./.###/..../...#
.##/###/#.. => .##./###./#.../#.##
###/###/#.. => ##.#/..../.##./##.#
.#./#.#/.#. => .##./.#.#/####/....
##./#.#/.#. => ##.#/#.##/####/.#..
#.#/#.#/.#. => ####/.##./##.#/...#
###/#.#/.#. => #..#/#.##/.##./###.
.#./###/.#. => .#../..../.##./##.#
##./###/.#. => ##.#/.#../#.../.###
#.#/###/.#. => ###./###./.#../###.
###/###/.#. => #..#/#.../#..#/.#.#
#.#/..#/##. => #..#/#.../##../###.
###/..#/##. => #.../.#../.###/#...
.##/#.#/##. => .#.#/.##./.#../##.#
###/#.#/##. => #.../..../##../.###
#.#/.##/##. => .#.#/##../.###/#.#.
###/.##/##. => ###./..#./##.#/.###
.##/###/##. => ..#./.#.#/##.#/#.#.
###/###/##. => ##../.#.#/#..#/.#.#
#.#/.../#.# => ##../###./..#./##.#
###/.../#.# => .#../##../..#./##.#
###/#../#.# => ###./#..#/####/....
#.#/.#./#.# => .###/..../.###/##.#
###/.#./#.# => ###./.###/..##/.#.#
###/##./#.# => ..#./..##/#..#/#.##
#.#/#.#/#.# => .#.#/.#../.#.#/#.##
###/#.#/#.# => .###/#.../##../.###
#.#/###/#.# => .#../...#/..../...#
###/###/#.# => #..#/##.#/..#./#...
###/#.#/### => .###/.#.#/..#./####
###/###/### => ##.#/..##/.#../..##`
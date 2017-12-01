package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	width  = 50
	height = 6
)

// screen is 50 wide by 6 high. 0,0 is top left
type screen [width][height]bool

// rect turns on pixels in top left of screen in rectangle a wide and b tall
func (s *screen) rect(a, b int) {
	for x := 0; x < a; x++ {
		for y := 0; y < b; y++ {
			s[x][y] = true
		}
	}
}

// rotateRow shifts row y right by by pixels. pixels that go off the end of the row
// wrap around to the start of the row
func (s *screen) rotateRow(y int, by int) {
	for i := 0; i < by; i++ {
		s.rotateRowOne(y)
	}
}

func (s *screen) rotateRowOne(y int) {
	last := false
	for x := 0; x < width; x++ {
		curr := s[x][y]
		if x > 0 {
			s[x][y] = last
		}
		last = curr
	}
	s[0][y] = last

}

// rotateColumn shifts column x down by by pixels, wrapping again
func (s *screen) rotateColumn(x int, by int) {
	for i := 0; i < by; i++ {
		s.rotateColumnOne(x)
	}
}

func (s *screen) rotateColumnOne(x int) {
	last := false
	for y := 0; y < height; y++ {
		curr := s[x][y]
		if y > 0 {
			s[x][y] = last
		}
		last = curr
	}
	s[x][0] = last
}

func (s *screen) count() int {
	count := 0
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if s[x][y] {
				count++
			}
		}
	}
	return count
}

func (s *screen) show() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if s[x][y] {
				io.WriteString(os.Stdout, "#")
			} else {
				io.WriteString(os.Stdout, " ")
			}
			if x%5 == 4 {
				io.WriteString(os.Stdout, "  ")
			}
		}
		io.WriteString(os.Stdout, "\n")
	}
}

func (s *screen) unshow() {
	fmt.Printf("\033[%dA", height)
}

func main() {

	s := screen{}
	// s.rect(3, 2)
	// s.rotateColumn(1, 1)
	// s.rotateRow(0, 4)
	// s.rotateColumn(1, 1)

	// s.show()

	for i := 0; i < height; i++ {
		io.WriteString(os.Stdout, "\n")
	}

	for _, instruction := range strings.Split(input, "\n") {
		s.unshow()

		words := strings.Fields(instruction)

		switch words[0] {
		case "rotate":
			switch words[1] {
			case "row":
				s.rotateRow(rotateParams(words[2:]))
			case "column":
				s.rotateColumn(rotateParams(words[2:]))
			}
		case "rect":
			s.rect(xByy(words[1]))
		}

		s.show()
		time.Sleep(time.Millisecond * 10)
	}

	fmt.Printf("%d lit\n", s.count())
}

func rotateParams(parts []string) (xy, by int) {
	by, _ = strconv.Atoi(parts[2])
	xy, _ = strconv.Atoi(strings.Split(parts[0], "=")[1])
	return
}

func xByy(val string) (x, y int) {
	parts := strings.Split(val, "x")

	x, _ = strconv.Atoi(parts[0])
	y, _ = strconv.Atoi(parts[1])
	return
}

var input = `rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 5
rect 1x1
rotate row y=0 by 3
rect 1x1
rotate row y=0 by 2
rect 1x1
rotate row y=0 by 3
rect 2x1
rotate row y=0 by 2
rect 1x2
rotate row y=1 by 5
rotate row y=0 by 3
rect 1x2
rotate column x=30 by 1
rotate column x=25 by 1
rotate column x=10 by 1
rotate row y=1 by 5
rotate row y=0 by 2
rect 1x2
rotate row y=0 by 5
rotate column x=0 by 1
rect 4x1
rotate row y=2 by 18
rotate row y=0 by 5
rotate column x=0 by 1
rect 3x1
rotate row y=2 by 12
rotate row y=0 by 5
rotate column x=0 by 1
rect 4x1
rotate column x=20 by 1
rotate row y=2 by 5
rotate row y=0 by 5
rotate column x=0 by 1
rect 4x1
rotate row y=2 by 15
rotate row y=0 by 15
rotate column x=10 by 1
rotate column x=5 by 1
rotate column x=0 by 1
rect 14x1
rotate column x=37 by 1
rotate column x=23 by 1
rotate column x=7 by 2
rotate row y=3 by 20
rotate row y=0 by 5
rotate column x=0 by 1
rect 4x1
rotate row y=3 by 5
rotate row y=2 by 2
rotate row y=1 by 4
rotate row y=0 by 4
rect 1x4
rotate column x=35 by 3
rotate column x=18 by 3
rotate column x=13 by 3
rotate row y=3 by 5
rotate row y=2 by 3
rotate row y=1 by 1
rotate row y=0 by 1
rect 1x5
rotate row y=4 by 20
rotate row y=3 by 10
rotate row y=2 by 13
rotate row y=0 by 10
rotate column x=5 by 1
rotate column x=3 by 3
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 9x1
rotate row y=4 by 10
rotate row y=3 by 10
rotate row y=1 by 10
rotate row y=0 by 10
rotate column x=7 by 2
rotate column x=5 by 1
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 9x1
rotate row y=4 by 20
rotate row y=3 by 12
rotate row y=1 by 15
rotate row y=0 by 10
rotate column x=8 by 2
rotate column x=7 by 1
rotate column x=6 by 2
rotate column x=5 by 1
rotate column x=3 by 1
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 9x1
rotate column x=46 by 2
rotate column x=43 by 2
rotate column x=24 by 2
rotate column x=14 by 3
rotate row y=5 by 15
rotate row y=4 by 10
rotate row y=3 by 3
rotate row y=2 by 37
rotate row y=1 by 10
rotate row y=0 by 5
rotate column x=0 by 3
rect 3x3
rotate row y=5 by 15
rotate row y=3 by 10
rotate row y=2 by 10
rotate row y=0 by 10
rotate column x=7 by 3
rotate column x=6 by 3
rotate column x=5 by 1
rotate column x=3 by 1
rotate column x=2 by 1
rotate column x=1 by 1
rotate column x=0 by 1
rect 9x1
rotate column x=19 by 1
rotate column x=10 by 3
rotate column x=5 by 4
rotate row y=5 by 5
rotate row y=4 by 5
rotate row y=3 by 40
rotate row y=2 by 35
rotate row y=1 by 15
rotate row y=0 by 30
rotate column x=48 by 4
rotate column x=47 by 3
rotate column x=46 by 3
rotate column x=45 by 1
rotate column x=43 by 1
rotate column x=42 by 5
rotate column x=41 by 5
rotate column x=40 by 1
rotate column x=33 by 2
rotate column x=32 by 3
rotate column x=31 by 2
rotate column x=28 by 1
rotate column x=27 by 5
rotate column x=26 by 5
rotate column x=25 by 1
rotate column x=23 by 5
rotate column x=22 by 5
rotate column x=21 by 5
rotate column x=18 by 5
rotate column x=17 by 5
rotate column x=16 by 5
rotate column x=13 by 5
rotate column x=12 by 5
rotate column x=11 by 5
rotate column x=3 by 1
rotate column x=2 by 5
rotate column x=1 by 5
rotate column x=0 by 1`

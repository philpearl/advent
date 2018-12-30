package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
)

func main() {

	part1("input.txt")
}

func part1(filename string) {
	clay, err := read(filename)
	if err != nil {
		log.Fatalln(err)
	}

	w := new(clay)
	current := point{x: 500, y: 0}

	for {
		changes := w.changes
		w.drop(current)
		if w.changes == changes {
			// No more changes on next drop of water
			break
		}
		// w.print()
		// fmt.Printf("%d changes \n", w.changes)
		// fmt.Println()
	}

	var count int
	for _, c := range w.arena {
		if c == '|' || c == '~' {
			count++
		}
	}
	fmt.Printf("%d wet cells\n", count-w.minY+1)

	count = 0
	for _, c := range w.arena {
		if c == '~' {
			count++
		}
	}
	fmt.Printf("%d retained\n", count)

}

type world struct {
	arena      []byte
	w          int
	minX, maxX int
	minY, maxY int
	changes    int
}

func new(clay []point) *world {
	minX, maxX, minY, maxY := extent(clay)
	_ = minY

	w := maxX - minX + 3
	arena := make([]byte, w*(maxY+1))

	for i := range arena {
		arena[i] = '.'
	}

	ww := &world{
		arena: arena,
		w:     w,
		minX:  minX,
		maxX:  maxX,
		minY:  minY,
		maxY:  maxY,
	}
	for _, c := range clay {
		ww.set(c, '#')
	}

	return ww
}

func (w *world) spread(dirn int, pos point) (wentDeeper bool) {

	for {
		below := pos
		below.y++
		if p := w.content(below); p == '#' || p == '~' {
			// current position is supported: move laterally if possible
			pos.x += dirn
			if p := w.content(pos); p == '#' {
				// Can't move laterally
				return wentDeeper
			}
			w.set(pos, '|')
		} else {
			wentDeeper = true
			pos = below
			w.set(pos, '|')
			w.drop(pos)
			return
		}
	}
}

func (w *world) drop(pos point) {
	for pos.y < w.maxY {
		below := pos
		below.y++
		if p := w.content(below); p == '#' || p == '~' {
			// current position is supported.
			deeper := w.spread(+1, pos)
			deeper2 := w.spread(-1, pos)
			if !deeper && !deeper2 {
				// In a container. Convert this level to ~, then drip in again from the top
				for p := pos; w.content(p) == '|'; p.x++ {
					w.set(p, '~')
				}
				l := pos
				l.x--
				for p := l; w.content(p) == '|'; p.x-- {
					w.set(p, '~')
				}
				return
			}
			break
		}
		w.set(below, '|')
		pos = below
	}
}

func (w *world) set(p point, c byte) {
	i := w.flatten(p)
	if w.arena[i] != c {
		w.changes++
		w.arena[i] = c
	}
}

func (w *world) content(p point) byte {
	return w.arena[w.flatten(p)]
}

func (w *world) flatten(p point) int {
	return (p.x - w.minX + 1) + w.w*p.y
}

func (w *world) print() {
	for i := 0; i < len(w.arena); i += w.w {
		fmt.Println(string(w.arena[i : i+w.w]))
	}
}

func extent(coords []point) (minX, maxX, minY, maxY int) {
	minX, minY = math.MaxInt64, math.MaxInt64
	for _, c := range coords {
		if c.y < minY {
			minY = c.y
		}
		if c.y > maxY {
			maxY = c.y
		}
		if c.x < minX {
			minX = c.x
		}
		if c.x > maxX {
			maxX = c.x
		}
	}
	return minX, maxX, minY, maxY
}

type point struct {
	x, y int
}

func read(filename string) ([]point, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var coords []point
	for {
		var z, start, end int
		var a1, a2 byte

		_, err := fmt.Fscanf(f, "%c=%d, %c=%d..%d", &a1, &z, &a2, &start, &end)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if a1 == 'x' {
			for y := start; y <= end; y++ {
				coords = append(coords, point{x: z, y: y})
			}
		} else {
			for x := start; x <= end; x++ {
				coords = append(coords, point{x: x, y: z})
			}
		}
	}

	return coords, nil
}

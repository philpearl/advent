package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	part1("input.txt")
	fmt.Println(time.Since(start))
}

func part1(filename string) {
	particles, err := loadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	var lastM = math.MaxInt64
	scale := 120
	count := 0
	for {
		tl, br := evolve(particles, scale)
		m := (br.x - tl.x + 1) * (br.y - tl.y + 1)
		if m > lastM {
			devolve(particles)
			print(particles, tl, br)
			fmt.Printf("count is %d\n", count)
			break
		}
		lastM = m
		count += scale
		if scale > 1 {
			scale--
		}
	}
}

func print(particles []particle, tl, br point) {
	arena, tl, br := fillArena(particles, tl, br)
	w := br.x - tl.x + 1
	h := br.y - tl.y + 1

	wr := bufio.NewWriter(os.Stdout)
	for y := 0; y < h; y++ {
		for _, x := range arena[y*w : (y+1)*w] {
			if x {
				wr.Write([]byte{'#'})
			} else {
				wr.Write([]byte{'.'})
			}
		}
		wr.Write([]byte{'\n'})
	}
	wr.Flush()
}

func fillArena(particles []particle, tl, br point) ([]bool, point, point) {
	w := br.x - tl.x + 1
	h := br.y - tl.y + 1
	arena := make([]bool, w*h)
	for _, p := range particles {
		x := p.position.x - tl.x
		y := p.position.y - tl.y
		arena[x+w*y] = true
	}
	return arena, tl, br
}

func evolve(particles []particle, scale int) (tl, br point) {
	tl.x = math.MaxInt64
	tl.y = math.MaxInt64
	br.x = math.MinInt64
	br.y = math.MinInt64

	for i := range particles {
		p := &particles[i]

		p.position.x += p.velocity.x * scale
		p.position.y += p.velocity.y * scale

		if p.position.x < tl.x {
			tl.x = p.position.x
		}
		if p.position.x > br.x {
			br.x = p.position.x
		}

		if p.position.y < tl.y {
			tl.y = p.position.y
		}
		if p.position.y > br.y {
			br.y = p.position.y
		}
	}

	return tl, br
}

func devolve(particles []particle) {
	for i := range particles {
		p := &particles[i]

		p.position.x -= p.velocity.x
		p.position.y -= p.velocity.y
	}
}

var lineRegex = regexp.MustCompile(`^position=<\s*(-?\d+),\s*(-?\d+)> velocity=<\s*(-?\d+),\s*(-?\d+)>$`)

type point struct {
	x, y int
}

type particle struct {
	position point
	velocity point
}

func loadFile(filename string) ([]particle, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	s := bufio.NewScanner(f)

	var particles []particle

	for s.Scan() {
		line := s.Text()
		matches := lineRegex.FindStringSubmatch(line)

		x, _ := strconv.Atoi(matches[1])
		y, _ := strconv.Atoi(matches[2])
		vx, _ := strconv.Atoi(matches[3])
		vy, _ := strconv.Atoi(matches[4])

		particles = append(particles, particle{
			position: point{
				x: x,
				y: y,
			},
			velocity: point{
				x: vx,
				y: vy,
			},
		})
	}

	return particles, s.Err()
}

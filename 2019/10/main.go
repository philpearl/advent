package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func main() {
	run(input)
}

type pos struct {
	x, y int
}
type polar struct {
	theta float64
	r     float64
	pos
}

func run(input string) {

	asteroids := make(map[pos]struct{}, len(input))
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		for x, c := range line {
			if c == '#' {
				asteroids[pos{x: x, y: y}] = struct{}{}
			}
		}
	}

	var maxCount int
	var bestAsteroid pos
	// From each asteroid we find the vector to every other asteroid. If that
	// vector passes through any integer locations we check those for other
	// asteroids. We count only asteroids where there is no other hit
	for source := range asteroids {
		var count int
		for target := range asteroids {
			if source == target {
				continue
			}

			vector := pos{
				x: target.x - source.x,
				y: target.y - source.y,
			}
			// So now we have a vector like (3, 15). This goes through (1, 5),
			// (2, 10) before hitting (3, 15). We basically want to find the
			// highest common factor between x & y.
			//
			x := vector.x
			if x < 0 {
				x = -x
			}
			y := vector.y
			if y < 0 {
				y = -y
			}

			f := hcf(x, y)
			if f == 1 {
				count++
				continue
			}
			var hit bool
			for i := 1; i < f; i++ {
				p := pos{
					x: source.x + (vector.x*i)/f,
					y: source.y + (vector.y*i)/f,
				}
				_, hit = asteroids[p]
				if hit {
					break
				}
			}
			if !hit {
				count++
			}
		}
		if count > maxCount {
			maxCount = count
			bestAsteroid = source
		}
	}
	fmt.Println(bestAsteroid, maxCount)

	// Starting from our best asteroid we vaporise asteroids!
	// Find all the asteroids visible from best
	delete(asteroids, bestAsteroid)
	for target := range asteroids {
		if target == bestAsteroid {
			continue
		}

		vector := pos{
			x: target.x - bestAsteroid.x,
			y: target.y - bestAsteroid.y,
		}
		// So now we have a vector like (3, 15). This goes through (1, 5),
		// (2, 10) before hitting (3, 15). We basically want to find the
		// highest common factor between x & y.
		//
		x := vector.x
		if x < 0 {
			x = -x
		}
		y := vector.y
		if y < 0 {
			y = -y
		}

		f := hcf(x, y)
		if f == 1 {
			continue
		}
		for i := 1; i < f; i++ {
			p := pos{
				x: bestAsteroid.x + (vector.x*i)/f,
				y: bestAsteroid.y + (vector.y*i)/f,
			}

			if _, hit := asteroids[p]; hit {
				delete(asteroids, target)
				break
			}
		}
	}
	// asteroids now has just the asteroids visible from best asteroid. Order them
	type angle struct {
		pos
		theta float64
	}

	// 0 -> pi/2
	fmt.Println(math.Atan(0.0/1.0), math.Atan(1.0), math.Atan(1.0/0.0000000001), math.Atan(-1.0), math.Atan(-0.000000001), math.Atan(-1.0/0.0000000001))

	angles := make([]angle, 0, len(asteroids))
	// We can take vectors from our asteroid to all other asteroids, work out the angle from the vertical and order them by angle.
	for asteroid := range asteroids {
		x := asteroid.x - bestAsteroid.x
		y := asteroid.y - bestAsteroid.y

		// Want up to be angle zero for straight up.
		var theta float64
		switch {
		case x >= 0 && y <= 0:
			theta = math.Atan(float64(x) / float64(-y))
		case x >= 0 && y >= 0:
			theta = math.Pi/4 + math.Atan(float64(y)/float64(x))
		case x <= 0 && y >= 0:
			theta = math.Pi/2 + math.Atan(float64(y)/float64(-x))
		case x <= 0 && y <= 0:
			theta = math.Pi*3/4 + math.Atan(float64(-y)/float64(-x))

		}
		angles = append(angles, angle{theta: theta, pos: asteroid})
	}
	sort.Slice(angles, func(i, j int) bool {
		return angles[i].theta < angles[j].theta
	})

	for i, ang := range angles {
		fmt.Println(i, ang)
	}

}

func hcf(a, b int) int {
	if a < b {
		a, b = b, a
	}
	if b == 0 {
		return a
	}

	return hcf(b, a%b)
}

var input = `#.....#...#.........###.#........#..
....#......###..#.#.###....#......##
......#..###.......#.#.#.#..#.......
......#......#.#....#.##....##.#.#.#
...###.#.#.......#..#...............
....##...#..#....##....#...#.#......
..##...#.###.....##....#.#..##.##...
..##....#.#......#.#...#.#...#.#....
.#.##..##......##..#...#.....##...##
.......##.....#.....##..#..#..#.....
..#..#...#......#..##...#.#...#...##
......##.##.#.#.###....#.#..#......#
#..#.#...#.....#...#...####.#..#...#
...##...##.#..#.....####.#....##....
.#....###.#...#....#..#......#......
.##.#.#...#....##......#.....##...##
.....#....###...#.....#....#........
...#...#....##..#.#......#.#.#......
.#..###............#.#..#...####.##.
.#.###..#.....#......#..###....##..#
#......#.#.#.#.#.#...#.#.#....##....
.#.....#.....#...##.#......#.#...#..
...##..###.........##.........#.....
..#.#..#.#...#.....#.....#...###.#..
.#..........#.......#....#..........
...##..#..#...#..#...#......####....
.#..#...##.##..##..###......#.......
.##.....#.......#..#...#..#.......#.
#.#.#..#..##..#..............#....##
..#....##......##.....#...#...##....
.##..##..#.#..#.................####
##.......#..#.#..##..#...#..........
#..##...#.##.#.#.........#..#..#....
.....#...#...#.#......#....#........
....#......###.#..#......##.....#..#
#..#...##.........#.....##.....#....`

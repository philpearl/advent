package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {

	part1("input.txt")
}

func part1(filename string) {
	world, err := read(filename)
	if err != nil {
		log.Fatalln(err)
	}

	world2 := make([][]byte, len(world))
	for i := range world2 {
		world2[i] = make([]byte, len(world[0]))
	}

	// we expect to see a cycle. We can look at the count of trees and lumber yards as a "hash" of the setup
	type counts struct {
		trees int
		lum   int
	}
	getCounts := func() counts {
		var counts counts
		for _, row := range world {
			for _, c := range row {
				switch c {
				case '|':
					counts.trees++
				case '#':
					counts.lum++
				}
			}
		}
		return counts
	}
	seen := make(map[counts]int)

	for m := 0; m < 1e9; m++ {
		var changes int
		for y := range world {
			for x, c := range world[y] {
				var trees, lum int
				for dx := -1; dx <= 1; dx++ {
					for dy := -1; dy <= 1; dy++ {
						if dx == 0 && dy == 0 {
							continue
						}
						if x+dx < 0 || x+dx >= len(world[0]) {
							continue
						}
						if y+dy < 0 || y+dy >= len(world) {
							continue
						}
						switch world[y+dy][x+dx] {
						case '#':
							lum++
						case '|':
							trees++
						}
					}
				}

				world2[y][x] = c
				switch c {
				case '.':
					// looking for 3 or more adj trees
					if trees >= 3 {
						world2[y][x] = '|'
						changes++
					}
				case '#':
					// looking for adj lumberyard and trees
					if trees == 0 || lum == 0 {
						world2[y][x] = '.'
						changes++
					}
				case '|':
					// looking for adj lumberyard * 3
					if lum >= 3 {
						world2[y][x] = '#'
						changes++
					}
				}
			}
		}
		world, world2 = world2, world
		if changes == 0 {
			break
		}

		// for _, row := range world {
		// 	fmt.Println(string(row))
		// }
		// fmt.Println()
		if m == 9 {
			counts := getCounts()
			fmt.Printf("Part 1: %d resources (%d trees, %d lum)\n", counts.trees*counts.lum, counts.trees, counts.lum)
		}

		// Bit of a hack - assume stability after a reasonable number of minutes. Our "hash" is extremely weak,
		// so we need to be in the stable section before looking
		if m > 1000 {
			// fmt.Printf("%d changes\n", changes)
			counts := getCounts()

			min, ok := seen[counts]
			if ok {
				fmt.Printf("cycle seen minute %d matches %d. Counts %#v\n", m, min, counts)

				// Calculate where we are in the cycle at 1e9
				offset := (1e9 - min) % (m - min)
				for counts, mm := range seen {
					if mm == min+offset-1 {
						fmt.Printf("%d resources (%d trees, %d lum)\n", counts.trees*counts.lum, counts.trees, counts.lum)
						break
					}
				}
				break
			}
			seen[counts] = m
		}
	}
}

func read(filename string) ([][]byte, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return bytes.Split(data, []byte{'\n'}), nil
}

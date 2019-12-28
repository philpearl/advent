package main

import "fmt"

func main() {
	run1([4]vec{
		{13, -13, -2},
		{16, 2, -15},
		{7, -18, -12},
		{-3, -8, -8},
	})

	run2([4]vec{
		{13, -13, -2},
		{16, 2, -15},
		{7, -18, -12},
		{-3, -8, -8},
	})
}

func run2(positions [4]vec) {
	// Find the period of each axis
	type state struct {
		p, v int32
	}
	states := make(map[[4]state]int)
	var periods [3]int
	for axis := 0; axis < 3; axis++ {
		for k := range states {
			delete(states, k)
		}

		var bodies [4]state
		for i, p := range positions {
			bodies[i].p = p[axis]
		}
		states[bodies] = 0

		for step := 1; ; step++ {
			for i := range bodies {
				ba := &bodies[i]
				for j, bb := range bodies {
					if i == j {
						continue
					}

					p := bb.p
					if ba.p < p {
						ba.v++
					} else if ba.p > p {
						ba.v--
					}
				}
			}
			// Update positions
			for i := range bodies {
				b := &bodies[i]
				b.p += b.v
			}

			if t, ok := states[bodies]; ok {
				fmt.Printf("axis %d back to step %d at step %d\n", axis, t, step)
				periods[axis] = step
				break
			}
			states[bodies] = step
		}
	}

	// Want lowest common multiple. So multiply then divide by GCD
	period := periods[0]
	for _, p := range periods[1:] {
		period = period * p / hcf(period, p)
	}
	fmt.Println(period)
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

func run1(positions [4]vec) {
	var bodies [4]body

	for i, p := range positions {
		bodies[i].p = p
	}

	// Strangely do this one axis at a time just to check they really are independent
	for axis := 0; axis < 3; axis++ {
		for step := 0; step < 1000; step++ {
			// Update velocity by applying gravity
			for i := range bodies {
				ba := &bodies[i]
				for j, b := range bodies {
					if i == j {
						continue
					}

					p := b.p[axis]
					if ba.p[axis] < p {
						ba.v[axis]++
					} else if ba.p[axis] > p {
						ba.v[axis]--
					}
				}
			}
			// Update positions
			for i := range bodies {
				b := &bodies[i]
				b.p[axis] += b.v[axis]
			}
		}
	}

	printEnergy(bodies)
}

func printEnergy(bodies [4]body) {
	// Calculate energies
	var energy int
	for _, b := range bodies {
		var ep int
		for _, p := range b.p {
			if p < 0 {
				ep -= int(p)
			} else {
				ep += int(p)
			}
		}
		var ev int
		for _, v := range b.v {
			if v < 0 {
				ev -= int(v)
			} else {
				ev += int(v)
			}
		}
		energy += ep * ev
	}
	fmt.Println(energy)
}

type body struct {
	p vec
	v vec
}

type vec [3]int32

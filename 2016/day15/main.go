package main

import (
	"fmt"
)

type disc struct {
	mod      int
	position int
}

type stack []disc

var discs = stack{
	{mod: 13, position: 11},
	{mod: 5, position: 0},
	{mod: 17, position: 11},
	{mod: 3, position: 0},
	{mod: 7, position: 2},
	{mod: 19, position: 17},
	{mod: 11, position: 0},
}

func main() {

	for offset := 0; ; offset++ {
		good := true
		for i := range discs {
			d := discs[i]
			slot := (d.position + offset + i + 1) % d.mod
			if slot != 0 {
				// fmt.Printf("offset %d disc %d is at slot %d\n", offset, i, slot)
				good = false
				break
			}
		}
		if good {
			fmt.Printf("Press button at time %d\n", offset)
			break
		}
	}

}

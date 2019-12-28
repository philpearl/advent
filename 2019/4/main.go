package main

import (
	"fmt"
	"os"
)

func main() {
	if err := run(input, number{4, 4, 4, 4, 4, 3}); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func run(input rang, start number) error {

	n := start
	var count, count2 int

	for {

		n[5]++
		if n[5] == 10 {
			n[4]++
			if n[4] == 10 {
				n[3]++
				if n[3] == 10 {
					n[2]++
					if n[2] == 10 {
						n[1]++
						if n[1] == 10 {
							n[0]++
							n[1] = n[0]
						}
						n[2] = n[1]
					}
					n[3] = n[2]
				}
				n[4] = n[3]
			}
			n[5] = n[4]
		}

		val := toInt(n)
		if val > input.max {
			break
		}

		// check for a double-digit
		if hasDouble(n) {
			count++
			if hasTrueDouble(n) {
				count2++
			}
		}
	}

	fmt.Println("result", count, count2)

	return nil
}

func hasDouble(n number) bool {
	for i := range n[1:] {
		if n[i] == n[i+1] {
			return true
		}
	}
	return false
}

func hasTrueDouble(n number) bool {
	for i := range n[1:] {
		if n[i] == n[i+1] {
			if (i == 0 || n[i-1] != n[i]) && (i == len(n)-2 || n[i+2] != n[i+1]) {
				return true
			}
		}
	}
	return false
}

func toInt(n number) int {
	var v int
	for _, d := range n {
		v = v*10 + int(d)
	}
	return v
}

type number [6]byte

type rang struct {
	min, max int
}

var input = rang{402328, 864247}

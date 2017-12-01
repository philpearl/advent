package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

var salt = "cuanljph"

func hash(val int) string {
	sum := md5.Sum([]byte(salt + strconv.Itoa(val)))
	h := hex.EncodeToString(sum[:])

	for i := 0; i < 2016; i++ {
		sum := md5.Sum([]byte(h))
		h = hex.EncodeToString(sum[:])
	}
	return h
}

func findRun(input string) (repeated rune) {
	var l, ll rune
	for _, c := range input {
		if c == ll && ll == l {
			return c
		}
		ll = l
		l = c
	}
	return 0
}

type hashBuffer struct {
	hashes []string
	index  int
}

func (hb *hashBuffer) next() string {
	h := hb.peek(0)
	hb.index++
	return h
}

func (hb *hashBuffer) peek(offset int) string {
	for len(hb.hashes) <= hb.index+offset {
		hb.hashes = append(hb.hashes, hash(len(hb.hashes)))
	}
	return hb.hashes[hb.index+offset]
}

func main() {
	hb := hashBuffer{}

	for keys := 0; keys < 64; {
		h := hb.next()
		if r := findRun(h); r != 0 {
			var ok bool
			// fmt.Printf("index %d found run of %c for hash %s\n", hb.index-1, r, h)
			lookFor := strings.Repeat(string(r), 5)
			for offset := 0; offset < 1000; offset++ {
				candidate := hb.peek(offset)
				if strings.Contains(candidate, lookFor) {
					ok = true
					break
				}
			}
			if ok {
				fmt.Printf("hash %s at index %d is good\n", h, hb.index-1)
				keys++
			}
		}
	}

}

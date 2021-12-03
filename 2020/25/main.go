package main

import (
	"fmt"
	"sync"
)

func main() {
	run(cardPK, doorPK)
}

func run(cardPK, doorPK int) {
	var wg sync.WaitGroup

	var cardLS int
	wg.Add(1)
	go func() {
		defer wg.Done()
		cardLS = findLS(7, cardPK)
		fmt.Println(cardLS)
	}()

	var doorLS int
	wg.Add(1)
	go func() {
		defer wg.Done()
		doorLS = findLS(7, doorPK)
		fmt.Println(doorLS)
	}()

	wg.Wait()
	fmt.Println(cardLS, doorLS)
	// These two should be the same!
	fmt.Println(transform(doorPK, cardLS))
	fmt.Println(transform(cardPK, doorLS))

}

var cardPK = 19241437
var doorPK = 17346587

func findLS(sn, target int) int {
	v := 1
	i := 0
	for v != target {
		v = (v * sn) % 20201227
		i++
	}
	return i
}

func transform(sn int, loopSize int) int {
	v := 1
	for i := 0; i < loopSize; i++ {
		v = (v * sn) % 20201227
	}
	return v
}

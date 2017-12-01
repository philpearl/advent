package main

import (
	"fmt"
)

var outputLen int = 35651584

var input = []byte("10010000000110000")

func reverseAndInvert(arena []byte, inputLen, outputOffset int) int {
	for i := inputLen - 1; i >= 0; i-- {
		v := arena[i]
		var o byte
		if v == '0' {
			o = '1'
		} else {
			o = '0'
		}
		arena[outputOffset] = o
		outputOffset++

		if outputOffset == outputLen {
			break
		}
	}
	return outputOffset
}

func checksum(input []byte) []byte {
	for {
		ll := len(input)
		if ll&1 == 1 {
			return input
		}
		for i := 0; i < ll/2; i++ {
			a, b := input[2*i], input[2*i+1]
			if a == b {
				input[i] = '1'
			} else {
				input[i] = '0'
			}
		}
		input = input[:ll/2]
	}
}

func main() {

	arena := make([]byte, outputLen)
	copy(arena, input)

	inputLen := len(input)
	outputOffset := inputLen
	for outputOffset < outputLen {

		arena[outputOffset] = '0'
		outputOffset++
		outputOffset = reverseAndInvert(arena, inputLen, outputOffset)
		inputLen = outputOffset
	}

	fmt.Println(string(arena))
	fmt.Println(string(checksum(arena)))
}

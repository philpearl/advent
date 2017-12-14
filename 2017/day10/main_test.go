package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {

	tests := []struct {
		listLen int
		lengths []uint8
		exp     int
	}{
		{
			listLen: 5,
			lengths: []uint8{3, 4, 1, 5},
			exp:     12,
		},
	}

	for i, test := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			assert.Equal(t, test.exp, part1(test.listLen, test.lengths))
		})
	}
}

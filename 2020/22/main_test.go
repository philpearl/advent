package main

import (
	"testing"
)

func TestStateToKey(t *testing.T) {
	allocs := testing.AllocsPerRun(10, func() {
		stateToKey([]byte{1, 2, 3}, []byte{4, 5, 6})
	})
	if allocs > 0 {
		t.Fatalf("key is %f allocs", allocs)
	}
}

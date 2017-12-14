package main

import "testing"

func BenchmarkThing(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		part1(input)
	}
}

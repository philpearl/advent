package main

import "testing"

func BenchmarkDance(b *testing.B) {
	b.ReportAllocs()
	runDance(input, b.N)
}

package day26

import (
	"fmt"
	"testing"
)

// L1 64KB per core
// L2 256KB per core
// L3 6MB per core
//
// Cache lines are 64 bytes - 8 8 byte int64s, 16 4 byte int32s
//
// With 4 byte ints, a cache line is 16 ints, and we have 1024 of these
//
// 64KB is 16*1024 int32s. But this is split between instructions and data
//
// So if we have 8*1024 int32s we'd expect them to all fit in the cache, and the
// access pattern shouldn't really matter. But we find it slows down at stride=32
//
// At stride = 32 we access
// - 0
// - 32*4=128
// - 256

func BenchmarkTest(b *testing.B) {
	NUM := 64 * 1024
	numbers := make([]int32, NUM)
	for i := range numbers {
		numbers[i] = int32(i)
	}
	superTotal := 0

	for stride := 1; stride < NUM; stride *= 2 {
		b.Run(fmt.Sprintf("stride_%d", stride), func(b *testing.B) {
			b.SetBytes(int64(NUM))
			var total int32
			for n := 0; n < b.N; n++ {
				for offset := 0; offset < stride; offset++ {
					for i := offset; i < NUM; i += stride {
						total += numbers[i]
					}
				}
			}
			superTotal += int(total)
		})
	}
	b.Logf("supertotal is %d", superTotal)
}

func BenchmarkTest_JustAdd(b *testing.B) {
	b.SetBytes(int64(1))
	var total int
	for n := 0; n < b.N; n++ {
		total += int(byte(n))
	}
	b.Logf("total is %d", total)
}

func BenchmarkTest_Sequential(b *testing.B) {
	NUM := 64 * 1024
	numbers := make([]int8, NUM)
	for i := range numbers {
		numbers[i] = int8(i)
	}

	b.ResetTimer()
	b.SetBytes(int64(NUM))
	var total int
	for n := 0; n < b.N; n++ {
		for i := 0; i < NUM; i++ {
			a := int(numbers[i])
			total += a
		}
	}

	b.Logf("total is %d", total)
}

func BenchmarkTest_ExceedCacheline(b *testing.B) {
	NUM := 64 * 1024
	numbers := make([]int8, NUM)
	for i := range numbers {
		numbers[i] = int8(i)
	}

	b.ResetTimer()
	b.SetBytes(int64(NUM))
	var total int
	for n := 0; n < b.N; n++ {
		for j := 0; j < 64; j++ {
			for i := j; i < NUM; i += 64 {
				total += int(numbers[i])
			}
		}
	}
	b.Logf("total is %d", total)
}

func BenchmarkTest_JustAdd2(b *testing.B) {
	NUM := 64 * 1024

	b.ResetTimer()
	b.SetBytes(int64(NUM))
	var total int
	for n := 0; n < b.N; n++ {
		for j := 0; j < 64; j++ {
			for i := j; i < NUM; i += 64 {
				total += int(uint8(i))
			}
		}
	}
	b.Logf("total is %d", total)
}

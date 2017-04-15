package fastrand

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"testing"
)

// BenchSink prevents the compiler from optimizing away benchmark loops.
var BenchSink uint32

var parallelisms = []int{1,2,4,8}

func BenchmarkUint32n(b *testing.B) {
	for _, parallelism := range parallelisms {
		b.Run(fmt.Sprintf("%d", parallelism), func(b *testing.B) {
			benchmarkUint32n(b, parallelism)
		})
	}
}

func benchmarkUint32n(b *testing.B, parallelism int) {
	b.SetParallelism(parallelism)
	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += Uint32n(1e6)
		}
		atomic.AddUint32(&BenchSink, s)
	})
}

func BenchmarkRNGUint32n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		var r RNG
		s := uint32(0)
		for pb.Next() {
			s += r.Uint32n(1e6)
		}
		atomic.AddUint32(&BenchSink, s)
	})
}

func BenchmarkMathRandInt31n(b *testing.B) {
	for _, parallelism := range parallelisms {
		b.Run(fmt.Sprintf("%d", parallelism), func(b *testing.B) {
			benchmarkRandInt32n(b, parallelism)
		})
	}
}

func benchmarkRandInt32n(b *testing.B, parallelism int) {
	b.SetParallelism(parallelism)
	b.RunParallel(func(pb *testing.PB) {
		s := uint32(0)
		for pb.Next() {
			s += uint32(rand.Int31n(1e6))
		}
		atomic.AddUint32(&BenchSink, s)
	})
}

func BenchmarkMathRandRNGInt31n(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		r := rand.New(rand.NewSource(42))
		s := uint32(0)
		for pb.Next() {
			s += uint32(r.Int31n(1e6))
		}
		atomic.AddUint32(&BenchSink, s)
	})
}

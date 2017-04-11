// Package fastrand implements fast random number generators.
package fastrand

import (
	cryptorand "crypto/rand"
	"fmt"
	"sync"
)

// Uint32 returns pseudorandom uint32.
//
// It is safe calling this function from concurrent goroutines.
func Uint32() uint32 {
	r := getRNG()
	x := r.Uint32()
	putRNG(r)
	return x
}

// Uint32N returns pseudorandom uint32 in the range [0..maxN).
//
// It is safe calling this function from concurrent goroutines.
func Uint32N(maxN uint32) uint32 {
	x := Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

func getRNG() *RNG {
	v := rngPool.Get()
	if v == nil {
		v = &RNG{
			x: getRandomUint32(),
		}
	}
	return v.(*RNG)
}

func putRNG(r *RNG) {
	rngPool.Put(r)
}

var rngPool sync.Pool

// RNG is a pseudorandom number generator.
//
// It is unsafe to call RNG methods from concurrent goroutines.
type RNG struct {
	x uint32
}

// Uint32 returns pseudorandom uint32.
//
// It is unsafe to call this method from concurrent goroutines.
func (r *RNG) Uint32() uint32 {
	if r.x == 0 {
		r.x = getRandomUint32()
	}

	// See https://en.wikipedia.org/wiki/Xorshift
	x := r.x
	x ^= x << 13
	x ^= x >> 17
	x ^= x << 5
	r.x = x
	return x
}

// Uint32N returns pseudorandom uint32 in the range [0..maxN).
//
// It is unsafe to call this method from concurrent goroutines.
func (r *RNG) Uint32N(maxN uint32) uint32 {
	x := r.Uint32()
	// See http://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
	return uint32((uint64(x) * uint64(maxN)) >> 32)
}

func getRandomUint32() uint32 {
	var buf [4]byte
	_, err := cryptorand.Read(buf[:])
	if err != nil {
		panic(fmt.Sprintf("BUG: cannot read random number: %s", err))
	}
	return uint32(buf[3]) | (uint32(buf[2]) << 8) | (uint32(buf[1]) << 16) | (uint32(buf[0]) << 24)
}

// Package egc implements the Elias gamma coding as specified by
// https://en.wikipedia.org/wiki/Elias_gamma_coding.

package egc

import (
	"math/bits"

	"github.com/RussellLuo/bite/bitmap"
)

func Encode(x uint) *bitmap.Bitmap {
	if x == 0 {
		panic("x is zero")
	}

	l := bits.Len(x)
	N := l - 1

	size := 2*N + 1
	b := bitmap.New(size)

	// Write out the lowest N+1 bits, i.e. the binary form of x.
	b.SetUint64(uint64(x))

	// The remaining highest N bits default to zero.

	return b
}

func Decode(b *bitmap.Bitmap) (uint, bool) {
	if b.Size() == 0 {
		return 0, false
	}

	x, ok := b.Uint64()
	return uint(x), ok
}

package bitmap

import (
	"math/big"
	"math/bits"
	"strings"
)

const (
 	_WS = bits.UintSize // word size in bits
)

type Bitmap struct {
	x *big.Int
	size int
}

func New(size int) *Bitmap {
	return &Bitmap{
		x: new(big.Int),
		size: size,
	}
}

func (b *Bitmap) Size() int {
	return b.size
}

// SetBit sets b's i'th bit to v (0 or 1). If b is not 0 or 1,
// SetBit will panic.
func (b *Bitmap) SetBit(i int, v uint) {
	b.x.SetBit(b.x, i, v)
}

// Bit returns the value of the i'th bit of b. The bit index i must be >= 0.
func (b *Bitmap) Bit(i int) uint {
	return b.x.Bit(i)
}

// SetUint64 sets b to x.
func (b *Bitmap) SetUint64(x uint64) {
	b.x.SetUint64(x)
}

// Uint64 returns the uint64 representation of b.
// If b cannot be represented in a uint64, Uint64 will return (0, false).
func (b *Bitmap) Uint64() (uint64, bool) {
	if !b.x.IsUint64() {
		return 0, false
	}
	return b.x.Uint64(), true
}

// SetString sets b to the value of s, interpreted in the given base,
// and returns a boolean indicating success.
//
// A prefix of ``0b'' or ``0B'' selects base 2, ``0'', ``0o'' or ``0O''
// selects base 8, and ``0x'' or ``0X'' selects base 16. Otherwise, the
// selected base is 10 and no prefix is accepted.
func (b *Bitmap) SetString(s string) bool {
	_, ok := b.x.SetString(s, 0)
	return ok
}

// String returns the string representation of b in the given base.
// Base must be between 2 and 62, inclusive.
func (b *Bitmap) String(base int) string {
	nonzero := b.x.Text(base)
	return strings.Repeat("0", b.size-len(nonzero)) + nonzero
}

// SetBytes interprets buf as the bytes of a big-endian unsigned
// integer, and sets b to that value.
func (b *Bitmap) SetBytes(buf []byte) {
	b.x.SetBytes(buf)
}

// Bytes returns the absolute value of b as a big-endian byte slice.
func (b *Bitmap) Bytes() []byte {
	return b.x.Bytes()
}

// Range returns bits [s, e) as another bitmap.
func (b *Bitmap) Range(s, e int) *Bitmap {
	if e <= s {
		panic("end is not greater than start")
	}

	size := e - s
	out := New(size)

	words := b.x.Bits()

	i := s / _WS
	if i >= len(words) {
		return out
	}

	shift := uint(s % _WS)
	word := words[i] >> shift

	for j := 0; j < size; j++ {
		out.SetBit(j, uint(word & 1))

		if shift >= _WS - 1 {
			// All bits are consumed, move to the next word.
			i++
			if i >= len(words) {
				break
			}

			shift = 0
			word = words[i]
		} else {
			// Right shift the word.
			shift++
			word >>= 1
		}
	}

	return out
}

// Package rle implements Run-length encoding as specified by
// https://en.wikipedia.org/wiki/Run-length_encoding, but in a bit-based way.

package rle

import (
	"strconv"
	"strings"

	"github.com/RussellLuo/bite/bitmap"
	"github.com/RussellLuo/bite/egc"
)

func Encode(b *bitmap.Bitmap) *bitmap.Bitmap {
	if b.Size() == 0 {
		return bitmap.New(0)
	}

	var binStr strings.Builder

	prev := b.Bit(b.Size() - 1)
	run := uint(1)

	// Record the highest bit.
	binStr.WriteString(strconv.Itoa(int(prev)))

	// Iterate in a big-endian order.
	for i := b.Size() - 2; i > -1; i-- {
		curr := b.Bit(i)

		if curr != prev {
			binStr.WriteString(egc.Encode(run).String(2))

			prev = curr
			run = 0
		}

		run++
	}

	// Encode the ending run.
	if run > 0 {
		binStr.WriteString(egc.Encode(run).String(2))
	}

	str := binStr.String()

	if len(str) > b.Size() {
		// The bit size is expanded after compression :(
		//
		// Use the original uncompressed bits instead, with an extra leading
		// zero to indicate the data is in uncompressed format.
		out := bitmap.New(b.Size() + 1)
		out.SetString("0b0" + b.String(2))
		return out
	}

	// The data is effectively compressed. Use the compressed bits, with an
	// extra leading one to indicate the data is in compressed format.
	out := bitmap.New(len(str) + 1)
	out.SetString("0b1" + str)
	return out
}

func Decode(b *bitmap.Bitmap) *bitmap.Bitmap {
	if b.Size() == 0 {
		return bitmap.New(0)
	}

	if b.Bit(b.Size()-1) == 0 {
		// The data is in uncompressed format.
		out := bitmap.New(b.Size() - 1)
		out.SetString("0b" + b.String(2)[1:])
		return out
	}

	// Decode the data since it's in compressed format.

	var binStr strings.Builder

	bit := b.Bit(b.Size() - 2)
	bitStr := strconv.Itoa(int(bit))

	N := 0 // the number of zero bit

	for i := b.Size() - 3; i > -1; {
		if b.Bit(i) == 1 {
			runB := b.Range(i-N, i+N+1) // EGC (elias gamma-coded) integer has 2N+1 bits
			run, _ := egc.Decode(runB)

			for j := 0; j < int(run); j++ {
				binStr.WriteString(bitStr)
			}

			// Invert the bit.
			bit = (bit + 1) & 1
			bitStr = strconv.Itoa(int(bit))

			// Move to the next EGC integer.
			i -= N + 1
			N = 0
		} else {
			i--
			N++
		}
	}

	str := binStr.String()
	out := bitmap.New(len(str))
	out.SetString("0b" + str)
	return out
}

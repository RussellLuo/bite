package rle

import (
	"testing"

	"github.com/RussellLuo/bite/bitmap"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		name    string
		inBin   string
		wantBin string
	}{
		{
			name:    "unexpectedly expanded",
			inBin:   "0000000011000011",
			wantBin: "00000000011000011",
		},
		{
			name:    "successfully compressed short data",
			inBin:   "0000000000000011",
			wantBin: "100001110010",
		},
		{
			name:    "successfully compressed medium data",
			inBin:   "00000000001111111000000010000000",
			wantBin: "1000010100011100111100111",
		},
		{
			name:    "successfully compressed long data",
			inBin:   "1111111111111111000000000000001111111110000111110000000000000011",
			wantBin: "110000100000001110000100100100001010001110010",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := bitmap.New(len(c.inBin))
			b.SetString("0b" + c.inBin)

			out := Encode(b)
			got := out.String(2)

			if got != c.wantBin {
				t.Errorf("Got (%#v) != Want (%#v)", got, c.wantBin)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		name    string
		inBin   string
		wantBin string
	}{
		{
			name:    "data in uncompressed format",
			inBin:   "00000000011000011",
			wantBin: "0000000011000011",
		},
		{
			name:    "short data in compressed format",
			inBin:   "100001110010",
			wantBin: "0000000000000011",
		},
		{
			name:    "medium data in compressed format",
			inBin:   "1000010100011100111100111",
			wantBin: "00000000001111111000000010000000",
		},
		{
			name:    "long data in compressed format",
			inBin:   "110000100000001110000100100100001010001110010",
			wantBin: "1111111111111111000000000000001111111110000111110000000000000011",
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			b := bitmap.New(len(c.inBin))
			b.SetString("0b" + c.inBin)

			out := Decode(b)
			got := out.String(2)

			if got != c.wantBin {
				t.Errorf("Got (%#v) != Want (%#v)", got, c.wantBin)
			}
		})
	}
}

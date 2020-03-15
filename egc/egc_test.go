package egc

import (
	"testing"

	"github.com/RussellLuo/bite/bitmap"
)

func TestEncode(t *testing.T) {
	cases := []struct {
		x   uint
		bin string
	}{
		{
			x:   1,
			bin: "1",
		},
		{
			x:   2,
			bin: "010",
		},
		{
			x:   3,
			bin: "011",
		},
		{
			x:   4,
			bin: "00100",
		},
		{
			x:   5,
			bin: "00101",
		},
		{
			x:   6,
			bin: "00110",
		},
		{
			x:   7,
			bin: "00111",
		},
		{
			x:   8,
			bin: "0001000",
		},
		{
			x:   9,
			bin: "0001001",
		},
		{
			x:   10,
			bin: "0001010",
		},
		{
			x:   11,
			bin: "0001011",
		},
		{
			x:   12,
			bin: "0001100",
		},
		{
			x:   13,
			bin: "0001101",
		},
		{
			x:   14,
			bin: "0001110",
		},
		{
			x:   15,
			bin: "0001111",
		},
		{
			x:   16,
			bin: "000010000",
		},
		{
			x:   17,
			bin: "000010001",
		},
	}

	for _, c := range cases {
		b := Encode(c.x)
		got := b.String(2)
		if got != c.bin {
			t.Errorf("Got (%#v) != Want (%#v)", got, c.bin)
		}
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		bin string
		x   uint
	}{
		{
			bin: "1",
			x:   1,
		},
		{
			bin: "010",
			x:   2,
		},
		{
			bin: "011",
			x:   3,
		},
		{
			bin: "00100",
			x:   4,
		},
		{
			bin: "00101",
			x:   5,
		},
		{
			bin: "00110",
			x:   6,
		},
		{
			bin: "00111",
			x:   7,
		},
		{
			bin: "0001000",
			x:   8,
		},
		{
			bin: "0001001",
			x:   9,
		},
		{
			bin: "0001010",
			x:   10,
		},
		{
			bin: "0001011",
			x:   11,
		},
		{
			bin: "0001100",
			x:   12,
		},
		{
			bin: "0001101",
			x:   13,
		},
		{
			bin: "0001110",
			x:   14,
		},
		{
			bin: "0001111",
			x:   15,
		},
		{
			bin: "000010000",
			x:   16,
		},
		{
			bin: "000010001",
			x:   17,
		},
	}

	for _, c := range cases {
		b := bitmap.New(len(c.bin))
		b.SetString("0b" + c.bin)

		got, ok := Decode(b)
		if !ok {
			t.Error("Decode fails")
		}

		if got != c.x {
			t.Errorf("Got (%#v) != Want (%#v)", got, c.x)
		}
	}
}

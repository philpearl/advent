package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestLineToInt(t *testing.T) {
	// the bits are reversed
	if i := lineToInt([]byte("..##.#..#.")); i != 0b0100101100 {
		t.Fatalf("%b not as expected", i)
	}
}

func TestReverse(t *testing.T) {
	if i := reverseEdge(0b0100101100); i != 0b0011010010 {
		t.Fatalf("%b not as expected", i)
	}
}

var allow = cmp.AllowUnexported(tile{}, square{})

func TestParseTile(t *testing.T) {
	ts := `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###`

	tt := parseTile(ts)

	if diff := cmp.Diff(tile{
		n: 2311,
		edges: [4]int16{
			0b0100101100,
			0b1001101000,
			0b1110011100,
			0b0100111110,
		},
		body: square{
			w:    10,
			data: []byte("..##.#..#.##..#.....#...##..#.####.#...###.##.###.##...#.###.#.#.#..##..#....#..###...#.#...###..###"),
		},
	}, tt, allow); diff != "" {
		t.Fatalf("tile not as expected. %s", diff)
	}
}

func TestFlipBody(t *testing.T) {
	tt := tile{
		flipped: true,
		body: square{
			data: []byte("ABCDEFGHI"),
			w:    3,
		},
	}
	tt.flipBody()
	if diff := cmp.Diff(tile{
		body: square{
			data: []byte("CBAFEDIHG"),
			w:    3,
		},
	}, tt, allow); diff != "" {
		t.Fatalf("not as expected: %s", diff)
	}
}

func TestRotateBody(t *testing.T) {
	tests := []struct {
		tile tile
		exp  tile
	}{
		{
			tile: tile{
				angle: 90,
				body: square{
					data: []byte("ABCDEFGHI"),
					w:    3,
				},
			},
			exp: tile{
				body: square{
					data: []byte("GDAHEBIFC"),
					w:    3,
				},
			},
		},
		{
			tile: tile{
				angle: 180,
				body: square{
					data: []byte("ABCDEFGHI"),
					w:    3,
				},
			},
			exp: tile{
				body: square{
					data: []byte("IHGFEDCBA"),
					w:    3,
				},
			},
		},
		{
			tile: tile{
				angle: 270,
				body: square{
					data: []byte("ABCDEFGHI"),
					w:    3,
				},
			},
			exp: tile{
				body: square{
					data: []byte("CFIBEHADG"),
					w:    3,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			test.tile.rotateBody()
			if diff := cmp.Diff(test.exp, test.tile, allow); diff != "" {
				t.Fatalf("not as expected: %s", diff)
			}
		})
	}
}

func TestJoin(t *testing.T) {
	sq1 := square{
		data: []byte("123456789"),
		w:    3,
	}
	sq2 := square{
		data: []byte("987654321"),
		w:    3,
	}
	sq3 := square{
		data: []byte("ABCDEFGHI"),
		w:    3,
	}
	sq4 := square{
		data: []byte("IHGFEDCBA"),
		w:    3,
	}
	sq := join([]square{sq1, sq2, sq3, sq4})

	if diff := cmp.Diff(square{
		w:    6,
		data: []byte("123987456654789321ABCIHGDEFFEDGHICBA"),
	}, sq, allow); diff != "" {
		t.Fatalf("not as expected. %s", diff)
	}
}

package core

import (
	"sort"
	"testing"
)

// moveSet renders moves as a stable sorted list of "from->to" strings for
// table comparison without ordering assumptions.
func moveSet(ms []Move) []string {
	out := make([]string, len(ms))
	for i, m := range ms {
		out[i] = squareName(m.From) + "->" + squareName(m.To)
	}
	sort.Strings(out)
	return out
}

func squareName(s Square) string {
	return string(rune('a'+s.File)) + string(rune('1'+s.Rank))
}

func TestPseudoLegalMoves_King(t *testing.T) {
	cases := []struct {
		name  string
		setup func() (GameState, Square) // returns state and the King's square
		want  []string                    // expected "from->to" set
	}{
		{
			name: "open board centre — 8 targets",
			setup: func() (GameState, Square) {
				var b Board
				b[3][4] = &Piece{Kind: King, Color: White} // e4
				return GameState{Board: b, ActiveColor: White}, sq(4, 3)
			},
			want: []string{
				"e4->d3", "e4->d4", "e4->d5",
				"e4->e3", "e4->e5",
				"e4->f3", "e4->f4", "e4->f5",
			},
		},
		{
			name: "corner a1 — 3 targets",
			setup: func() (GameState, Square) {
				var b Board
				b[0][0] = &Piece{Kind: King, Color: White}
				return GameState{Board: b, ActiveColor: White}, sq(0, 0)
			},
			want: []string{"a1->a2", "a1->b1", "a1->b2"},
		},
		{
			name: "corner h8 — 3 targets",
			setup: func() (GameState, Square) {
				var b Board
				b[7][7] = &Piece{Kind: King, Color: Black}
				return GameState{Board: b, ActiveColor: Black}, sq(7, 7)
			},
			want: []string{"h8->g7", "h8->g8", "h8->h7"},
		},
		{
			name: "blocked by own piece — same-color blocker excluded",
			setup: func() (GameState, Square) {
				var b Board
				b[3][4] = &Piece{Kind: King, Color: White}
				b[4][4] = &Piece{Kind: Pawn, Color: White} // own pawn on e5 blocks
				b[3][5] = &Piece{Kind: Pawn, Color: White} // own pawn on f4 blocks
				return GameState{Board: b, ActiveColor: White}, sq(4, 3)
			},
			want: []string{
				"e4->d3", "e4->d4", "e4->d5",
				"e4->e3",
				"e4->f3", "e4->f5",
			},
		},
		{
			name: "captures enemy piece — opposing-colour target allowed",
			setup: func() (GameState, Square) {
				var b Board
				b[3][4] = &Piece{Kind: King, Color: White}
				b[4][4] = &Piece{Kind: Pawn, Color: Black} // enemy pawn on e5 capturable
				return GameState{Board: b, ActiveColor: White}, sq(4, 3)
			},
			want: []string{
				"e4->d3", "e4->d4", "e4->d5",
				"e4->e3", "e4->e5",
				"e4->f3", "e4->f4", "e4->f5",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			state, from := tc.setup()
			got := moveSet(PseudoLegalMoves(state, from))
			if !equalStrings(got, tc.want) {
				t.Errorf("PseudoLegalMoves =\n  %v\nwant\n  %v", got, tc.want)
			}
		})
	}
}

func TestPseudoLegalMoves_Queen(t *testing.T) {
	cases := []struct {
		name  string
		setup func() (GameState, Square)
		want  []string
	}{
		{
			name: "open board centre — 27 targets",
			setup: func() (GameState, Square) {
				var b Board
				b[3][4] = &Piece{Kind: Queen, Color: White} // e4
				return GameState{Board: b, ActiveColor: White}, sq(4, 3)
			},
			want: []string{
				"e4->a4", "e4->a8",
				"e4->b1", "e4->b4", "e4->b7",
				"e4->c2", "e4->c4", "e4->c6",
				"e4->d3", "e4->d4", "e4->d5",
				"e4->e1", "e4->e2", "e4->e3", "e4->e5", "e4->e6", "e4->e7", "e4->e8",
				"e4->f3", "e4->f4", "e4->f5",
				"e4->g2", "e4->g4", "e4->g6",
				"e4->h1", "e4->h4", "e4->h7",
			},
		},
		{
			name: "blocked by own piece on west ray",
			setup: func() (GameState, Square) {
				var b Board
				b[3][4] = &Piece{Kind: Queen, Color: White} // e4
				b[3][2] = &Piece{Kind: Rook, Color: White}  // c4 — own piece blocks west ray
				return GameState{Board: b, ActiveColor: White}, sq(4, 3)
			},
			// d4 reachable, c4 blocked (own piece), a4 and b4 unreachable
			want: []string{
				"e4->a8",
				"e4->b1", "e4->b7",
				"e4->c2", "e4->c6",
				"e4->d3", "e4->d4", "e4->d5",
				"e4->e1", "e4->e2", "e4->e3", "e4->e5", "e4->e6", "e4->e7", "e4->e8",
				"e4->f3", "e4->f4", "e4->f5",
				"e4->g2", "e4->g4", "e4->g6",
				"e4->h1", "e4->h4", "e4->h7",
			},
		},
		{
			name: "captures enemy piece on west ray",
			setup: func() (GameState, Square) {
				var b Board
				b[3][4] = &Piece{Kind: Queen, Color: White} // e4
				b[3][2] = &Piece{Kind: Rook, Color: Black}  // c4 — enemy piece; capturable, then stops
				return GameState{Board: b, ActiveColor: White}, sq(4, 3)
			},
			// d4 reachable, c4 capture included, a4 and b4 unreachable past enemy piece
			want: []string{
				"e4->a8",
				"e4->b1", "e4->b7",
				"e4->c2", "e4->c4", "e4->c6",
				"e4->d3", "e4->d4", "e4->d5",
				"e4->e1", "e4->e2", "e4->e3", "e4->e5", "e4->e6", "e4->e7", "e4->e8",
				"e4->f3", "e4->f4", "e4->f5",
				"e4->g2", "e4->g4", "e4->g6",
				"e4->h1", "e4->h4", "e4->h7",
			},
		},
		{
			name: "corner a1 — 21 targets",
			setup: func() (GameState, Square) {
				var b Board
				b[0][0] = &Piece{Kind: Queen, Color: White} // a1
				return GameState{Board: b, ActiveColor: White}, sq(0, 0)
			},
			want: []string{
				"a1->a2", "a1->a3", "a1->a4", "a1->a5", "a1->a6", "a1->a7", "a1->a8",
				"a1->b1", "a1->b2",
				"a1->c1", "a1->c3",
				"a1->d1", "a1->d4",
				"a1->e1", "a1->e5",
				"a1->f1", "a1->f6",
				"a1->g1", "a1->g7",
				"a1->h1", "a1->h8",
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			state, from := tc.setup()
			got := moveSet(PseudoLegalMoves(state, from))
			if !equalStrings(got, tc.want) {
				t.Errorf("PseudoLegalMoves =\n  %v\nwant\n  %v", got, tc.want)
			}
		})
	}
}

func equalStrings(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

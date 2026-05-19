package core

import "testing"

// sq is a tiny helper to keep table rows readable.
func sq(file, rank uint8) Square { return Square{File: file, Rank: rank} }

func TestIsSquareAttacked_ByEnemyKing(t *testing.T) {
	// Board with a single Black King on d5 (file=3, rank=4).
	var b Board
	b[4][3] = &Piece{Kind: King, Color: Black}

	cases := []struct {
		name   string
		target Square
		want   bool
	}{
		// All 8 neighbours of d5 are attacked.
		{"c4 attacked", sq(2, 3), true},
		{"d4 attacked", sq(3, 3), true},
		{"e4 attacked", sq(4, 3), true},
		{"c5 attacked", sq(2, 4), true},
		{"e5 attacked", sq(4, 4), true},
		{"c6 attacked", sq(2, 5), true},
		{"d6 attacked", sq(3, 5), true},
		{"e6 attacked", sq(4, 5), true},
		// Non-neighbours and own square are not attacked.
		{"d5 own square not attacked", sq(3, 4), false},
		{"b3 two away not attacked", sq(1, 2), false},
		{"a1 far corner not attacked", sq(0, 0), false},
		{"d7 two away not attacked", sq(3, 6), false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsSquareAttacked(b, tc.target, Black); got != tc.want {
				t.Errorf("IsSquareAttacked(%+v, by Black) = %v, want %v", tc.target, got, tc.want)
			}
		})
	}
}

func TestIsSquareAttacked_ByQueen(t *testing.T) {
	// White Queen on d5 (file=3, rank=4).
	var b Board
	b[4][3] = &Piece{Kind: Queen, Color: White}

	cases := []struct {
		name   string
		target Square
		want   bool
	}{
		// Same rank.
		{"h5 same rank attacked", sq(7, 4), true},
		{"a5 same rank attacked", sq(0, 4), true},
		// Same file.
		{"d8 same file attacked", sq(3, 7), true},
		{"d1 same file attacked", sq(3, 0), true},
		// Diagonals.
		{"g8 NE diagonal attacked", sq(6, 7), true},
		{"a2 SW diagonal attacked", sq(0, 1), true},
		{"f3 SE diagonal attacked", sq(5, 2), true},
		{"b7 NW diagonal attacked", sq(1, 6), true},
		// Off-axis squares (not rank, file, or diagonal).
		{"e7 off-axis not attacked", sq(4, 6), false},
		{"b4 off-axis not attacked", sq(1, 3), false},
		// Own square not attacked.
		{"d5 own square not attacked", sq(3, 4), false},
		// Blocker on ray: piece between queen and target blocks the attack.
		{"f5 blocked by e5 not attacked", sq(5, 4), false},
	}

	// Black piece on e5: blocks the Queen's east ray without being a White attacker.
	blocker := b
	blocker[4][4] = &Piece{Kind: Rook, Color: Black}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			board := b
			if tc.name == "f5 blocked by e5 not attacked" {
				board = blocker
			}
			if got := IsSquareAttacked(board, tc.target, White); got != tc.want {
				t.Errorf("IsSquareAttacked(%+v, by White) = %v, want %v", tc.target, got, tc.want)
			}
		})
	}
}

func TestIsSquareAttacked_ByRook(t *testing.T) {
	// White Rook on d5 (file=3, rank=4).
	var b Board
	b[4][3] = &Piece{Kind: Rook, Color: White}

	cases := []struct {
		name   string
		target Square
		want   bool
	}{
		// Same rank.
		{"h5 same rank attacked", sq(7, 4), true},
		{"a5 same rank attacked", sq(0, 4), true},
		// Same file.
		{"d8 same file attacked", sq(3, 7), true},
		{"d1 same file attacked", sq(3, 0), true},
		// Diagonal squares not attacked.
		{"e6 diagonal not attacked", sq(4, 5), false},
		{"c4 diagonal not attacked", sq(2, 3), false},
		// Own square not attacked.
		{"d5 own square not attacked", sq(3, 4), false},
		// Blocker on ray: piece between rook and target blocks the attack.
		{"f5 blocked by e5 not attacked", sq(5, 4), false},
	}

	// Black piece on e5: blocks the Rook's east ray without being a White attacker.
	blocker := b
	blocker[4][4] = &Piece{Kind: King, Color: Black}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			board := b
			if tc.name == "f5 blocked by e5 not attacked" {
				board = blocker
			}
			if got := IsSquareAttacked(board, tc.target, White); got != tc.want {
				t.Errorf("IsSquareAttacked(%+v, by White) = %v, want %v", tc.target, got, tc.want)
			}
		})
	}
}

func TestIsSquareAttacked_WrongColorIgnored(t *testing.T) {
	// White King on e4. Asking "is d5 attacked by Black?" must be false even
	// though a King sits adjacent — the attacker is the wrong color.
	var b Board
	b[3][4] = &Piece{Kind: King, Color: White}

	if got := IsSquareAttacked(b, sq(3, 4), Black); got {
		t.Errorf("IsSquareAttacked(d5, by Black) = true, want false (only White King present)")
	}
}

func TestIsSquareAttacked_NoAttackerOnEmptyBoard(t *testing.T) {
	var b Board
	if got := IsSquareAttacked(b, sq(3, 3), White); got {
		t.Errorf("empty board: IsSquareAttacked = true, want false")
	}
}

func TestIsSquareAttacked_ByBishop(t *testing.T) {
	// White Bishop on d5 (file=3, rank=4).
	var b Board
	b[4][3] = &Piece{Kind: Bishop, Color: White}

	cases := []struct {
		name   string
		target Square
		board  Board
		want   bool
	}{
		{"g8 NE diagonal attacked", sq(6, 7), b, true},
		{"a2 SW diagonal attacked", sq(0, 1), b, true},
		{"f3 SE diagonal attacked", sq(5, 2), b, true},
		{"b7 NW diagonal attacked", sq(1, 6), b, true},
		{"same rank not attacked", sq(7, 4), b, false},
		{"same file not attacked", sq(3, 7), b, false},
		{"off-axis not attacked", sq(4, 6), b, false},
		{"own square not attacked", sq(3, 4), b, false},
	}

	// Add blocker test: Black piece on e6 blocks NE ray past it.
	withBlocker := b
	withBlocker[5][4] = &Piece{Kind: Rook, Color: Black} // e6 blocks NE ray
	cases = append(cases,
		struct {
			name   string
			target Square
			board  Board
			want   bool
		}{"f7 blocked by e6 not attacked", sq(5, 6), withBlocker, false},
	)

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsSquareAttacked(tc.board, tc.target, White); got != tc.want {
				t.Errorf("IsSquareAttacked(%+v, by White) = %v, want %v", tc.target, got, tc.want)
			}
		})
	}
}

func TestIsSquareAttacked_ByKnight(t *testing.T) {
	// White Knight on d5 (file=3, rank=4).
	var b Board
	b[4][3] = &Piece{Kind: Knight, Color: White}

	cases := []struct {
		name   string
		target Square
		want   bool
	}{
		// All 8 L-shaped targets from d5.
		{"c3 attacked", sq(2, 2), true},
		{"e3 attacked", sq(4, 2), true},
		{"b4 attacked", sq(1, 3), true},
		{"f4 attacked", sq(5, 3), true},
		{"b6 attacked", sq(1, 5), true},
		{"f6 attacked", sq(5, 5), true},
		{"c7 attacked", sq(2, 6), true},
		{"e7 attacked", sq(4, 6), true},
		// Non-L-shaped squares not attacked.
		{"d4 adjacent not attacked", sq(3, 3), false},
		{"e6 diagonal not attacked", sq(4, 5), false},
		{"a5 same rank not attacked", sq(0, 4), false},
	}

	// Verify knight ignores blockers: pack all adjacent squares, targets still attacked.
	withBlockers := b
	withBlockers[3][3] = &Piece{Kind: Pawn, Color: Black} // d4
	withBlockers[5][3] = &Piece{Kind: Pawn, Color: Black} // d6
	withBlockers[4][2] = &Piece{Kind: Pawn, Color: Black} // c5
	withBlockers[4][4] = &Piece{Kind: Pawn, Color: Black} // e5

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsSquareAttacked(b, tc.target, White); got != tc.want {
				t.Errorf("IsSquareAttacked(%+v, by White) = %v, want %v", tc.target, got, tc.want)
			}
		})
	}

	t.Run("ignores blockers on adjacent squares", func(t *testing.T) {
		if got := IsSquareAttacked(withBlockers, sq(2, 2), White); !got {
			t.Errorf("Knight should still attack c3 even with all adjacent squares occupied")
		}
	})
}

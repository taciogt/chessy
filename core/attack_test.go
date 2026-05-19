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

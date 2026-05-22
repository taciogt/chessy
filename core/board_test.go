package core

import "testing"

func TestStartingBoard_PieceCount(t *testing.T) {
	b := NewStartingBoard()

	var total, white, black int
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			p := b[rank][file]
			if p == nil {
				continue
			}
			total++
			switch p.Color {
			case White:
				white++
			case Black:
				black++
			}
		}
	}

	if total != 32 {
		t.Errorf("total pieces = %d, want 32", total)
	}
	if white != 16 {
		t.Errorf("white pieces = %d, want 16", white)
	}
	if black != 16 {
		t.Errorf("black pieces = %d, want 16", black)
	}
}

func TestStartingBoard_Layout(t *testing.T) {
	b := NewStartingBoard()

	// Rank index 0 = rank "1" (White back rank).
	// Rank index 7 = rank "8" (Black back rank).
	// File index 0 = file "a".
	tests := []struct {
		name string
		rank int
		file int
		want *Piece // nil = expected empty
	}{
		// White back rank (a1..h1)
		{"a1 white rook", 0, 0, &Piece{Kind: Rook, Color: White}},
		{"b1 white knight", 0, 1, &Piece{Kind: Knight, Color: White}},
		{"c1 white bishop", 0, 2, &Piece{Kind: Bishop, Color: White}},
		{"d1 white queen", 0, 3, &Piece{Kind: Queen, Color: White}},
		{"e1 white king", 0, 4, &Piece{Kind: King, Color: White}},
		{"f1 white bishop", 0, 5, &Piece{Kind: Bishop, Color: White}},
		{"g1 white knight", 0, 6, &Piece{Kind: Knight, Color: White}},
		{"h1 white rook", 0, 7, &Piece{Kind: Rook, Color: White}},
		// White pawns (a2..h2)
		{"a2 white pawn", 1, 0, &Piece{Kind: Pawn, Color: White}},
		{"d2 white pawn", 1, 3, &Piece{Kind: Pawn, Color: White}},
		{"h2 white pawn", 1, 7, &Piece{Kind: Pawn, Color: White}},
		// Empty middle ranks (sample)
		{"a3 empty", 2, 0, nil},
		{"e4 empty", 3, 4, nil},
		{"d5 empty", 4, 3, nil},
		{"h6 empty", 5, 7, nil},
		// Black pawns (a7..h7)
		{"a7 black pawn", 6, 0, &Piece{Kind: Pawn, Color: Black}},
		{"e7 black pawn", 6, 4, &Piece{Kind: Pawn, Color: Black}},
		{"h7 black pawn", 6, 7, &Piece{Kind: Pawn, Color: Black}},
		// Black back rank (a8..h8)
		{"a8 black rook", 7, 0, &Piece{Kind: Rook, Color: Black}},
		{"b8 black knight", 7, 1, &Piece{Kind: Knight, Color: Black}},
		{"c8 black bishop", 7, 2, &Piece{Kind: Bishop, Color: Black}},
		{"d8 black queen", 7, 3, &Piece{Kind: Queen, Color: Black}},
		{"e8 black king", 7, 4, &Piece{Kind: King, Color: Black}},
		{"f8 black bishop", 7, 5, &Piece{Kind: Bishop, Color: Black}},
		{"g8 black knight", 7, 6, &Piece{Kind: Knight, Color: Black}},
		{"h8 black rook", 7, 7, &Piece{Kind: Rook, Color: Black}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := b[tc.rank][tc.file]
			if tc.want == nil {
				if got != nil {
					t.Errorf("b[%d][%d] = %+v, want nil", tc.rank, tc.file, *got)
				}
				return
			}
			if got == nil {
				t.Fatalf("b[%d][%d] = nil, want %+v", tc.rank, tc.file, *tc.want)
			}
			if *got != *tc.want {
				t.Errorf("b[%d][%d] = %+v, want %+v", tc.rank, tc.file, *got, *tc.want)
			}
		})
	}
}

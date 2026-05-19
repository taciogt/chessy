package core

import "testing"

func TestLegalMoves_KingsCanNeverBeAdjacent(t *testing.T) {
	// White King on e4, Black King on g4. White to move.
	// Pseudo-legal King targets for e4: d3, d4, d5, e3, e5, f3, f4, f5.
	// f3, f4, f5 are all adjacent to the Black King on g4 and would leave
	// the White King under attack. Legal set must exclude those three.
	var b Board
	b[3][4] = &Piece{Kind: King, Color: White} // e4
	b[3][6] = &Piece{Kind: King, Color: Black} // g4
	state := GameState{Board: b, ActiveColor: White}

	got := moveSet(LegalMoves(state))
	want := []string{
		"e4->d3", "e4->d4", "e4->d5",
		"e4->e3", "e4->e5",
	}
	if !equalStrings(got, want) {
		t.Errorf("LegalMoves =\n  %v\nwant\n  %v", got, want)
	}
}

func TestLegalMoves_OnlyActiveColorPieces(t *testing.T) {
	// White King on e4, Black King on a8. White to move.
	// The Black King's pseudo-legal moves must NOT appear in LegalMoves.
	var b Board
	b[3][4] = &Piece{Kind: King, Color: White}
	b[7][0] = &Piece{Kind: King, Color: Black}
	state := GameState{Board: b, ActiveColor: White}

	for _, m := range LegalMoves(state) {
		piece := b[m.From.Rank][m.From.File]
		if piece == nil {
			t.Fatalf("legal move %+v originates from empty square", m)
		}
		if piece.Color != White {
			t.Errorf("legal move %+v originates from %v piece, want only White", m, piece.Color)
		}
	}
}

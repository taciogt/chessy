package core

import "testing"

func TestComputeStatus_StartingPositionOngoing(t *testing.T) {
	// Even though only King movement is implemented this slice, the starting
	// position must report Ongoing — terminal detection must respect the
	// fact that not every piece's moves are generated yet.
	if got := ComputeStatus(NewGame()); got != Ongoing {
		t.Errorf("ComputeStatus(starting position) = %d, want Ongoing (%d)", got, Ongoing)
	}
}

func TestComputeStatus_Stalemate_KvK(t *testing.T) {
	// White to move, alone, with two enemy Kings (stand-ins for any later
	// piece types) controlling every escape square. WKa1 has no legal moves
	// and is not under attack → Stalemate.
	//
	//   3 k . . . . . . .   (BK a3 attacks a2, b2)
	//   2 . . k . . . . .   (BK c2 attacks b1, b2)
	//   1 K . . . . . . .   (WK a1 — no escape, not attacked)
	//     a b c d e f g h
	var b Board
	b[0][0] = &Piece{Kind: King, Color: White}
	b[2][0] = &Piece{Kind: King, Color: Black}
	b[1][2] = &Piece{Kind: King, Color: Black}
	state := GameState{Board: b, ActiveColor: White}

	if got := ComputeStatus(state); got != Stalemate {
		t.Errorf("ComputeStatus = %d, want Stalemate (%d)", got, Stalemate)
	}
}

func TestComputeStatus_Checkmate_KvK(t *testing.T) {
	// WKa1 attacked by BKb2. Escape squares a2, b1, b2 all unsafe:
	//   a2 — attacked by BK a3
	//   b1 — attacked by BK b2
	//   b2 — capture, but new WK on b2 would be attacked by BK a3
	// All filtered → Checkmate. Two enemy Kings stand in for the attacker +
	// defender pair that a single later-implemented piece (e.g. Queen) would
	// play in a real mate; this slice can only generate King attacks.
	//
	//   3 k . . . . . . .   (BK a3 defends b2 and covers a2)
	//   2 . k . . . . . .   (BK b2 — the attacker)
	//   1 K . . . . . . .   (WK a1 — in check, no legal reply)
	//     a b c d e f g h
	var b Board
	b[0][0] = &Piece{Kind: King, Color: White}
	b[1][1] = &Piece{Kind: King, Color: Black}
	b[2][0] = &Piece{Kind: King, Color: Black}
	state := GameState{Board: b, ActiveColor: White}

	if got := ComputeStatus(state); got != Checkmate {
		t.Errorf("ComputeStatus = %d, want Checkmate (%d)", got, Checkmate)
	}
}

func TestComputeStatus_Check_KingAttackedWithEscape(t *testing.T) {
	// WKe4 attacked by BKf4. The two-Kings-adjacent position can't arise in
	// a legal game, but it's the cleanest way to exercise the Check branch
	// while we only have King-attack detection. WK has escape squares (d3,
	// d4, d5, plus capturing the BK), so this is Check, not Checkmate.
	var b Board
	b[3][4] = &Piece{Kind: King, Color: White}
	b[3][5] = &Piece{Kind: King, Color: Black}
	state := GameState{Board: b, ActiveColor: White}

	if got := ComputeStatus(state); got != Check {
		t.Errorf("ComputeStatus = %d, want Check (%d)", got, Check)
	}
}

func TestComputeStatus_OngoingWhenKingsCanShuffle(t *testing.T) {
	// Pure KvK, kings far apart — both have many legal moves. Ongoing.
	var b Board
	b[0][0] = &Piece{Kind: King, Color: White}
	b[7][7] = &Piece{Kind: King, Color: Black}
	state := GameState{Board: b, ActiveColor: White}

	if got := ComputeStatus(state); got != Ongoing {
		t.Errorf("ComputeStatus = %d, want Ongoing (%d)", got, Ongoing)
	}
}

func TestComputeStatus_Checkmate_KQK(t *testing.T) {
	// Back-rank mate: WK g6, WQ h7, BK h8. Black to move.
	// BK is in check (WQ on h7 attacks h8 via same file).
	// g8 — attacked by WQ (NW diagonal from h7)
	// g7 — attacked by WK (adjacent)
	// h7 — WQ occupies it; capture impossible (new BK on h7 attacked by WK on g6)
	// All moves illegal → Checkmate.
	//
	//   8 . . . . . . . k   (BK h8 — in check, no escape)
	//   7 . . . . . . . Q   (WQ h7 — gives check, defended by WK)
	//   6 . . . . . . K .   (WK g6)
	//     a b c d e f g h
	var b Board
	b[7][7] = &Piece{Kind: King, Color: Black}  // h8
	b[6][7] = &Piece{Kind: Queen, Color: White} // h7
	b[5][6] = &Piece{Kind: King, Color: White}  // g6
	state := GameState{Board: b, ActiveColor: Black}

	if got := ComputeStatus(state); got != Checkmate {
		t.Errorf("ComputeStatus = %d, want Checkmate (%d)", got, Checkmate)
	}
}

func TestComputeStatus_Check_KQK(t *testing.T) {
	// WQ a5 gives check to BK a8 (same file). BK has escape squares b8 and b7.
	//
	//   8 k . . . . . . .   (BK a8 — in check, has escapes)
	//   5 Q . . . . . . .   (WQ a5 — checks along a-file)
	//   3 . . K . . . . .   (WK c3)
	//     a b c d e f g h
	var b Board
	b[7][0] = &Piece{Kind: King, Color: Black}  // a8
	b[4][0] = &Piece{Kind: Queen, Color: White} // a5
	b[2][2] = &Piece{Kind: King, Color: White}  // c3
	state := GameState{Board: b, ActiveColor: Black}

	if got := ComputeStatus(state); got != Check {
		t.Errorf("ComputeStatus = %d, want Check (%d)", got, Check)
	}
}

func TestComputeStatus_KQK_ActiveQueenSide_Ongoing(t *testing.T) {
	// White (K+Q) to move, kings far apart — White has many legal moves → Ongoing.
	// This exercises the guard path for the active side having a Queen; after the
	// guard fix (Queen added to implemented set), status is computed via proper
	// legal-move enumeration rather than the phased-implementation downgrade.
	//
	//   8 . . . . . . . k   (BK h8 — far away)
	//   1 K Q . . . . . .   (WK a1, WQ b1)
	//     a b c d e f g h
	var b Board
	b[0][0] = &Piece{Kind: King, Color: White}  // a1
	b[0][1] = &Piece{Kind: Queen, Color: White} // b1
	b[7][7] = &Piece{Kind: King, Color: Black}  // h8
	state := GameState{Board: b, ActiveColor: White}

	if got := ComputeStatus(state); got != Ongoing {
		t.Errorf("ComputeStatus = %d, want Ongoing (%d)", got, Ongoing)
	}
}

func TestComputeStatus_Checkmate_KRK(t *testing.T) {
	// BK a8 in check by WR a1 (a-file). WK c7 covers b8 and b7.
	// BK escapes: a7 attacked by WR, b7 attacked by WK, b8 attacked by WK.
	//
	//   8 k . . . . . . .   (BK a8 — in check, no escape)
	//   7 . . K . . . . .   (WK c7 — covers b8, b7)
	//   1 R . . . . . . .   (WR a1 — gives check)
	//     a b c d e f g h
	var b Board
	b[7][0] = &Piece{Kind: King, Color: Black}  // a8
	b[0][0] = &Piece{Kind: Rook, Color: White}  // a1
	b[6][2] = &Piece{Kind: King, Color: White}  // c7
	state := GameState{Board: b, ActiveColor: Black}

	if got := ComputeStatus(state); got != Checkmate {
		t.Errorf("ComputeStatus = %d, want Checkmate (%d)", got, Checkmate)
	}
}

func TestComputeStatus_Check_KRK(t *testing.T) {
	// BK a8 in check by WR a1 (a-file). WK c6 covers b7 but NOT b8.
	// BK can escape to b8 → Check, not Checkmate.
	//
	//   8 k . . . . . . .   (BK a8 — in check, can escape to b8)
	//   6 . . K . . . . .   (WK c6 — covers b7, not b8)
	//   1 R . . . . . . .   (WR a1 — gives check)
	//     a b c d e f g h
	var b Board
	b[7][0] = &Piece{Kind: King, Color: Black}  // a8
	b[0][0] = &Piece{Kind: Rook, Color: White}  // a1
	b[5][2] = &Piece{Kind: King, Color: White}  // c6
	state := GameState{Board: b, ActiveColor: Black}

	if got := ComputeStatus(state); got != Check {
		t.Errorf("ComputeStatus = %d, want Check (%d)", got, Check)
	}
}

func TestComputeStatus_KRK_ActiveRookSide_Ongoing(t *testing.T) {
	// White (K+R) to move, kings far apart — White has many legal moves → Ongoing.
	// Exercises the guard path for the active side having a Rook; after adding
	// Rook to the implemented set, status is computed via proper legal-move
	// enumeration rather than the phased-implementation downgrade.
	//
	//   8 . . . . . . . k   (BK h8 — far away)
	//   1 K R . . . . . .   (WK a1, WR b1)
	//     a b c d e f g h
	var b Board
	b[0][0] = &Piece{Kind: King, Color: White} // a1
	b[0][1] = &Piece{Kind: Rook, Color: White} // b1
	b[7][7] = &Piece{Kind: King, Color: Black} // h8
	state := GameState{Board: b, ActiveColor: White}

	if got := ComputeStatus(state); got != Ongoing {
		t.Errorf("ComputeStatus = %d, want Ongoing (%d)", got, Ongoing)
	}
}

func TestComputeStatus_Checkmate_ScholarsMate(t *testing.T) {
	// Scholar's Mate final position after 1.e4 e5 2.Bc4 Nc6 3.Qh5 Nf6 4.Qxf7#
	// Black to move, in check from WQ on f7 (attacks e8 diagonally).
	//
	//   8 r . b q k b n r   (BK e8 — in check, no escape)
	//   7 p p p p . Q p p   (WQ f7 — gives check; Black f7 pawn was captured)
	//   6 . . n . . n . .   (Black knights c6, f6)
	//   5 . . . . p . . .   (Black e-pawn)
	//   4 . . B . P . . .   (WB c4, WP e4)
	//   3 . . . . . . . .
	//   2 P P P P . P P P   (White pawns)
	//   1 R N B . K . N R   (White back rank)
	//     a b c d e f g h
	//
	// Escape analysis:
	//   d7 — own Black pawn blocks
	//   d8 — own Black Queen blocks
	//   e7 — attacked by WQ f7 (same rank)
	//   f7 — WQ there; if captured, WB c4 attacks via NE diagonal (d5, e6, f7)
	//   f8 — own Black Bishop blocks
	var b Board
	// White pieces
	b[0][0] = &Piece{Kind: Rook, Color: White}   // a1
	b[0][1] = &Piece{Kind: Knight, Color: White} // b1
	b[0][2] = &Piece{Kind: Bishop, Color: White} // c1
	b[0][4] = &Piece{Kind: King, Color: White}   // e1
	b[0][6] = &Piece{Kind: Knight, Color: White} // g1
	b[0][7] = &Piece{Kind: Rook, Color: White}   // h1
	b[3][2] = &Piece{Kind: Bishop, Color: White} // c4 (moved from f1)
	b[3][4] = &Piece{Kind: Pawn, Color: White}   // e4
	b[6][5] = &Piece{Kind: Queen, Color: White}  // f7 (gives check)
	b[1][0] = &Piece{Kind: Pawn, Color: White}   // a2
	b[1][1] = &Piece{Kind: Pawn, Color: White}   // b2
	b[1][2] = &Piece{Kind: Pawn, Color: White}   // c2
	b[1][3] = &Piece{Kind: Pawn, Color: White}   // d2
	b[1][5] = &Piece{Kind: Pawn, Color: White}   // f2
	b[1][6] = &Piece{Kind: Pawn, Color: White}   // g2
	b[1][7] = &Piece{Kind: Pawn, Color: White}   // h2
	// Black pieces
	b[7][0] = &Piece{Kind: Rook, Color: Black}   // a8
	b[7][2] = &Piece{Kind: Bishop, Color: Black} // c8
	b[7][3] = &Piece{Kind: Queen, Color: Black}  // d8 — blocks d8 escape
	b[7][4] = &Piece{Kind: King, Color: Black}   // e8
	b[7][5] = &Piece{Kind: Bishop, Color: Black} // f8 — blocks f8 escape
	b[7][6] = &Piece{Kind: Knight, Color: Black} // g8
	b[7][7] = &Piece{Kind: Rook, Color: Black}   // h8
	b[5][2] = &Piece{Kind: Knight, Color: Black} // c6
	b[5][5] = &Piece{Kind: Knight, Color: Black} // f6
	b[4][4] = &Piece{Kind: Pawn, Color: Black}   // e5
	b[6][0] = &Piece{Kind: Pawn, Color: Black}   // a7
	b[6][1] = &Piece{Kind: Pawn, Color: Black}   // b7
	b[6][2] = &Piece{Kind: Pawn, Color: Black}   // c7
	b[6][3] = &Piece{Kind: Pawn, Color: Black}   // d7 — blocks d7 escape
	b[6][6] = &Piece{Kind: Pawn, Color: Black}   // g7
	b[6][7] = &Piece{Kind: Pawn, Color: Black}   // h7

	state := GameState{Board: b, ActiveColor: Black}

	if got := ComputeStatus(state); got != Checkmate {
		t.Errorf("ComputeStatus(Scholar's Mate) = %d, want Checkmate (%d)", got, Checkmate)
	}
}

func TestGameStatus_DistinctValues(t *testing.T) {
	values := map[string]GameStatus{
		"Ongoing":   Ongoing,
		"Check":     Check,
		"Checkmate": Checkmate,
		"Stalemate": Stalemate,
	}

	seen := map[GameStatus]string{}
	for name, v := range values {
		if other, dup := seen[v]; dup {
			t.Errorf("%s and %s share value %d, must be distinct", name, other, v)
		}
		seen[v] = name
	}

	if len(seen) != 4 {
		t.Errorf("got %d distinct GameStatus values, want 4", len(seen))
	}
}

package core

import "testing"

func TestApplyMove_Relocates(t *testing.T) {
	g := NewGame()
	// e2 -> e4 in [rank][file]: (rank=1, file=4) -> (rank=3, file=4)
	m := Move{
		From: Square{File: 4, Rank: 1},
		To:   Square{File: 4, Rank: 3},
	}

	g2 := ApplyMove(g, m)

	if g2.Board[1][4] != nil {
		t.Errorf("From square e2 should be empty after move, got %+v", *g2.Board[1][4])
	}
	got := g2.Board[3][4]
	if got == nil {
		t.Fatalf("To square e4 should hold the moved pawn, got nil")
	}
	if got.Kind != Pawn || got.Color != White {
		t.Errorf("To square e4 = %+v, want White Pawn", *got)
	}
}

func TestApplyMove_DoesNotMutateInput(t *testing.T) {
	g := NewGame()
	m := Move{
		From: Square{File: 4, Rank: 1},
		To:   Square{File: 4, Rank: 3},
	}

	_ = ApplyMove(g, m)

	// Input board must still show pawn on e2 and empty e4.
	if p := g.Board[1][4]; p == nil || p.Kind != Pawn || p.Color != White {
		t.Errorf("input board e2 = %+v, want White Pawn (input must not mutate)", p)
	}
	if g.Board[3][4] != nil {
		t.Errorf("input board e4 = %+v, want nil (input must not mutate)", *g.Board[3][4])
	}
	// History must still be empty.
	if len(g.History) != 0 {
		t.Errorf("input history length = %d, want 0", len(g.History))
	}
}

func TestApplyMove_FlipsActiveColor(t *testing.T) {
	g := NewGame() // White to move
	m := Move{From: Square{File: 4, Rank: 1}, To: Square{File: 4, Rank: 3}}

	after := ApplyMove(g, m)
	if after.ActiveColor != Black {
		t.Errorf("ActiveColor after White move = %d, want Black (%d)", after.ActiveColor, Black)
	}

	after2 := ApplyMove(after, Move{From: Square{File: 4, Rank: 6}, To: Square{File: 4, Rank: 4}})
	if after2.ActiveColor != White {
		t.Errorf("ActiveColor after Black move = %d, want White (%d)", after2.ActiveColor, White)
	}

	// Caller's state must remain untouched.
	if g.ActiveColor != White {
		t.Errorf("input ActiveColor = %d, want White (immutability)", g.ActiveColor)
	}
}

func TestApplyMove_AppendsHistory(t *testing.T) {
	g := NewGame()
	m := Move{
		From: Square{File: 4, Rank: 1},
		To:   Square{File: 4, Rank: 3},
	}

	g2 := ApplyMove(g, m)

	if got, want := len(g2.History), 1; got != want {
		t.Fatalf("returned History length = %d, want %d", got, want)
	}
	if g2.History[0] != m {
		t.Errorf("History[0] = %+v, want %+v", g2.History[0], m)
	}
}

func TestApplyMove_HistoryNotAliased(t *testing.T) {
	// Seed a state whose History has spare capacity. If ApplyMove reuses the
	// caller's backing array, two derived states will clobber each other.
	g := NewGame()
	filler := Move{From: Square{File: 0, Rank: 0}, To: Square{File: 7, Rank: 7}}
	g.History = make([]Move, 0, 8)
	g.History = append(g.History, filler)

	m1 := Move{From: Square{File: 4, Rank: 1}, To: Square{File: 4, Rank: 3}}
	m2 := Move{From: Square{File: 3, Rank: 1}, To: Square{File: 3, Rank: 3}}

	g1 := ApplyMove(g, m1)
	g2 := ApplyMove(g, m2)

	if g1.History[1] != m1 {
		t.Errorf("history aliased: g1.History[1] = %+v, want %+v", g1.History[1], m1)
	}
	if g2.History[1] != m2 {
		t.Errorf("history aliased: g2.History[1] = %+v, want %+v", g2.History[1], m2)
	}
}

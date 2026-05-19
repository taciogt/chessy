package core

import "testing"

func TestNewGame_InitialState(t *testing.T) {
	g := NewGame()

	if g.ActiveColor != White {
		t.Errorf("ActiveColor = %d, want White (%d)", g.ActiveColor, White)
	}
	if len(g.History) != 0 {
		t.Errorf("History length = %d, want 0", len(g.History))
	}
	if g.EnPassantTarget != nil {
		t.Errorf("EnPassantTarget = %+v, want nil", g.EnPassantTarget)
	}
	if g.CastlingRights != (CastlingRights{}) {
		t.Errorf("CastlingRights = %+v, want zero-valued", g.CastlingRights)
	}
	// Board must be the starting position (sanity-check one square).
	if p := g.Board[0][4]; p == nil || p.Kind != King || p.Color != White {
		t.Errorf("Board[0][4] = %+v, want White King", p)
	}
}

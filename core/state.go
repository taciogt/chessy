package core

// CastlingRights tracks whether each side may still castle on each flank.
// Zero value = no rights yet (filled in by a future rules phase).
type CastlingRights struct {
	WhiteKingside  bool
	WhiteQueenside bool
	BlackKingside  bool
	BlackQueenside bool
}

// GameState is the full snapshot of a chess match.
// Immutable by convention: ApplyMove returns a new value.
type GameState struct {
	Board           Board
	ActiveColor     Color
	History         []Move
	CastlingRights  CastlingRights
	EnPassantTarget *Square
}

// NewGame returns the canonical opening position: standard board,
// White to move, empty history, no castling rights yet, no en-passant target.
func NewGame() GameState {
	return GameState{
		Board:       NewStartingBoard(),
		ActiveColor: White,
	}
}

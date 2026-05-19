package core

// GameStatus is the outcome state of a position.
type GameStatus uint8

const (
	Ongoing GameStatus = iota
	Check
	Checkmate
	Stalemate
)

// ComputeStatus returns the current outcome of s for the side to move.
//
// Terminal statuses (Checkmate, Stalemate) require an exhaustive legal-move
// search. While the engine is built up piece-by-piece, this slice only
// generates moves for Kings; so "no legal moves" is only trusted when the
// active side has nothing but Kings on the board. In any other case we
// downgrade to Check (if the King is attacked) or Ongoing — never falsely
// terminating a game just because, say, Pawn moves aren't wired in yet.
func ComputeStatus(s GameState) GameStatus {
	kingSq, hasKing := findKing(s.Board, s.ActiveColor)
	var attacked bool
	if hasKing {
		attacked = IsSquareAttacked(s.Board, kingSq, opposite(s.ActiveColor))
	}

	if !activeSideHasOnlyImplementedPieces(s) {
		if attacked {
			return Check
		}
		return Ongoing
	}

	noMoves := len(LegalMoves(s)) == 0
	switch {
	case attacked && noMoves:
		return Checkmate
	case attacked:
		return Check
	case noMoves:
		return Stalemate
	default:
		return Ongoing
	}
}

// activeSideHasOnlyImplementedPieces is a phased-implementation guard. Today
// only Kings have movement logic, so a position is only safe to terminate
// when the active side consists entirely of Kings. Each future slice extends
// the implemented-kinds set as it lands.
func activeSideHasOnlyImplementedPieces(s GameState) bool {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			p := s.Board[rank][file]
			if p != nil && p.Color == s.ActiveColor && p.Kind != King {
				return false
			}
		}
	}
	return true
}

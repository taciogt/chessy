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

// activeSideHasOnlyImplementedPieces is a phased-implementation guard that
// prevents false terminal-state detection. All piece types are now implemented.
func activeSideHasOnlyImplementedPieces(s GameState) bool {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			p := s.Board[rank][file]
			if p == nil || p.Color != s.ActiveColor {
				continue
			}
			switch p.Kind {
			case King, Queen, Rook, Bishop, Knight, Pawn:
				// all piece types implemented
			default:
				return false
			}
		}
	}
	return true
}

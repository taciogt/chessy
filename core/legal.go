package core

// LegalMoves returns every move the active color may play in s: pseudo-legal
// moves filtered to those that do not leave the active King attacked.
func LegalMoves(s GameState) []Move {
	var legal []Move
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			p := s.Board[rank][file]
			if p == nil || p.Color != s.ActiveColor {
				continue
			}
			from := Square{File: uint8(file), Rank: uint8(rank)}
			for _, m := range PseudoLegalMoves(s, from) {
				if !leavesKingInCheck(s, m) {
					legal = append(legal, m)
				}
			}
		}
	}
	return legal
}

func leavesKingInCheck(s GameState, m Move) bool {
	next := ApplyMove(s, m)
	kingSq, ok := findKing(next.Board, s.ActiveColor)
	if !ok {
		return false
	}
	return IsSquareAttacked(next.Board, kingSq, opposite(s.ActiveColor))
}

func findKing(b Board, color Color) (Square, bool) {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			p := b[rank][file]
			if p != nil && p.Kind == King && p.Color == color {
				return Square{File: uint8(file), Rank: uint8(rank)}, true
			}
		}
	}
	return Square{}, false
}

func opposite(c Color) Color {
	if c == White {
		return Black
	}
	return White
}

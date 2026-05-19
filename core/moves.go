package core

// PseudoLegalMoves returns the moves the piece on `from` can geometrically make
// in the position s, ignoring whether the active King would be left in check.
// Returns nil when `from` is empty or holds a piece of the inactive color.
//
// MVP scope: only the King has movement logic. Other piece types return nil and
// will be filled in by later slices.
func PseudoLegalMoves(s GameState, from Square) []Move {
	p := s.Board[from.Rank][from.File]
	if p == nil || p.Color != s.ActiveColor {
		return nil
	}
	switch p.Kind {
	case King:
		return kingPseudoLegalMoves(s.Board, from, p.Color)
	default:
		return nil
	}
}

// kingPseudoLegalMoves enumerates the up-to-8 single-step targets, dropping
// off-board destinations and squares occupied by same-colour pieces.
func kingPseudoLegalMoves(b Board, from Square, mover Color) []Move {
	moves := make([]Move, 0, 8)
	for df := -1; df <= 1; df++ {
		for dr := -1; dr <= 1; dr++ {
			if df == 0 && dr == 0 {
				continue
			}
			tf := int(from.File) + df
			tr := int(from.Rank) + dr
			if tf < 0 || tf > 7 || tr < 0 || tr > 7 {
				continue
			}
			target := b[tr][tf]
			if target != nil && target.Color == mover {
				continue
			}
			moves = append(moves, Move{
				From: from,
				To:   Square{File: uint8(tf), Rank: uint8(tr)},
			})
		}
	}
	return moves
}

package core

// PseudoLegalMoves returns the moves the piece on `from` can geometrically make
// in the position s, ignoring whether the active King would be left in check.
// Returns nil when `from` is empty or holds a piece of the inactive color.
func PseudoLegalMoves(s GameState, from Square) []Move {
	p := s.Board[from.Rank][from.File]
	if p == nil || p.Color != s.ActiveColor {
		return nil
	}
	switch p.Kind {
	case King:
		return kingPseudoLegalMoves(s.Board, from, p.Color)
	case Queen:
		return queenPseudoLegalMoves(s.Board, from, p.Color)
	case Rook:
		return rookPseudoLegalMoves(s.Board, from, p.Color)
	case Bishop:
		return bishopPseudoLegalMoves(s.Board, from, p.Color)
	case Knight:
		return knightPseudoLegalMoves(s.Board, from, p.Color)
	case Pawn:
		return pawnPseudoLegalMoves(s.Board, from, p.Color)
	default:
		return nil
	}
}

// pawnPseudoLegalMoves returns the up-to-4 moves a pawn can make: one square
// forward (if empty), two squares forward from the starting rank (if both
// squares are empty), and diagonal captures (one square forward-diagonal if
// occupied by an enemy piece). En passant and promotion are out of scope for MVP.
func pawnPseudoLegalMoves(b Board, from Square, mover Color) []Move {
	moves := make([]Move, 0, 4)
	var forwardDir int
	var startRank uint8
	if mover == White {
		forwardDir = 1
		startRank = 1
	} else {
		forwardDir = -1
		startRank = 6
	}

	nextRank := int(from.Rank) + forwardDir
	if nextRank < 0 || nextRank > 7 {
		return moves // pawn at back rank — no promotion in MVP
	}

	// Forward 1: only if destination is empty.
	if b[nextRank][from.File] == nil {
		moves = append(moves, Move{From: from, To: Square{File: from.File, Rank: uint8(nextRank)}})

		// Forward 2: only from starting rank and only when forward 1 was clear.
		if from.Rank == startRank {
			doubleRank := int(from.Rank) + 2*forwardDir
			if b[doubleRank][from.File] == nil {
				moves = append(moves, Move{From: from, To: Square{File: from.File, Rank: uint8(doubleRank)}})
			}
		}
	}

	// Diagonal captures: only when an enemy piece occupies the target square.
	for _, df := range [2]int{-1, 1} {
		tf := int(from.File) + df
		if tf < 0 || tf > 7 {
			continue
		}
		target := b[nextRank][tf]
		if target != nil && target.Color != mover {
			moves = append(moves, Move{From: from, To: Square{File: uint8(tf), Rank: uint8(nextRank)}})
		}
	}

	return moves
}

// knightPseudoLegalMoves enumerates all L-shaped jump targets, dropping off-board
// destinations and squares occupied by same-colour pieces. Knights ignore blockers.
func knightPseudoLegalMoves(b Board, from Square, mover Color) []Move {
	moves := make([]Move, 0, 8)
	for _, d := range [8][2]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}} {
		tf := int(from.File) + d[0]
		tr := int(from.Rank) + d[1]
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
	return moves
}

// bishopPseudoLegalMoves slides along the 4 diagonal rays until blocked by the
// board edge, a same-colour piece (excluded), or an enemy piece (captured, then stops).
func bishopPseudoLegalMoves(b Board, from Square, mover Color) []Move {
	moves := make([]Move, 0, 13)
	for _, d := range [4][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		for step := 1; ; step++ {
			tf := int(from.File) + d[0]*step
			tr := int(from.Rank) + d[1]*step
			if tf < 0 || tf > 7 || tr < 0 || tr > 7 {
				break
			}
			target := b[tr][tf]
			if target != nil && target.Color == mover {
				break
			}
			moves = append(moves, Move{
				From: from,
				To:   Square{File: uint8(tf), Rank: uint8(tr)},
			})
			if target != nil {
				break
			}
		}
	}
	return moves
}

// queenPseudoLegalMoves slides along all 8 directions until blocked by the
// board edge, a same-colour piece (excluded), or an enemy piece (included as a
// capture, then stops).
func queenPseudoLegalMoves(b Board, from Square, mover Color) []Move {
	moves := make([]Move, 0, 27)
	for df := -1; df <= 1; df++ {
		for dr := -1; dr <= 1; dr++ {
			if df == 0 && dr == 0 {
				continue
			}
			for step := 1; ; step++ {
				tf := int(from.File) + df*step
				tr := int(from.Rank) + dr*step
				if tf < 0 || tf > 7 || tr < 0 || tr > 7 {
					break
				}
				target := b[tr][tf]
				if target != nil && target.Color == mover {
					break
				}
				moves = append(moves, Move{
					From: from,
					To:   Square{File: uint8(tf), Rank: uint8(tr)},
				})
				if target != nil {
					break // enemy piece captured — stop sliding
				}
			}
		}
	}
	return moves
}

// rookPseudoLegalMoves slides along the 4 orthogonal rays until blocked by the
// board edge, a same-colour piece (excluded), or an enemy piece (captured, then stops).
func rookPseudoLegalMoves(b Board, from Square, mover Color) []Move {
	moves := make([]Move, 0, 14)
	for _, d := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		for step := 1; ; step++ {
			tf := int(from.File) + d[0]*step
			tr := int(from.Rank) + d[1]*step
			if tf < 0 || tf > 7 || tr < 0 || tr > 7 {
				break
			}
			target := b[tr][tf]
			if target != nil && target.Color == mover {
				break
			}
			moves = append(moves, Move{
				From: from,
				To:   Square{File: uint8(tf), Rank: uint8(tr)},
			})
			if target != nil {
				break
			}
		}
	}
	return moves
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

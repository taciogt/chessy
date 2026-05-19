package core

// IsSquareAttacked reports whether any piece of byColor attacks target on b.
// MVP scope: only enemy Kings produce attacks; later slices extend this per
// piece type as their movement logic lands.
func IsSquareAttacked(b Board, target Square, byColor Color) bool {
	for rank := 0; rank < 8; rank++ {
		for file := 0; file < 8; file++ {
			p := b[rank][file]
			if p == nil || p.Color != byColor {
				continue
			}
			from := Square{File: uint8(file), Rank: uint8(rank)}
			switch p.Kind {
			case King:
				if kingAttacks(from, target) {
					return true
				}
			case Queen:
				if queenAttacks(b, from, target) {
					return true
				}
			case Rook:
				if rookAttacks(b, from, target) {
					return true
				}
			case Bishop:
				if bishopAttacks(b, from, target) {
					return true
				}
			case Knight:
				if knightAttacks(from, target) {
					return true
				}
			}
		}
	}
	return false
}

// bishopAttacks slides along the 4 diagonal rays; returns true if target is
// reached before any intervening piece blocks the ray.
func bishopAttacks(b Board, from, target Square) bool {
	for _, d := range [4][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
		for step := 1; ; step++ {
			tf := int(from.File) + d[0]*step
			tr := int(from.Rank) + d[1]*step
			if tf < 0 || tf > 7 || tr < 0 || tr > 7 {
				break
			}
			sq := Square{File: uint8(tf), Rank: uint8(tr)}
			if sq == target {
				return true
			}
			if b[tr][tf] != nil {
				break
			}
		}
	}
	return false
}

// knightAttacks returns true when target is reachable from from by an L-shaped jump.
func knightAttacks(from, target Square) bool {
	df := abs(int(from.File) - int(target.File))
	dr := abs(int(from.Rank) - int(target.Rank))
	return (df == 1 && dr == 2) || (df == 2 && dr == 1)
}

// queenAttacks slides along each of the 8 rays from from; returns true if
// target is reached before any intervening piece blocks the ray.
func queenAttacks(b Board, from, target Square) bool {
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
				sq := Square{File: uint8(tf), Rank: uint8(tr)}
				if sq == target {
					return true
				}
				if b[tr][tf] != nil {
					break // any piece blocks the ray before target
				}
			}
		}
	}
	return false
}

// rookAttacks slides along the 4 orthogonal rays; returns true if target is
// reached before any intervening piece blocks the ray.
func rookAttacks(b Board, from, target Square) bool {
	for _, d := range [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
		for step := 1; ; step++ {
			tf := int(from.File) + d[0]*step
			tr := int(from.Rank) + d[1]*step
			if tf < 0 || tf > 7 || tr < 0 || tr > 7 {
				break
			}
			sq := Square{File: uint8(tf), Rank: uint8(tr)}
			if sq == target {
				return true
			}
			if b[tr][tf] != nil {
				break
			}
		}
	}
	return false
}

// kingAttacks is true when from is within one square of target (Chebyshev distance 1).
func kingAttacks(from, target Square) bool {
	df := int(from.File) - int(target.File)
	dr := int(from.Rank) - int(target.Rank)
	if df == 0 && dr == 0 {
		return false
	}
	return abs(df) <= 1 && abs(dr) <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

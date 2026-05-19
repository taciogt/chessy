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
			if p.Kind == King && kingAttacks(Square{File: uint8(file), Rank: uint8(rank)}, target) {
				return true
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

package core

// GameStatus is the outcome state of a position.
type GameStatus uint8

const (
	Ongoing GameStatus = iota
	Check
	Checkmate
	Stalemate
)

// ComputeStatus reports whether the game is ongoing or terminated.
// MVP stub: always returns Ongoing. Real rules ship in a later slice.
func ComputeStatus(GameState) GameStatus {
	return Ongoing
}

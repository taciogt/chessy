package core

// Move represents a piece transition between two squares.
// Promotion is the zero value (no promotion) until the promotion phase lands.
type Move struct {
	From      Square
	To        Square
	Promotion PieceKind
}

package core

// Square identifies a single cell on the board.
// File 0..7 maps to a..h; Rank 0..7 maps to ranks 1..8.
type Square struct {
	File uint8
	Rank uint8
}

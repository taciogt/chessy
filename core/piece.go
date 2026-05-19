package core

type Color uint8

const (
	White Color = iota
	Black
)

type PieceKind uint8

const (
	Pawn PieceKind = iota + 1
	Rook
	Knight
	Bishop
	Queen
	King
)

type Piece struct {
	Kind  PieceKind
	Color Color
}

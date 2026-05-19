package core

type Board [8][8]*Piece

func NewStartingBoard() Board {
	var b Board
	backRank := [8]PieceKind{Rook, Knight, Bishop, Queen, King, Bishop, Knight, Rook}
	for f := 0; f < 8; f++ {
		b[0][f] = &Piece{Kind: backRank[f], Color: White}
		b[1][f] = &Piece{Kind: Pawn, Color: White}
		b[6][f] = &Piece{Kind: Pawn, Color: Black}
		b[7][f] = &Piece{Kind: backRank[f], Color: Black}
	}
	return b
}

package core

// ApplyMove returns a new GameState with the piece relocated from m.From to m.To,
// the side to move flipped, and m appended to the move history. The input s is
// not mutated.
//
// MVP scope: no legality check is performed here; callers should source moves
// from LegalMoves. The returned History does not share its backing array with
// the input, so two states derived from the same input are independent.
func ApplyMove(s GameState, m Move) GameState {
	s.Board[m.To.Rank][m.To.File] = s.Board[m.From.Rank][m.From.File]
	s.Board[m.From.Rank][m.From.File] = nil

	newHistory := make([]Move, len(s.History)+1)
	copy(newHistory, s.History)
	newHistory[len(s.History)] = m
	s.History = newHistory

	s.ActiveColor = opposite(s.ActiveColor)

	return s
}

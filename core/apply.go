package core

// ApplyMove returns a new GameState with the piece relocated from m.From to m.To
// and m appended to the move history. MVP stub: no legality check, no turn flip.
//
// The returned History does not share its backing array with the input, so two
// states derived from the same input are independent.
func ApplyMove(s GameState, m Move) GameState {
	s.Board[m.To.Rank][m.To.File] = s.Board[m.From.Rank][m.From.File]
	s.Board[m.From.Rank][m.From.File] = nil

	newHistory := make([]Move, len(s.History)+1)
	copy(newHistory, s.History)
	newHistory[len(s.History)] = m
	s.History = newHistory

	return s
}

package core

import (
	"fmt"
	"regexp"
)

// SyntaxError is returned when the SAN input does not match the expected grammar.
type SyntaxError struct {
	Input string
}

func (e SyntaxError) Error() string { return fmt.Sprintf("invalid SAN: %q", e.Input) }

// NoLegalMoveError is returned when no legal move matches the parsed SAN.
type NoLegalMoveError struct {
	Input string
}

func (e NoLegalMoveError) Error() string {
	return fmt.Sprintf("no legal move for SAN: %q", e.Input)
}

// AmbiguousMoveError is returned when more than one legal move matches the SAN
// and no sufficient disambiguator was provided.
type AmbiguousMoveError struct {
	Input      string
	Candidates []Move
}

func (e AmbiguousMoveError) Error() string {
	return fmt.Sprintf("ambiguous SAN %q: %d candidates", e.Input, len(e.Candidates))
}

// sanRegex parses piece moves (K, Q, R, B, N).
// Group 1: piece letter
// Group 2: optional disambiguator (full square, file, or rank)
// Group 3: destination square
var sanRegex = regexp.MustCompile(
	`^([KQRBN])([a-h][1-8]|[a-h]|[1-8])?x?([a-h][1-8])[+#]?$`,
)

var pieceLetterToKind = map[byte]PieceKind{
	'K': King,
	'Q': Queen,
	'R': Rook,
	'B': Bishop,
	'N': Knight,
}

// ParseSAN converts a SAN piece-move string into a Move using state to resolve
// which piece moves and to validate legality. Pawn moves and castling are out of scope.
func ParseSAN(state GameState, input string) (Move, error) {
	m := sanRegex.FindStringSubmatch(input)
	if m == nil {
		return Move{}, SyntaxError{Input: input}
	}

	kind := pieceLetterToKind[m[1][0]]
	disambig := m[2]
	destStr := m[3]
	dest := Square{File: destStr[0] - 'a', Rank: destStr[1] - '1'}

	var candidates []Move
	for _, lm := range LegalMoves(state) {
		p := state.Board[lm.From.Rank][lm.From.File]
		if p == nil || p.Kind != kind || lm.To != dest {
			continue
		}
		if !matchesDisambig(lm.From, disambig) {
			continue
		}
		candidates = append(candidates, lm)
	}

	switch len(candidates) {
	case 0:
		return Move{}, NoLegalMoveError{Input: input}
	case 1:
		return candidates[0], nil
	default:
		return Move{}, AmbiguousMoveError{Input: input, Candidates: candidates}
	}
}

// matchesDisambig returns true when from satisfies the disambiguator string.
// Empty disambig always matches.
func matchesDisambig(from Square, disambig string) bool {
	switch len(disambig) {
	case 0:
		return true
	case 1:
		if disambig[0] >= 'a' && disambig[0] <= 'h' {
			return from.File == disambig[0]-'a'
		}
		return from.Rank == disambig[0]-'1'
	case 2:
		return from.File == disambig[0]-'a' && from.Rank == disambig[1]-'1'
	}
	return false
}

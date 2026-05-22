package core

import (
	"errors"
	"testing"
)

func sparseBoard(pieces map[string]Piece) Board {
	var b Board
	for s, p := range pieces {
		f := s[0] - 'a'
		r := s[1] - '1'
		pc := p
		b[r][f] = &pc
	}
	return b
}

func sqn(s string) Square {
	return Square{File: s[0] - 'a', Rank: s[1] - '1'}
}

// --- Tracer bullet: plain King move ---

func TestParseSAN_KingMove(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e1": {Kind: King, Color: White},
		"h8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "Ke2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("e1"), To: sqn("e2")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

// --- Syntax errors ---

func TestParseSAN_SyntaxError(t *testing.T) {
	state := GameState{Board: NewStartingBoard(), ActiveColor: White}
	cases := []string{"", "xyz", "K9", "Pf3", "ke2", "Ne9"}
	for _, input := range cases {
		t.Run(input, func(t *testing.T) {
			_, err := ParseSAN(state, input)
			var se SyntaxError
			if !errors.As(err, &se) {
				t.Errorf("ParseSAN(%q) err = %T(%v), want SyntaxError", input, err, err)
			}
		})
	}
}

// --- Trailing check/checkmate markers stripped ---

func TestParseSAN_CheckMarkerStripped(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e1": {Kind: King, Color: White},
		"h8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}
	want := Move{From: sqn("e1"), To: sqn("e2")}

	for _, input := range []string{"Ke2+", "Ke2#"} {
		got, err := ParseSAN(state, input)
		if err != nil {
			t.Fatalf("ParseSAN(%q) unexpected error: %v", input, err)
		}
		if got != want {
			t.Errorf("ParseSAN(%q) = %+v, want %+v", input, got, want)
		}
	}
}

// --- Capture ---

func TestParseSAN_KingCapture(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e1": {Kind: King, Color: White},
		"d2": {Kind: Rook, Color: Black},
		"h8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "Kxd2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("e1"), To: sqn("d2")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

// --- Each piece kind ---

func TestParseSAN_PieceKinds(t *testing.T) {
	cases := []struct {
		name  string
		board map[string]Piece
		san   string
		want  Move
	}{
		{
			name: "Knight Nf3",
			board: map[string]Piece{
				"g1": {Kind: Knight, Color: White},
				"e1": {Kind: King, Color: White},
				"e8": {Kind: King, Color: Black},
			},
			san:  "Nf3",
			want: Move{From: sqn("g1"), To: sqn("f3")},
		},
		{
			name: "Bishop Bb5",
			board: map[string]Piece{
				"c4": {Kind: Bishop, Color: White},
				"e1": {Kind: King, Color: White},
				"e8": {Kind: King, Color: Black},
			},
			san:  "Bb5",
			want: Move{From: sqn("c4"), To: sqn("b5")},
		},
		{
			name: "Rook Rh4",
			board: map[string]Piece{
				"h1": {Kind: Rook, Color: White},
				"e1": {Kind: King, Color: White},
				"a8": {Kind: King, Color: Black},
			},
			san:  "Rh4",
			want: Move{From: sqn("h1"), To: sqn("h4")},
		},
		{
			name: "Queen Qd4",
			board: map[string]Piece{
				"d1": {Kind: Queen, Color: White},
				"e1": {Kind: King, Color: White},
				"a8": {Kind: King, Color: Black},
			},
			san:  "Qd4",
			want: Move{From: sqn("d1"), To: sqn("d4")},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			state := GameState{Board: sparseBoard(tc.board), ActiveColor: White}
			got, err := ParseSAN(state, tc.san)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Errorf("ParseSAN(%q) = %+v, want %+v", tc.san, got, tc.want)
			}
		})
	}
}

// --- Capture for non-King pieces ---

func TestParseSAN_PieceCapture(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"d3": {Kind: Knight, Color: White},
		"e5": {Kind: Pawn, Color: Black},
		"a1": {Kind: King, Color: White},
		"h8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "Nxe5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("d3"), To: sqn("e5")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

// --- No legal move error ---

func TestParseSAN_NoLegalMove(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e1": {Kind: King, Color: White},
		"h8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	_, err := ParseSAN(state, "Qd4")
	var nle NoLegalMoveError
	if !errors.As(err, &nle) {
		t.Errorf("ParseSAN err = %T(%v), want NoLegalMoveError", err, err)
	}
}

// --- Ambiguous move error ---

func TestParseSAN_AmbiguousMove(t *testing.T) {
	// Two White Rooks on a1 and h1 can both reach e1
	b := sparseBoard(map[string]Piece{
		"a1": {Kind: Rook, Color: White},
		"h1": {Kind: Rook, Color: White},
		"e8": {Kind: King, Color: White},
		"a8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	_, err := ParseSAN(state, "Re1")
	var ame AmbiguousMoveError
	if !errors.As(err, &ame) {
		t.Errorf("ParseSAN err = %T(%v), want AmbiguousMoveError", err, err)
	}
}

// --- File disambiguation ---

func TestParseSAN_FileDisambiguation(t *testing.T) {
	// White Rooks on a1 and h1; "Rae1" selects the a1 rook
	b := sparseBoard(map[string]Piece{
		"a1": {Kind: Rook, Color: White},
		"h1": {Kind: Rook, Color: White},
		"e8": {Kind: King, Color: White},
		"a8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "Rae1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("a1"), To: sqn("e1")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

// --- Rank disambiguation ---

func TestParseSAN_RankDisambiguation(t *testing.T) {
	// White Knights on b1 and b3 can both reach d2; "N1d2" selects b1
	b := sparseBoard(map[string]Piece{
		"b1": {Kind: Knight, Color: White},
		"b3": {Kind: Knight, Color: White},
		"e1": {Kind: King, Color: White},
		"h8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "N1d2")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("b1"), To: sqn("d2")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

// --- Pawn moves ---

func TestParseSAN_PawnSingleAdvance(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e2": {Kind: Pawn, Color: White},
		"e1": {Kind: King, Color: White},
		"e8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "e3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("e2"), To: sqn("e3")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

func TestParseSAN_PawnDoubleAdvance(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e2": {Kind: Pawn, Color: White},
		"e1": {Kind: King, Color: White},
		"e8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "e4")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("e2"), To: sqn("e4")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

func TestParseSAN_PawnCapture(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e4": {Kind: Pawn, Color: White},
		"d5": {Kind: Pawn, Color: Black},
		"e1": {Kind: King, Color: White},
		"e8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "exd5")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("e4"), To: sqn("d5")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

func TestParseSAN_PawnCheckMarkerStripped(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e2": {Kind: Pawn, Color: White},
		"e1": {Kind: King, Color: White},
		"e8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}
	want := Move{From: sqn("e2"), To: sqn("e3")}

	for _, input := range []string{"e3+", "e3#"} {
		got, err := ParseSAN(state, input)
		if err != nil {
			t.Fatalf("ParseSAN(%q) unexpected error: %v", input, err)
		}
		if got != want {
			t.Errorf("ParseSAN(%q) = %+v, want %+v", input, got, want)
		}
	}
}

func TestParseSAN_PawnSyntaxError(t *testing.T) {
	state := GameState{Board: NewStartingBoard(), ActiveColor: White}
	cases := []string{"e9", "xe4", "9e", "exd9"}
	for _, input := range cases {
		t.Run(input, func(t *testing.T) {
			_, err := ParseSAN(state, input)
			var se SyntaxError
			if !errors.As(err, &se) {
				t.Errorf("ParseSAN(%q) err = %T(%v), want SyntaxError", input, err, err)
			}
		})
	}
}

func TestParseSAN_PawnCaptureEmptySquare(t *testing.T) {
	b := sparseBoard(map[string]Piece{
		"e4": {Kind: Pawn, Color: White},
		"e1": {Kind: King, Color: White},
		"e8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	_, err := ParseSAN(state, "exd5")
	var nle NoLegalMoveError
	if !errors.As(err, &nle) {
		t.Errorf("ParseSAN err = %T(%v), want NoLegalMoveError", err, err)
	}
}

// --- Full-square disambiguation ---

func TestParseSAN_FullSquareDisambiguation(t *testing.T) {
	// White Queens on h4 and a1; both can reach e1; "Qh4e1" selects h4
	b := sparseBoard(map[string]Piece{
		"h4": {Kind: Queen, Color: White},
		"a1": {Kind: Queen, Color: White},
		"g1": {Kind: King, Color: White},
		"h8": {Kind: King, Color: Black},
	})
	state := GameState{Board: b, ActiveColor: White}

	got, err := ParseSAN(state, "Qh4e1")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	want := Move{From: sqn("h4"), To: sqn("e1")}
	if got != want {
		t.Errorf("ParseSAN = %+v, want %+v", got, want)
	}
}

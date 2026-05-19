package core

import "testing"

func TestComputeStatus_ReturnsOngoing(t *testing.T) {
	start := NewGame()
	afterMove := ApplyMove(start, Move{
		From: Square{File: 4, Rank: 1},
		To:   Square{File: 4, Rank: 3},
	})

	cases := []struct {
		name  string
		state GameState
	}{
		{"starting position", start},
		{"after e2-e4", afterMove},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := ComputeStatus(tc.state); got != Ongoing {
				t.Errorf("ComputeStatus = %d, want Ongoing (%d)", got, Ongoing)
			}
		})
	}
}

func TestGameStatus_DistinctValues(t *testing.T) {
	values := map[string]GameStatus{
		"Ongoing":   Ongoing,
		"Check":     Check,
		"Checkmate": Checkmate,
		"Stalemate": Stalemate,
	}

	seen := map[GameStatus]string{}
	for name, v := range values {
		if other, dup := seen[v]; dup {
			t.Errorf("%s and %s share value %d, must be distinct", name, other, v)
		}
		seen[v] = name
	}

	if len(seen) != 4 {
		t.Errorf("got %d distinct GameStatus values, want 4", len(seen))
	}
}

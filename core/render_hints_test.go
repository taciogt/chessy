package core_test

import (
	"testing"

	"github.com/taciogt/chessy/core"
)

func TestRenderHintsZeroValue(t *testing.T) {
	h := core.RenderHints{}
	if h.PreviewMove != nil {
		t.Error("expected nil PreviewMove")
	}
	if len(h.HighlightSquares) != 0 {
		t.Error("expected empty HighlightSquares")
	}
}

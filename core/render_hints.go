package core

// RenderHints carries optional overlay information for the Renderer.
// PreviewMove highlights a move about to be played; HighlightSquares
// marks squares for beginner/debug overlays.
type RenderHints struct {
	PreviewMove      *Move
	HighlightSquares []Square
}

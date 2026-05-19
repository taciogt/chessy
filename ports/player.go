// Package ports defines the contracts that adapters must satisfy.
// Both interfaces use core types so the game-loop orchestrator can compose
// Player + Renderer + core.GameState in a single signature.
package ports

import "github.com/taciogt/chessy/core"

// Player produces a Move for the active side from the given state.
// Implementations include terminal humans, Minimax agents, and UCI bridges.
type Player interface {
	SelectMove(state core.GameState) core.Move
}

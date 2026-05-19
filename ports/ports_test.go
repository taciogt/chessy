package ports_test

import (
	"testing"

	"github.com/taciogt/chessy/core"
	"github.com/taciogt/chessy/ports"
)

// noop implements both ports.Player and ports.Renderer with do-nothing bodies.
type noop struct{}

func (noop) SelectMove(core.GameState) core.Move { return core.Move{} }
func (noop) Render(core.GameState)                {}

// Compile-time assertions: interface contracts are satisfiable.
var (
	_ ports.Player   = noop{}
	_ ports.Renderer = noop{}
)

func TestPlayerAndRendererAreSatisfiable(t *testing.T) {
	var p ports.Player = noop{}
	var r ports.Renderer = noop{}

	_ = p.SelectMove(core.NewGame())
	r.Render(core.NewGame())
}

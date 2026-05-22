package ports

import "github.com/taciogt/chessy/core"

// Renderer displays a GameState. Implementations include a terminal UI, a web
// frontend, or a graphical client.
type Renderer interface {
	Render(state core.GameState, hints core.RenderHints)
}

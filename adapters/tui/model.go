package tui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/taciogt/chessy/core"
	"github.com/taciogt/chessy/ports"
)

// Compile-time assertion: Model satisfies the Renderer port.
var _ ports.Renderer = Model{}

var pieceSymbol = map[core.Color]map[core.PieceKind]string{
	core.White: {
		core.King:   "♔",
		core.Queen:  "♕",
		core.Rook:   "♖",
		core.Bishop: "♗",
		core.Knight: "♘",
		core.Pawn:   "♙",
	},
	core.Black: {
		core.King:   "♚",
		core.Queen:  "♛",
		core.Rook:   "♜",
		core.Bishop: "♝",
		core.Knight: "♞",
		core.Pawn:   "♟",
	},
}

// Model is a Bubble Tea model that renders a chess GameState.
type Model struct {
	state core.GameState
	hints core.RenderHints
}

// New creates a Model for the given GameState.
func New(state core.GameState) Model {
	return Model{state: state}
}

func (m Model) Init() tea.Cmd { return nil }

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

// View renders the board with rank/file labels and a turn indicator.
// Rank 8 is at the top; files a–h run left to right (white's perspective).
func (m Model) View() string {
	var sb strings.Builder

	for rank := 7; rank >= 0; rank-- {
		fmt.Fprintf(&sb, " %d  ", rank+1)
		for file := 0; file < 8; file++ {
			piece := m.state.Board[rank][file]
			if piece == nil {
				sb.WriteString("·  ")
			} else {
				fmt.Fprintf(&sb, "%s  ", pieceSymbol[piece.Color][piece.Kind])
			}
		}
		sb.WriteString("\n")
	}

	sb.WriteString("\n     ")
	for _, f := range "abcdefgh" {
		fmt.Fprintf(&sb, "%c  ", f)
	}
	sb.WriteString("\n\n")

	if m.state.ActiveColor == core.White {
		sb.WriteString("White to move\n")
	} else {
		sb.WriteString("Black to move\n")
	}

	sb.WriteString("\nPress q to quit\n")
	return sb.String()
}

// Render implements ports.Renderer by running a Bubble Tea program for state.
func (m Model) Render(state core.GameState, hints core.RenderHints) {
	m.state = state
	m.hints = hints
	p := tea.NewProgram(m)
	//nolint:errcheck
	p.Run()
}

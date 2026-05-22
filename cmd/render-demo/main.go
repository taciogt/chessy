package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/taciogt/chessy/adapters/tui"
	"github.com/taciogt/chessy/core"
)

func main() {
	p := tea.NewProgram(tui.New(core.NewGame()))
	if _, err := p.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

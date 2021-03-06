package main

import (
	"fmt"
	"os"

	"github.com/happyRip/Tic-Tac-Go/pkg/game"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	p := tea.NewProgram(game.InitialModel)

	err := p.Start()
	if err != nil {
		fmt.Printf("We've encountered an error: %v", err)
		os.Exit(1)
	}
}

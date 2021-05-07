package game

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.String() {

		case "ctrl+c", "q":
			m.view = "quitting"
			return m, tea.Quit

		case "up", "k", "w":
			m.cursor.Up()

		case "down", "j", "s":
			m.cursor.Down()

		case "left", "h", "a":
			m.cursor.Left()

		case "right", "l", "d":
			m.cursor.Right()

		case "enter", " ":
			switch m.view {

			case "game":
				if m.IsFieldEmpty() {
					m.PlaceToken()
				}

			case "winning", "draw":
				m.NewRound()
			}
		}
	}
	return m, nil
}

func (m Model) View() string {
	s := "\nTic Tac Go\n"

	switch m.view {

	case "game":
		s += fmt.Sprintf("\nScore\n%v : %v\n\n", m.score[0], m.score[1])

		for i, row := range m.board {

			render := row

			if m.cursor.Y == i {
				render[m.cursor.X] = "."
			}

			s += fmt.Sprintf(" %s | %s | %s\n", render[0], render[1], render[2])
			if i < 2 {
				s += "---+---+---\n"
			}
		}
		s += fmt.Sprintf("\nCursor position: %v, %v\n", m.cursor.X, m.cursor.Y)

	case "winning":
		s += fmt.Sprintf("\n%s has won!\n\nPress ENTER to start a new game.\n", m.CurrentPlayerName())

	case "draw":
		s += "\nNo winner this time.\n\nPress ENTER to start a new game.\n"

	case "quitting":
		return fmt.Sprintf("\nFinal score:\n%v : %v\n\n", m.score[0], m.score[1])
	}

	s += "\nPress q to quit.\n"

	return s
}

func (m *Model) ResetView() {
	m.view = "game"
}

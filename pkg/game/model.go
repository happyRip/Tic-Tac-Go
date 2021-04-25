package game

import (
	c "tic-tac-go/pkg/cursor"
	u "tic-tac-go/pkg/utility"
)

type Model struct {
	playerName [2]string
	score      [2]int
	board      [3][3]string
	cursor     c.Cursor
	isCircle   bool
	view       string
}

var InitialModel = Model{
	playerName: [2]string{"x", "o"},
	board:      u.ClearBoard(),
	view:       "game",
}

func (m *Model) HomeCursor() {
	m.cursor = c.Cursor{}

	for m.board[m.cursor.Y][m.cursor.X] != " " {
		if m.cursor.X < 2 {
			m.cursor.X++
		} else if m.cursor.Y < 2 {
			m.cursor.X = 0
			m.cursor.Y++
		} else {
			m.view = "draw"
			break
		}
	}
}

func (m Model) IsFieldEmpty() bool {
	return m.board[m.cursor.Y][m.cursor.X] == " "
}

func (m *Model) PlaceToken() {
	if m.isCircle {
		m.board[m.cursor.Y][m.cursor.X] = "o"
	} else {
		m.board[m.cursor.Y][m.cursor.X] = "x"
	}
	m.isCircle = !m.isCircle

	if m.HaveWon() {
		if m.isCircle {
			m.score[1]++
		} else {
			m.score[0]++
		}
	}

	m.HomeCursor()
}

func (m *Model) HaveWon() bool {
	for _, row := range m.board {
		if row[0] == row[1] && row[1] == row[2] && row[0] != " " {

			m.view = "winning"
			return true
		}
	}

	for i := 0; i < 3; i++ {
		if m.board[0][i] == m.board[1][i] && m.board[1][i] == m.board[2][i] && m.board[0][i] != " " {

			m.view = "winning"
			return true
		}
	}

	if m.board[0][0] == m.board[1][1] && m.board[1][1] == m.board[2][2] && m.board[0][0] != " " {

		m.view = "winning"
		return true
	}

	if m.board[0][2] == m.board[1][1] && m.board[1][1] == m.board[2][0] && m.board[1][1] != " " {

		m.view = "winning"
		return true
	}

	return false
}

func (m *Model) NewRound() {
	m.board = u.ClearBoard()
	m.isCircle = false
	m.HomeCursor()
	m.ResetView()
}

func (m Model) CurrentPlayerName() string {
	if !m.isCircle {
		return m.playerName[1]
	}
	return m.playerName[0]
}

package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func RunTUI() (tea.Model, error) {
	program := tea.NewProgram(initialModel())
	return program.Run()
}

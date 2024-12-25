package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	status string
}

func (m model) Init() tea.Cmd {
	return tea.SetWindowTitle("Termi")
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := m.status

	return s
}

func initialModel() model {
	return model{status: "ijinnamowowo"}
}

func RunTUI() (tea.Model, error) {
	program := tea.NewProgram(initialModel(), tea.WithAltScreen())
	return program.Run()
}

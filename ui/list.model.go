package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
}

func (m *Model) InitListPage() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (m *Model) UpdateListPage(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "ctrl+c":
			return m, tea.Quit
		case "ctrl+r":
			m.Cp = MAIN_PAGE
		case "l":
			fmt.Println(m.NoteService.NoteList())
		}
	case tea.WindowSizeMsg:
		with = msg.Width
	}
	return m, nil
}

func (m *Model) ViewListPage() string {
	return "list page"
}

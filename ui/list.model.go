package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
)

type List struct {
}

func (sm *StateManager) InitListPage() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (sm *StateManager) UpdateListPage(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "ctrl+c":
			return sm, tea.Quit
		case "ctrl+r":
			sm.Cp = MAIN_PAGE
		case "l":
			fmt.Println(sm.NoteService.NoteList())
		}
	case tea.WindowSizeMsg:
		with = msg.Width
	}
	return sm, nil
}

func (sm *StateManager) ViewListPage() string {
	return "list page"
}

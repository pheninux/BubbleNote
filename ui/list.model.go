package ui

import tea "github.com/charmbracelet/bubbletea"

type ListModel struct {
}

func (b *BaseModel) InitListPage() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (b *BaseModel) UpdateListPage(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q", "ctrl+c":
			return b, tea.Quit
		case "ctrl+r":
			b.Cp = MAIN_PAGE
		}
	case tea.WindowSizeMsg:
		with = msg.Width
	}
	return b, nil
}

func (b *BaseModel) ViewListPage() string {
	return "list page"
}

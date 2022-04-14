package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"strings"
)

type MainModel struct {
	selectedChoice int
}

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("90"))
	normalStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	with         = 200
)

func (b *BaseModel) InitMainPage() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (b *BaseModel) UpdateMainPage(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return b, tea.Quit
		case "tab":
			if b.selectedChoice == 0 {
				b.selectedChoice++
			} else {
				b.selectedChoice--
			}
		case "enter":
			switch b.selectedChoice {
			case 0:
				b.Cp = 1
			case 1:
				b.Cp = 2
			}
		}
	case tea.WindowSizeMsg:
		with = msg.Width
	}
	return b, nil
}

func (b *BaseModel) ViewMainPage() string {
	choicesTpl := strings.Builder{}
	choice := ""
	// block note choices : choice tpl
	for i, v := range choices {
		cursor := " "
		if b.selectedChoice == i {
			cursor = ">"
			choice = termenv.String(v).Foreground(termenv.EnvColorProfile().Color("121")).String()
		} else {
			choice = v
		}
		//if m.cp == 0 {
		//	if _, ok := m.selected[i]; ok {
		//		if m.choiseCursor == 0 {
		//			return m.notePageView()
		//		} else {
		//			return m.listPageView()
		//		}
		//	}
		//}
		choicesTpl.WriteString(fmt.Sprintf("%s %s", cursor, choice))
	}

	return lipgloss.JoinVertical(lipgloss.Center, b.titleView(), lipgloss.NewStyle().Align(lipgloss.Left).Render(choicesTpl.String()), b.mainHelpView())
}

func (b *BaseModel) mainHelpView() string {
	return fmt.Sprintf("\n\n total notes : %v \n tab: switch modes • q: exit", lipgloss.NewStyle().Foreground(lipgloss.Color("#0FE066")).Render("12"))
}

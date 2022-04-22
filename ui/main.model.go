package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"strconv"
	"strings"
)

type Main struct {
	selectedChoice int
	help           help.Model
	mainKeyMap     mainKeyMap
}

var (
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("90"))
	normalStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	with         = 200
)

type mainKeyMap struct {
	tab  key.Binding
	quit key.Binding
}

func (sm *StateManager) NewMainModel() Main {

	return Main{mainKeyMap: InitMainKeyMap(), help: help.New()}
}

func InitMainKeyMap() mainKeyMap {

	return mainKeyMap{
		tab: key.NewBinding(
			key.WithKeys("tab"),
			key.WithHelp("tab", "switch")),
		quit: key.NewBinding(
			key.WithKeys("ctrl+q"),
			key.WithHelp("ctrl+q", "exit programme")),
	}
}

func (sm *StateManager) InitMainPage() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (sm *StateManager) UpdateMainPage(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return sm, tea.Quit
		case "tab":
			if sm.Page.Main.selectedChoice == 0 {
				sm.Page.Main.selectedChoice++
			} else {
				sm.Page.Main.selectedChoice--
			}
		case "ctrl+z":

		case "enter":
			switch sm.Page.Main.selectedChoice {
			case 0:
				sm.Cp = 1
			case 1:
				sm.Cp = 2
			}
		}
	case tea.WindowSizeMsg:
		with = msg.Width
	}
	return sm, nil
}

func (sm *StateManager) ViewMainPage() string {

	choicesTpl := strings.Builder{}
	choice := ""
	// block note choices : choice tpl
	for i, v := range choices {
		cursor := " "
		if sm.Page.Main.selectedChoice == i {
			cursor = ">"
			choice = termenv.String(v).Foreground(termenv.EnvColorProfile().Color("121")).String()
		} else {
			choice = v
		}
		choicesTpl.WriteString(fmt.Sprintf("%s %s", cursor, choice))
	}

	return lipgloss.JoinVertical(lipgloss.Center, sm.titleView(), lipgloss.NewStyle().Align(lipgloss.Left).MarginTop(3).Render(choicesTpl.String()), sm.mainHelpView())

}

func (sm *StateManager) mainHelpView() string {
	_, t := sm.NoteService.NoteList()
	return fmt.Sprintf("\n\n total notes : %s \n\n\n %s", lipgloss.NewStyle().Foreground(lipgloss.Color("#0FE066")).Render(strconv.Itoa(t)),
		sm.Page.Main.help.ShortHelpView([]key.Binding{
			sm.Page.Main.mainKeyMap.quit,
			sm.Page.Main.mainKeyMap.tab,
		}))
}

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

func (m *Model) NewMainModel() Main {

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

func (m *Model) InitMainPage() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (m *Model) UpdateMainPage(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+q":
			return m, tea.Quit
		case "tab":
			if m.Page.Main.selectedChoice == 0 {
				m.Page.Main.selectedChoice++
			} else {
				m.Page.Main.selectedChoice--
			}
		case "ctrl+z":

		case "enter":
			switch m.Page.Main.selectedChoice {
			case 0:
				m.Cp = 1
			case 1:
				m.Cp = 2
			}
		}
	case tea.WindowSizeMsg:
		with = msg.Width
	}
	return m, nil
}

func (m *Model) ViewMainPage() string {

	choicesTpl := strings.Builder{}
	choice := ""
	// block note choices : choice tpl
	for i, v := range choices {
		cursor := " "
		if m.Page.Main.selectedChoice == i {
			cursor = ">"
			choice = termenv.String(v).Foreground(termenv.EnvColorProfile().Color("121")).String()
		} else {
			choice = v
		}
		choicesTpl.WriteString(fmt.Sprintf("%s %s", cursor, choice))
	}

	return lipgloss.JoinVertical(lipgloss.Center, m.titleView(), lipgloss.NewStyle().Align(lipgloss.Left).Render(choicesTpl.String()), m.mainHelpView())
}

func (m *Model) mainHelpView() string {
	_, t := m.NoteService.NoteList()
	return fmt.Sprintf("\n\n total notes : %s \n %s", lipgloss.NewStyle().Foreground(lipgloss.Color("#0FE066")).Render(strconv.Itoa(t)),
		m.Page.Main.help.ShortHelpView([]key.Binding{
			m.Page.Main.mainKeyMap.quit,
			m.Page.Main.mainKeyMap.tab,
		}))
}

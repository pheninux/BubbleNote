package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type bar struct {
	status   string
	msg      string
	encoding string
	version  string
}

type model struct {
	bar    bar
	with   int
	height int
}

func (m *model) Init() tea.Cmd {
	m.initModelBar()
	return func() tea.Msg {
		return nil
	}
}

func (m *model) initModelBar() {
	b := bar{
		status:   statusStyle.Render("STATUS"),
		encoding: encodingStyle.Render("UTF-8"),
		version:  versionStyle.Render("4.5.98"),
	}
	m.bar = b
}

func (m *model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.with, m.height = msg.Width, msg.Height
	}
	return m, nil
}

func (m *model) View() string {
	fmt.Sprintf(" w : %v .  h :  %v", m.with, m.height)
	ui := strings.Builder{}
	ui.WriteString(lipgloss.Place(m.with, m.height, lipgloss.Bottom, lipgloss.Bottom, lipgloss.JoinHorizontal(lipgloss.Center,
		m.bar.status,
		msgStyle.Width(m.with-(w(m.bar.status)+w(m.bar.version)+w(m.bar.encoding))).Render("EN COUR"),
		m.bar.encoding,
		m.bar.version)))
	return ui.String()
}

func main() {

	if err := tea.NewProgram(&model{}, tea.WithAltScreen()).Start(); err != nil {
		panic(err)
	}

}

package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var dockStyle = lipgloss.NewStyle()
var titleRender = lipgloss.NewStyle()
var contentRender = lipgloss.NewStyle().MarginTop(0)

type ui struct {
	t, r, b, l int
}
type model struct {
	ui
	with   int
	height int
}

func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.with, m.height = msg.Width, msg.Height
		m.ui.t, m.ui.r, m.ui.b, m.ui.l = dockStyle.GetMargin()

	}
	return m, nil
}

func (m model) View() string {
	return m.getTitle() + m.getContent()
}

func (m model) getTitle() string {

	return lipgloss.Place(m.with, 3,
		lipgloss.Center, lipgloss.Top,
		lipgloss.NewStyle().Render(fmt.Sprintf("with => %v  , height => %v ", m.with, m.height)),
		lipgloss.WithWhitespaceChars("#"), lipgloss.WithWhitespaceForeground(lipgloss.Color("67")))
}
func (m model) getHelp() string {
	return lipgloss.Place(m.with-m.ui.r-m.ui.l, m.height-13, lipgloss.Center, lipgloss.Bottom, fmt.Sprintf("q => quit \"/  s => save"), lipgloss.WithWhitespaceChars("#"), lipgloss.WithWhitespaceForeground(lipgloss.Color("12")))
}

func (m model) getContent() string {
	return lipgloss.Place(m.with-m.ui.r-m.ui.l, m.height-m.t-m.b, lipgloss.Center, lipgloss.Top, contentRender.Render(fmt.Sprintf("im the content")), lipgloss.WithWhitespaceChars("#"), lipgloss.WithWhitespaceForeground(lipgloss.Color("40")))

}
func main() {

	if err := tea.NewProgram(&model{}, tea.WithAltScreen()).Start(); err != nil {
		panic(err)
	}

}

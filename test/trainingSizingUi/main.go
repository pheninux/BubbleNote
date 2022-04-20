package main

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

//style bar
var (
	with          = 50
	height        = 30
	w             = lipgloss.Width
	h             = lipgloss.Height
	barStyleColor = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{
		Light: "#353533",
		Dark:  "#F2F2D9",
	}).Background(lipgloss.AdaptiveColor{
		Light: "#CACAC6",
		Dark:  "#DCDCD5",
	})
	textFg        = lipgloss.NewStyle().Foreground(lipgloss.Color("#FEFFFE")).Padding(0, 1)
	statusStyle   = textFg.Copy().Inherit(barStyleColor).Background(lipgloss.Color("#EA52BC"))
	msgStyle      = textFg.Copy().Inherit(barStyleColor)
	encodingStyle = textFg.Copy().Inherit(barStyleColor).Background(lipgloss.Color("#F5BAF7"))
	versionStyle  = textFg.Copy().Inherit(barStyleColor).Background(lipgloss.Color("#9152EA"))
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
		with, height = msg.Width, msg.Height
	}
	return m, nil
}

func (m *model) View() string {
	fmt.Sprintf(" w : %v .  h :  %v", with, height)
	ui := strings.Builder{}
	ui.WriteString(lipgloss.Place(with, height, lipgloss.Bottom, lipgloss.Bottom, lipgloss.JoinHorizontal(lipgloss.Center,
		m.bar.status,
		msgStyle.Width(with-(w(m.bar.status)+w(m.bar.version)+w(m.bar.encoding))).Render("EN COUR"),
		m.bar.encoding,
		m.bar.version)))
	return ui.String()
}

func main() {

	if err := tea.NewProgram(&model{}, tea.WithAltScreen()).Start(); err != nil {
		panic(err)
	}

}

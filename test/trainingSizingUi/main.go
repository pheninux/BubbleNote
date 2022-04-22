package main

import (
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

	//column := m.with / 3
	ui := strings.Builder{}
	ui.WriteString(lipgloss.Place(m.with, (m.height-m.height)+1, lipgloss.Center, lipgloss.Top, lipgloss.JoinHorizontal(lipgloss.Center,
		m.bar.status,
		msgStyle.Width(m.with-(w(m.bar.status)+w(m.bar.version)+w(m.bar.encoding))).Render("EN COUR"),
		m.bar.encoding,
		m.bar.version), lipgloss.WithWhitespaceChars("#"), lipgloss.WithWhitespaceForeground(lipgloss.Color("34"))))

	/*ui.WriteString(lipgloss.Place(m.with, (m.height-m.height)+1, lipgloss.Right,
	lipgloss.Center,
	lipgloss.NewStyle().MarginTop(5).Foreground(lipgloss.Color("57")).Render("coucou")))*/

	col1 := lipgloss.Place(5, 3,
		lipgloss.Left,
		lipgloss.Center,
		lipgloss.NewStyle().Background(lipgloss.Color("57")).Padding(0, 1).MarginRight(1).Render("Wikipédia est une encyclopédie universelle et multilingue créée par Jimmy Wales et Larry Sanger le 15 janvier 2001. Il s'agit d'une œuvre libre, c'est-à-dire que chacun est libre de la rediffuser. Gérée en wiki dans le site web wikipedia.org grâce au logiciel MediaWiki, elle \n"))

	col2 := lipgloss.Place(5, 12,
		lipgloss.Left,
		lipgloss.Center,
		lipgloss.NewStyle().Background(lipgloss.Color("24")).Padding(0, 1).MarginRight(1).Render("Wikipédia est une encyclopédie universelle et multilingue créée par Jimmy Wales et Larry Sanger le 15 janvier 2001. Il s'agit d'une œuvre libre, c'est-à-dire que chacun est libre de la rediffuser. Gérée en wiki dans le site web wikipedia.org grâce au logiciel MediaWiki, elle permet à tous les internautes d'écrire et de modifier des articles, ce qui lui"))

	col3 := lipgloss.Place(5, 12,
		lipgloss.Left,
		lipgloss.Center,
		lipgloss.NewStyle().Background(lipgloss.Color("10")).Padding(0, 1).MarginRight(1).Render("Wikipédia est une encyclopédie universelle et multilingue créée par Jimmy Wales et Larry Sanger le 15 janvier 2001. Il s'agit d'une œuvre libre, c'est-à-dire que chacun est libre de la rediffuser. Gérée en wiki dans le site web wikipedia.org grâce au logiciel MediaWiki, elle permet à tous les internautes d'écrire et de modifier des articles, ce qui lui"))
	ui.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, col1, col2, col3))
	return ui.String()
}

func main() {

	if err := tea.NewProgram(&model{}, tea.WithAltScreen()).Start(); err != nil {
		panic(err)
	}

}

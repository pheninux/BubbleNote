package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type NoteModel struct {
	Ti         textinput.Model
	noteKeyMap notKeyMap
	help       help.Model
}

type notKeyMap struct {
	previous key.Binding
	new      key.Binding
	save     key.Binding
	quit     key.Binding
}

func InitNoteKeyMap() notKeyMap {

	return notKeyMap{
		previous: key.NewBinding(
			key.WithKeys("ctrl+p"),
			key.WithHelp("ctrl+p", "previous")),
		new: key.NewBinding(
			key.WithKeys("ctrl+n"),
			key.WithHelp("ctrl+n", "new note")),
		save: key.NewBinding(
			key.WithKeys("ctrl+s"),
			key.WithHelp("ctrl+s", "save note")),
		quit: key.NewBinding(
			key.WithKeys("ctrl+q"),
			key.WithHelp("ctrl+q", "quit programme")),
	}
}

func (b *BaseModel) NewNoteModel() *NoteModel {
	ti := textinput.New()
	ti.Placeholder = "Write your note"
	ti.Focus()
	ti.BackgroundStyle = lipgloss.NewStyle().Background(lipgloss.Color("100"))
	ti.CursorStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	ti.CharLimit = 156
	ti.Width = 100

	return &NoteModel{Ti: ti, noteKeyMap: InitNoteKeyMap(), help: help.New()}
}

func (b *BaseModel) InitNotePage() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (b *BaseModel) UpdateNotePage(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, b.noteKeyMap.quit):
			return b, tea.Quit
		case key.Matches(msg, b.noteKeyMap.previous):
			b.Cp = MAIN_PAGE
		case key.Matches(msg, b.noteKeyMap.new):
		//todo impliment new note
		case key.Matches(msg, b.noteKeyMap.save):
			//todo save note
		}

	case tea.WindowSizeMsg:
		with = msg.Width
	}
	b.NoteModel.Ti, cmd = b.NoteModel.Ti.Update(msg)
	return b, cmd
}

func (b *BaseModel) ViewNotePage() string {
	ui := strings.Builder{}
	ui.WriteString(b.titleView())
	ui.WriteString(b.NoteModel.Ti.View())
	ui.WriteString(b.notehelpView())
	return lipgloss.NewStyle().Align(lipgloss.Center).Width(with - 2).Render(ui.String())
}

func (b *BaseModel) notehelpView() string {
	return fmt.Sprintf("\n\n total notes : %v \n\n %s", lipgloss.NewStyle().Foreground(lipgloss.Color("#0FE066")).Render("12"),
		b.help.ShortHelpView([]key.Binding{
			b.noteKeyMap.previous,
			b.noteKeyMap.new,
			b.noteKeyMap.save,
			b.noteKeyMap.quit,
		}))
}

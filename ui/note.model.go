package ui

import (
	"BubbleNote/pkg/model"
	"fmt"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
	"time"
)

type Note struct {
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

func (sm *StateManager) NewNoteModel() Note {
	ti := textinput.New()
	ti.Placeholder = "Write your note"
	ti.Focus()
	ti.BackgroundStyle = lipgloss.NewStyle().Background(lipgloss.Color("100"))
	ti.CursorStyle = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	ti.TextStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	ti.CharLimit = 156
	ti.Width = 100

	return Note{Ti: ti, noteKeyMap: InitNoteKeyMap(), help: help.New()}
}

func (sm *StateManager) InitNotePage() tea.Cmd {
	return func() tea.Msg {
		return nil
	}
}

func (sm *StateManager) UpdateNotePage(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, sm.Page.Note.noteKeyMap.quit):
			return sm, tea.Quit
		case key.Matches(msg, sm.Page.Note.noteKeyMap.previous):
			sm.Cp = MAIN_PAGE
		case key.Matches(msg, sm.Page.Note.noteKeyMap.new):
		//todo impliment new note
		case key.Matches(msg, sm.Page.Note.noteKeyMap.save):
			n := model.Note{}
			n.Content = sm.Page.Note.Ti.Value()
			n.CreatedAt = time.Now()
			if err := sm.NoteService.SaveNote(n); err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("success")
			}
		}

	case tea.WindowSizeMsg:
		with = msg.Width
	}
	sm.Page.Note.Ti, cmd = sm.Page.Note.Ti.Update(msg)
	return sm, cmd
}

func (sm *StateManager) ViewNotePage() string {

	ui := strings.Builder{}
	ui.WriteString(sm.titleView())
	ui.WriteString(sm.Page.Note.Ti.View())
	ui.WriteString(sm.notehelpView())
	return lipgloss.NewStyle().Align(lipgloss.Center).Width(with - 2).Render(ui.String())
}

func (sm *StateManager) notehelpView() string {
	return fmt.Sprintf("\n\n total notes : %v \n\n %s", lipgloss.NewStyle().Foreground(lipgloss.Color("#0FE066")).Render("12"),
		sm.Page.Note.help.ShortHelpView([]key.Binding{
			sm.Page.Note.noteKeyMap.previous,
			sm.Page.Note.noteKeyMap.new,
			sm.Page.Note.noteKeyMap.save,
			sm.Page.Note.noteKeyMap.quit,
		}))
}

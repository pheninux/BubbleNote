package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	Page        Pages
	Cp          int // page cursor
	NoteService interface {
		SaveNote(note string) error
		NoteList() string
	}
}

type Pages struct {
	Main Main
	Note Note
	List List
}

var (
	choices = []string{" Creates a new note \n", " List all notes"}
)

const (
	MAIN_PAGE = iota
	NOTE_PAGE
	LIST_PAGE
)

func (m *Model) Init() tea.Cmd {

	return func() tea.Msg {
		return nil
	}
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch m.Cp {
	case MAIN_PAGE:
		return m.UpdateMainPage(msg)
	case NOTE_PAGE:
		return m.UpdateNotePage(msg)
	case LIST_PAGE:
		return m.UpdateListPage(msg)
	}
	return m, nil
}

func (m *Model) View() string {
	switch m.Cp {
	case MAIN_PAGE:
		return m.ViewMainPage()
	case NOTE_PAGE:
		return m.ViewNotePage()
	case LIST_PAGE:
		return m.ViewListPage()
	case 5:
		return "error when saving note"
	case 6:
		return "note saved success"
	default:
		return "other pages"
	}
}

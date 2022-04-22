package ui

import (
	"BubbleNote/pkg/model"
	tea "github.com/charmbracelet/bubbletea"
)

type StateManager struct {
	Page        Pages
	Cp          int // page cursor
	NoteService interface {
		SaveNote(n model.Note) error
		NoteList() (string, int)
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

func (sm *StateManager) Init() tea.Cmd {

	return func() tea.Msg {
		return nil
	}
}

func (sm *StateManager) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch sm.Cp {
	case MAIN_PAGE:
		return sm.UpdateMainPage(msg)
	case NOTE_PAGE:
		return sm.UpdateNotePage(msg)
	case LIST_PAGE:
		return sm.UpdateListPage(msg)
	}
	return sm, nil
}

func (sm *StateManager) View() string {
	switch sm.Cp {
	case MAIN_PAGE:
		return sm.ViewMainPage()
	case NOTE_PAGE:
		return sm.ViewNotePage()
	case LIST_PAGE:
		return sm.ViewListPage()
	case 5:
		return "error when saving note"
	case 6:
		return "note saved success"
	default:
		return "other pages"
	}
}

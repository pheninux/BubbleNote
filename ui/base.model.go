package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type BaseModel struct {
	*MainModel
	*NoteModel
	*ListModel
	Cp int // page cursor
}

var (
	choices = []string{" Creates a new note \n", " List all notes"}
)

const (
	MAIN_PAGE = iota
	NOTE_PAGE
	LIST_PAGE
)

func (b *BaseModel) Init() tea.Cmd {

	return func() tea.Msg {
		return nil
	}
}

func (b *BaseModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch b.Cp {
	case MAIN_PAGE:
		return b.UpdateMainPage(msg)
	case NOTE_PAGE:
		return b.UpdateNotePage(msg)
	case LIST_PAGE:
		return b.UpdateListPage(msg)
	}
	return b, nil
}

func (b *BaseModel) View() string {
	switch b.Cp {
	case MAIN_PAGE:
		return b.ViewMainPage()
	case NOTE_PAGE:
		return b.ViewNotePage()
	case LIST_PAGE:
		return b.ViewListPage()
	default:
		return "other pages"
	}
}

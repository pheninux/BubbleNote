package main

import (
	"BubbleNote/ui"
	tea "github.com/charmbracelet/bubbletea"
)

type application struct {
	model ui.BaseModel
}

func main() {
	bm := ui.BaseModel{}
	app := application{ui.BaseModel{
		MainModel: &ui.MainModel{},
		NoteModel: bm.NewNoteModel(),
		ListModel: &ui.ListModel{},
		Cp:        0,
	}}
	if err := tea.NewProgram(&app.model).Start(); err != nil {
		panic(err)
	}
}

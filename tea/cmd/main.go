package main

import (
	"BubbleNote/pkg/dao"
	"BubbleNote/ui"
	badger "badger-3.2103.2"
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type application struct {
	model   ui.BaseModel
	noteRep dao.NoteRepo
}

func main() {
	bm := ui.BaseModel{}
	app := application{model: ui.BaseModel{
		MainModel: &ui.MainModel{},
		NoteModel: bm.NewNoteModel(),
		ListModel: &ui.ListModel{},
		Cp:        0,
	}, noteRep: dao.NoteRepo{Db: openDB()}}
	if err := tea.NewProgram(&app.model).Start(); err != nil {
		panic(err)
	}
}

func openDB() *badger.DB {
	db, err := badger.Open(badger.DefaultOptions("./db/"))
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	return db
}

package main

import (
	"BubbleNote/pkg/dao"
	"BubbleNote/ui"
	bolt "github.com/boltdb/bolt"
	tea "github.com/charmbracelet/bubbletea"
	badger "github.com/dgraph-io/badger/v3"
	"log"
	"os"
)

func main() {

	db, err := openBoltDB("./db/bubble.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	m := ui.StateManager{}
	m = ui.StateManager{
		Page: ui.Pages{
			Main: m.NewMainModel(),
			Note: m.NewNoteModel(),
			List: ui.List{},
		},
		Cp:          0,
		NoteService: &dao.NoteRepo{Db: db},
	}
	if err := tea.NewProgram(&m, tea.WithAltScreen()).Start(); err != nil {

	}
}

func openBadgerDB(path string) (*badger.DB, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.MkdirAll(path, 0755)
	}
	opts := badger.DefaultOptions(path)
	opts.Dir = path
	opts.Logger = nil
	opts.ValueDir = path
	opts.SyncWrites = false
	opts.ValueThreshold = 256
	opts.CompactL0OnClose = true
	db, err := badger.Open(opts)
	if err != nil {
		log.Println("badger open failed", "path", path, "err", err)
		return nil, err
	}
	return db, nil
}

func openBoltDB(path string) (*bolt.DB, error) {

	db, err := bolt.Open(path, 0600, nil)
	if err := createBucket(db); err != nil {
		return db, err
	}
	return db, err
}

func createBucket(db *bolt.DB) error {
	err := db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("notes"))
		return err
	})
	return err
}

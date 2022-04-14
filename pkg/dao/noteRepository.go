package dao

import (
	badger "badger-3.2103.2"
)

type NoteRepo struct {
	Db *badger.DB
}

func (nr *NoteRepo) SaveNote() {

}

package dao

import (
	"BubbleNote/pkg/model"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
)

type NoteRepo struct {
	Db *bolt.DB
}

func (nr *NoteRepo) SaveNote(n model.Note) error {

	err := nr.Db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("notes"))
		id, _ := bucket.NextSequence()
		n.ID = int(id)
		b, err := json.Marshal(n)
		if err != nil {
			fmt.Println(err)
		}
		err = bucket.Put([]byte(string(n.ID)), b)
		return err
	})
	return err
}

func (nr *NoteRepo) NoteList() (l []model.Note, t int, err error) {
	total := 0

	nr.Db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("notes"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			n := model.Note{}
			err = json.Unmarshal([]byte(v), &n)
			l = append(l, n)
			total++
		}
		return nil
	})
	return l, total, err
}

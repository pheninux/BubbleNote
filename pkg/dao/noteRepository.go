package dao

import (
	"fmt"
	"github.com/boltdb/bolt"
	"strings"
)

type NoteRepo struct {
	//Db  *badger.DB
	Db  *bolt.DB
	key int
}

func (nr *NoteRepo) SaveNote(note string) error {

	//err := nr.Db.Update(func(txn *badger.Txn) error {
	//	nr.key++
	//	err := txn.Set([]byte(string(nr.key)), []byte(note))
	//	return err
	//})
	//return err

	err := nr.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("notes"))
		key, _ := b.NextSequence()
		err := b.Put([]byte(string(key)), []byte(note))
		return err
	})
	return err
}

func (nr *NoteRepo) NoteList() string {
	//result := strings.Builder{}
	//err := nr.Db.View(func(txn *badger.Txn) error {
	//	opts := badger.DefaultIteratorOptions
	//	opts.PrefetchSize = 10
	//	it := txn.NewIterator(opts)
	//	defer it.Close()
	//	for it.Rewind(); it.Valid(); it.Next() {
	//		item := it.Item()
	//		k := item.Key()
	//		err := item.Value(func(v []byte) error {
	//			result.WriteString(fmt.Sprintf("key=%s, value=%s\n", k, v))
	//			return nil
	//		})
	//		if err != nil {
	//			return err
	//		}
	//	}
	//	return nil
	//})
	//if err != nil {
	//	result.WriteString(fmt.Sprintf("Failed to iterator keys and values from the cache."))
	//}
	//return result.String()
	result := strings.Builder{}
	nr.Db.View(func(tx *bolt.Tx) error {
		// Assume bucket exists and has keys
		b := tx.Bucket([]byte("notes"))

		c := b.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			result.WriteString(fmt.Sprintf("key=%s, value=%s\n", k, v))
		}

		return nil
	})
	return result.String()
}

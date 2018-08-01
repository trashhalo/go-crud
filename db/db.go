package db

import (
	"bytes"
	"encoding/gob"
	"github.com/dgraph-io/badger"
	"github.com/satori/go.uuid"
)

var db *badger.DB

type Todo struct {
	Key  string
	Text string
	Done bool
}

func CreateTodo(text string) error {
	uuidKey := uuid.NewV4()
	key := uuidKey.Bytes()
	todo := Todo{
		Key:  uuidKey.String(),
		Text: text,
		Done: false,
	}
	var value bytes.Buffer
	enc := gob.NewEncoder(&value)
	err := enc.Encode(todo)
	if err != nil {
		return err
	}
	return db.Update(func(txn *badger.Txn) error {
		return txn.Set(key, value.Bytes())
	})
}

func UpdateTodo(keyStr string, done bool) error {
	uuidKey, err := uuid.FromString(keyStr)
	if err != nil {
		return err
	}
	return db.Update(func(txn *badger.Txn) error {
		key := uuidKey.Bytes()
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		v, err := item.Value()
		if err != nil {
			return err
		}
		dec := gob.NewDecoder(bytes.NewBuffer(v))
		var todo Todo
		err = dec.Decode(&todo)
		if err != nil {
			return err
		}
		todo.Done = done
		var value bytes.Buffer
		enc := gob.NewEncoder(&value)
		err = enc.Encode(todo)
		if err != nil {
			return err
		}
		return txn.Set(key, value.Bytes())
	})
}

func ListTodos() ([]Todo, error) {
	var todos []Todo
	err := db.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		defer it.Close()
		for it.Rewind(); it.Valid(); it.Next() {
			item := it.Item()
			v, err := item.Value()
			if err != nil {
				return err
			}
			dec := gob.NewDecoder(bytes.NewBuffer(v))
			var todo Todo
			err = dec.Decode(&todo)
			if err != nil {
				return err
			}
			todos = append(todos, todo)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func OpenDB() error {
	var err error
	opts := badger.DefaultOptions
	opts.Dir = "/tmp/badger"
	opts.ValueDir = "/tmp/badger"
	db, err = badger.Open(opts)
	if err != nil {
		return err
	}
	return nil
}

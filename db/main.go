package db

import (
	"encoding/json"
	"os"
)

type db[T any] struct {
	entries map[string]map[string]interface{}
}

func newDb[T any](filepath string) (*db[T], error) {
	byt, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	instances := []map[string]interface{}{}

	err = json.Unmarshal(byt, &instances)

	if err != nil {
		return nil, err
	}

	entries := map[string]map[string]interface{}{}

	for _, ins := range instances {
		if id, ok := ins["id"]; ok {
			if idStr, isStr := id.(string); isStr {
				entries[idStr] = ins
			}
		}
	}

	return &db[T]{
		entries: entries,
	}, nil
}

func (db *db[T]) GetEntry(id string) (*T, error) {

	inter := db.entries[id]

	byt, err := json.Marshal(&inter)

	if err != nil {
		return nil, err
	}

	var entry T

	err = json.Unmarshal(byt, &entry)

	if err != nil {
		return nil, err
	}

	return &entry, err
}

func (db *db[T]) GetEntryBytes(id string) ([]byte, error) {
	inter := db.entries[id]

	byt, err := json.Marshal(&inter)

	if err != nil {
		return nil, err
	}

	return byt, err
}

package db

import (
	"kc.com/kc/types"
)

var ItemDb *db[types.Item]

func NewItemDb() error {
	var err error
	ItemDb, err = newDb[types.Item]("./db/item.json")

	if err != nil {
		return err
	}

	return nil
}

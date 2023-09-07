package items

import (
	"kc.com/kc/db"
	"kc.com/kc/types"
)

var ItemSpawner = &Item{}

type Item struct {
}

func Create(id string) (bool, types.IItem) {
	item, err := db.ItemDb.GetEntry(id)

	if err != nil {
		return false, nil
	}

	switch item.Class {
	case types.ItemClassHarvest:
		return true, NewHarvest(item)
	case types.ItemClassSeed:
		return true, NewSeed(item)
	}

	return false, nil
}

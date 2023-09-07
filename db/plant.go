package db

import "kc.com/kc/types"

var PlantDb *db[types.Plant]

func NewPlantDb() error {
	var err error
	PlantDb, err = newDb[types.Plant]("./db/plant.json")

	if err != nil {
		return err
	}

	return nil
}

package spawners

import (
	"fmt"
	"time"

	"kc.com/kc/db"
	"kc.com/kc/managers"
	"kc.com/kc/objects"
)

var ObjectSpawner = &Object{}

type Object struct {
}

func spawnPlant(objectType string, objectId string, xPos float64, yPos float64) error {

	id := fmt.Sprintf("object_%s_%s_%d", objectType, objectId, time.Now().UnixNano())

	p, err := db.PlantDb.GetEntry(objectId)

	if err != nil {
		return err
	}

	obj, err := objects.NewObject(id, p.ImagePath, xPos, yPos, 32, 32)

	if err != nil {
		return err
	}

	plant := objects.NewPlant(obj, p.Lifecycle, p.GrowRate, p.Bountiful, p.MultiHarvest, p.Fertile)

	if err != nil {
		return err
	}

	managers.ObjectManager.Register(id, plant)
	managers.CollisionManager.AppendCollidable(plant)

	return nil
}

func (o *Object) Spawn(objectType string, objectId string, xPos float64, yPos float64) error {
	switch objectType {
	case "plant":
		err := spawnPlant(objectType, objectId, xPos, yPos)
		if err != nil {
			return err
		}
	}

	return nil
}

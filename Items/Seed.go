package items

import (
	"encoding/json"

	"kc.com/kc/spawners"
	"kc.com/kc/types"
)

type SeedData struct {
	PlantId string
}

type Seed struct {
	*types.Item
	*SeedData
}

func (s *Seed) Use(xPos float64, yPos float64) bool {
	spawners.ObjectSpawner.Spawn("plant", s.SeedData.PlantId, xPos, yPos)

	return true
}

func NewSeed(item *types.Item) *Seed {

	data := &SeedData{}

	json.Unmarshal(item.ClassData, &data)

	return &Seed{
		item,
		data,
	}
}

package items

import (
	"encoding/json"

	"kc.com/kc/types"
)

type HarvestData struct {
}

type Harvest struct {
	*types.Item
}

func NewHarvest(item *types.Item) *Harvest {

	data := &HarvestData{}

	json.Unmarshal(item.ClassData, &data)

	return &Harvest{
		item,
	}
}

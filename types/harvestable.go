package types

type Harvestable interface {
	Harvest() ([]*InventoryItem, bool)
}

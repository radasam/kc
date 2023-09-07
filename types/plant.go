package types

type Plant struct {
	GlobalId     string
	ObjectId     string `db:"id"`
	ImagePath    string
	Lifecycle    []int
	GrowRate     []int
	Bountiful    float64
	MultiHarvest float64
	Fertile      float64
}

package objects

import (
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"kc.com/kc/managers"
	"kc.com/kc/types"
)

var LIGHT_LEVEL_PENALTY = []float64{0.01, 0.05, 0.2, 0.6, 0.9, 1}

type Plant struct {
	*Object
	ticks          float64
	lifecycleState int
	state          int
	lifecycle      []int
	growRate       []int
	bountiful      float64
	multiHarvest   float64
	hardy          float64
	nutritious     float64
	vigorous       float64
	fertile        float64
}

func (p *Plant) GetRad() float64 {
	return 16
}

func (p *Plant) IsSolid() bool {
	return false
}

func (p *Plant) Despawn() {
	managers.CollisionManager.RemoveCollideable(p)
	managers.ObjectManager.Despawn(p.Object.GetID())
}

func (p *Plant) GetVerts() []types.Vert {
	return p.Object.GetVerts()
}

func (p *Plant) OnCollision(t types.CollisionType, c types.Collideable) {
}

func (p *Plant) RollBountiful() int {
	roll := rand.Float64()
	return int(math.Ceil(roll / p.bountiful))
}

func (p *Plant) RollMultiHarvest() bool {
	roll := rand.Float64()
	return roll > p.multiHarvest
}

func (p *Plant) RollFertile() bool {
	roll := rand.Float64()
	return roll > p.fertile
}

func (p *Plant) Harvest() ([]*types.InventoryItem, bool) {
	items := []*types.InventoryItem{}
	if p.state == p.lifecycle[4] {
		multiHarvest := p.RollMultiHarvest()

		if multiHarvest {
			p.state = p.lifecycle[3]
			p.Object.SetState(p.lifecycle[3])
		} else {
			p.state += 1
			p.Object.IncrementState()
		}

		items = append(items, &types.InventoryItem{ItemId: "tomato", Count: p.RollBountiful(), Weight: 5})

		seed := p.RollFertile()

		if seed {
			items = append(items, &types.InventoryItem{ItemId: "tomato-seed", Count: 1, Weight: 1})
		}

		return items, true

	}

	return nil, false

}

func (p *Plant) Draw(screen *ebiten.Image) {
	p.ticks += 1 * LIGHT_LEVEL_PENALTY[managers.WeatherManager.GetWeather()]
	if p.ticks > float64(p.growRate[p.lifecycleState]) && p.state != p.lifecycle[4] && p.state < p.lifecycle[5] {
		if p.lifecycle[p.lifecycleState] < p.state {
			p.lifecycleState += 1
		}
		p.state += 1
		p.Object.IncrementState()
		p.ticks = 0
	}

	if p.state == p.lifecycle[5] {
		managers.ObjectManager.Despawn(p.GetID())
	} else {
		p.Object.Draw(screen)
	}
}

func NewPlant(object *Object, lifecycle []int, growRate []int, bountiful float64, multiHarvest float64, fertile float64) *Plant {
	return &Plant{Object: object, state: 0, lifecycle: lifecycle, growRate: growRate, bountiful: bountiful, multiHarvest: multiHarvest, fertile: fertile}
}

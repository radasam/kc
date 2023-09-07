package characters

import (
	"fmt"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"kc.com/kc/bases"
	"kc.com/kc/events"
	"kc.com/kc/managers"
	"kc.com/kc/types"
)

type Farmer struct {
	initialised bool
	id          string
	xPos        float64
	yPos        float64
	sprite      types.Sprite
	w           float64
	h           float64
	bases.Inventory
	bases.Infobox
}

func (f *Farmer) Draw(screen *ebiten.Image) {
	f.sprite.Draw(screen, f.xPos-f.w/2, f.yPos-f.h/2)
	f.Infobox.Draw(screen, f.xPos, f.yPos)
	if f.Inventory.Show {
		f.Inventory.Draw(screen, f.xPos+f.w/2, f.yPos-f.h/2)
	}

}

func (f *Farmer) Despawn() {
	managers.CollisionManager.RemoveCollideable(f)
	managers.CharacterManager.Despawn(f.id)
}

func (f *Farmer) Move(x float64, y float64) {
	f.xPos += x
	f.yPos += y
	managers.CollisionManager.UpdateCollidable(f)
	ct, d := managers.CollisionManager.CheckCollisions(f)

	if ct != types.NoCollision {
		if ct == types.HardCollision {
			if d.IsSolid() {
				f.xPos -= x
				f.yPos -= y
				return
			} else if h, ok := d.(types.Harvestable); ok {
				items, success := h.Harvest()
				if success {
					for _, item := range items {
						f.Inventory.AddItem(item.ItemId, item.Count, item.Weight)
						ev, err := events.NewInventoryEvent(item.ItemId, item.Count)
						if err != nil {
							println(err)
						}
						if ev == nil {
							println("is nil")
						} else {
							f.Infobox.AddEvent(ev)
						}
					}

				}
			}
		}

		if ct == types.SoftCollision {
			if h, ok := d.(types.Harvestable); ok {
				items, success := h.Harvest()
				if success {
					for _, item := range items {
						f.Inventory.AddItem(item.ItemId, item.Count, item.Weight)
						ev, err := events.NewInventoryEvent(item.ItemId, item.Count)
						if err != nil {
							println(err)
						}
						if ev == nil {
							println("is nil")
						} else {
							f.Infobox.AddEvent(ev)
						}
					}

				}
			}
		}
	}
}

func (f *Farmer) Init() {
	if !f.initialised {
		managers.KeypressManager.Bind("test", ebiten.KeyW, func() { f.Move(0, -1) })
		managers.KeypressManager.Bind("test", ebiten.KeyS, func() { f.Move(0, 1) })
		managers.KeypressManager.Bind("test", ebiten.KeyA, func() { f.Move(-1, 0) })
		managers.KeypressManager.Bind("test", ebiten.KeyD, func() { f.Move(1, 0) })
		managers.KeypressManager.BindPress("inv", ebiten.KeyB, f.Inventory.Toggle)
		managers.KeypressManager.BindPress("use", ebiten.KeyU, func() { f.Inventory.UseItem(f.xPos, f.yPos) })
		f.initialised = true
	}

}

func (f *Farmer) GetPos() []float64 {
	return []float64{f.xPos, f.yPos}
}

func (f *Farmer) GetRad() float64 {
	return 64
}

func (f *Farmer) GetDim() []float64 {
	return []float64{f.w, f.h}
}

func (f *Farmer) GetID() string {
	return f.id
}

func (f *Farmer) IsSolid() bool {
	return false
}

func (f *Farmer) GetVerts() []types.Vert {
	verts := []types.Vert{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			verts = append(verts, types.Vert{X: f.xPos + (f.w * float64(i)) - f.w/2, Y: f.yPos + (f.h * float64(j)) - f.h/2, Id: fmt.Sprintf("%s-%d", f.id, (i*2)+(j))})
		}
	}

	return verts
}

func (f *Farmer) OnCollision(t types.CollisionType, c types.Collideable) {

}

func SpawnFarmer() error {
	path, err := filepath.Abs("./farmer-sheet.png")

	if err != nil {
		return err
	}

	sprite, err := NewDirectionalSprite(path, 50, 50, 32, 32)

	if err != nil {
		return err
	}

	farmer := &Farmer{sprite: sprite, xPos: 50, yPos: 50, id: "farmer", w: 32, h: 32, Inventory: *bases.NewInventory(1000), Infobox: *bases.NewInfoBox()}

	managers.CharacterManager.Register("farmer", farmer)
	managers.CollisionManager.AppendCollidable(farmer)

	farmer.Inventory.AddItem("tomato-seed", 1, 1)

	return nil
}

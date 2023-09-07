package objects

import (
	"kc.com/kc/managers"
	"kc.com/kc/types"
)

type Wall struct {
	*Object
}

func (w *Wall) GetRad() float64 {
	return 16
}

func (w *Wall) IsSolid() bool {
	return true
}

func (w *Wall) OnCollision(t types.CollisionType, c types.Collideable) {

}

func SpawnWall(id string, filepath string, xPos float64, yPos float64) error {
	obj, err := NewObject(id, filepath, xPos, yPos, 8, 32)

	if err != nil {
		return err
	}

	wall := &Wall{Object: obj}

	managers.ObjectManager.Register(id, wall)
	managers.CollisionManager.AppendCollidable(wall)

	return nil
}

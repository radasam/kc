package managers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"kc.com/kc/types"
)

var ObjectManager *Object

type Object struct {
	objMap    map[string]types.Object
	initQueue []func()
}

func (om *Object) OnUpdate() {
	for _, v := range om.initQueue {
		v()
	}

	om.initQueue = []func(){}
}

func (om *Object) Draw(screen *ebiten.Image) {
	for _, c := range om.objMap {
		c.Draw(screen)
	}
}

func (om *Object) Register(id string, obj types.Object) {
	om.initQueue = append(om.initQueue, obj.Init)
	om.objMap[id] = obj
}

func (om *Object) Despawn(id string) {
	delete(om.objMap, id)
}

func NewObjectManager() {
	ObjectManager = &Object{
		objMap:    map[string]types.Object{},
		initQueue: []func(){},
	}
}

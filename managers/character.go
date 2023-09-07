package managers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"kc.com/kc/types"
)

var CharacterManager *Character

type Character struct {
	charMap   map[string]types.Character
	initQueue []func()
}

func (cm *Character) OnUpdate() {
	for _, v := range cm.initQueue {
		v()
	}

	cm.initQueue = []func(){}
}

func (cm *Character) Draw(screen *ebiten.Image) {
	for _, c := range cm.charMap {
		c.Draw(screen)
	}
}

func (cm *Character) Register(id string, char types.Character) {
	cm.initQueue = append(cm.initQueue, char.Init)
	cm.charMap[id] = char
}

func (cm *Character) Despawn(id string) {
	delete(cm.charMap, id)
}

func NewCharacterManager() {
	CharacterManager = &Character{
		charMap:   map[string]types.Character{},
		initQueue: []func(){},
	}
}

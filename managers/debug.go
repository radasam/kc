package managers

import (
	"github.com/hajimehoshi/ebiten/v2"
	"kc.com/kc/types"
)

var DebugManager *Debug

type Debug struct {
	drawMap map[string]types.Drawable
}

func (d *Debug) AddDraw(id string, drawable types.Drawable) {
	d.drawMap[id] = drawable
}

func (d *Debug) RemoveDraw(id string) {
	delete(d.drawMap, id)
}

func (d *Debug) Draw(screen *ebiten.Image) {
	for _, d2 := range d.drawMap {
		d2.Draw(screen)
	}
}

func NewDebugManager() {
	DebugManager = &Debug{
		drawMap: map[string]types.Drawable{},
	}
}

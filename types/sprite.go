package types

import "github.com/hajimehoshi/ebiten/v2"

type Sprite interface {
	Draw(screen *ebiten.Image, x float64, y float64)
}

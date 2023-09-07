package types

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	Draw(screen *ebiten.Image)
	Init()
}

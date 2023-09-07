package types

import "github.com/hajimehoshi/ebiten/v2"

type Character interface {
	Draw(screen *ebiten.Image)
	Init()
}

package characters

import (
	"bytes"
	"image"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Sprite struct {
	image  *ebiten.Image
	width  float64
	height float64
}

func (s *Sprite) Draw(screen *ebiten.Image, x float64, y float64, col int, row int, flipX bool) {
	op := &ebiten.DrawImageOptions{}

	if flipX {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(x+s.width, y)
	} else {
		op.GeoM.Translate(x, y)
	}
	screen.DrawImage(ebiten.NewImageFromImage(s.image.SubImage(image.Rect(0+(col*32), 0+(row*32), 32+(col*32), 32+(row*32)))), op)
}

func NewSprite(filepath string, width float64, height float64) (*Sprite, error) {
	r, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(r))
	if err != nil {
		return nil, err
	}

	return &Sprite{image: ebiten.NewImageFromImage(img), width: width, height: height}, nil
}

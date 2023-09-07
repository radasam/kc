package images

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type SpriteSheet struct {
	sheet *ebiten.Image
	h     int
	w     int
	row   int
}

func (s *SpriteSheet) GetImage() *ebiten.Image {
	return ebiten.NewImageFromImage(s.sheet.SubImage(image.Rect(0, 0+(s.row*s.h), s.w, s.h+(s.row*s.h))))
}

func (s *SpriteSheet) SetRow(row int) {
	s.row = row
}

func NewSpriteSheet(image *ebiten.Image, w int, h int, initRow int) *SpriteSheet {
	return &SpriteSheet{
		sheet: image,
		w:     w,
		h:     h,
		row:   initRow,
	}
}

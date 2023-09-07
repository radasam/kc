package helpers

import (
	"bytes"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func ImagefromFile(filepath string) (*ebiten.Image, error) {
	r, err := os.ReadFile(filepath)

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(r))
	if err != nil {
		return nil, err
	}

	return ebiten.NewImageFromImage(img), nil
}

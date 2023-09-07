package helpers

import (
	"image"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
)

var NumbersHelper Numbers

type Numbers struct {
	image *ebiten.Image
}

func (n *Numbers) ToImage(str string) *ebiten.Image {
	indexList := []int{}
	imageCount := 0

	for i := 0; i < len(str); i++ {
		if str[i] == '+' {
			indexList = append(indexList, 0)
			imageCount += 1
		} else {
			index, err := strconv.Atoi(string(str[i]))
			if err == nil {
				indexList = append(indexList, index+1)
				imageCount += 1
			}
		}
	}

	newImage := ebiten.NewImage(8*imageCount, 12)

	for i, v := range indexList {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(i*8), 0)
		newImage.DrawImage(ebiten.NewImageFromImage(n.image.SubImage(image.Rect((v*8), 0, 8+(v*8), 12))), op)
	}

	return newImage
}

func NewNumbers(sheetPath string) error {
	image, err := ImagefromFile(sheetPath)

	if err != nil {
		return err
	}

	NumbersHelper = Numbers{image: image}

	return nil
}

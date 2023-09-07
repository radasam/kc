package ui

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var UI_SPACING = 2

func DrawCursor(screen *ebiten.Image, xPos int, yPos int, w int, h int) {
	vector.StrokeLine(screen, float32(xPos-1), float32(yPos-1), float32(xPos+w+1), float32(yPos-1), 1, color.White, false)
	vector.StrokeLine(screen, float32(xPos+w+1), float32(yPos-1), float32(xPos+w+1), float32(yPos+h+1), 1, color.White, false)
	vector.StrokeLine(screen, float32(xPos+w+1), float32(yPos+h+1), float32(xPos-1), float32(yPos+h+1), 1, color.White, false)
	vector.StrokeLine(screen, float32(xPos-1), float32(yPos+h+1), float32(xPos-1), float32(yPos-2), 1, color.White, false)
}

func CreatePanel(w int, h int) *ebiten.Image {
	img := ebiten.NewImage(w, h)
	bg := ebiten.NewImage(w, h)

	fillColor := color.NRGBA{R: 255, G: 255, B: 255, A: 200}

	vector.DrawFilledRect(bg, 0, 0, float32(w), float32(h), fillColor, true)

	ops := &ebiten.DrawImageOptions{}
	ops.ColorScale.ScaleAlpha(0.5)
	img.DrawImage(bg, ops)

	return img
}

func CreateGrid(items []*ebiten.Image, cols int) *ebiten.Image {
	totalItems := len(items)
	maxWidth := 0
	maxHeight := 0

	for _, i2 := range items {
		if i2.Bounds().Size().X > maxWidth {
			maxWidth = i2.Bounds().Size().X
		}

		if i2.Bounds().Size().Y > maxHeight {
			maxHeight = i2.Bounds().Size().Y
		}
	}

	if totalItems < cols {
		cols = totalItems
	}

	totalRows := math.Ceil(float64(totalItems) / float64(cols))

	panel := CreatePanel((maxWidth+UI_SPACING*2)*cols, int(totalRows)*int(maxHeight+UI_SPACING*2))

	row := 0
	col := 0

	for _, i2 := range items {
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Translate((float64(UI_SPACING)*(float64(col)+1.0))+(float64(maxWidth)*float64(col)), (float64(UI_SPACING)*(float64(row)+1.0))+(float64(maxHeight)*float64(row)))
		panel.DrawImage(i2, ops)
		col += 1
		if col >= cols {
			col = 0
			row += 1
		}
	}

	return panel

}

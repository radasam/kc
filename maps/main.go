package maps

import (
	"bytes"
	"encoding/json"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var CurrMap *Map

type MapData struct {
	Tilesets [][][]int
}

type Map struct {
	mapData      MapData
	tileset      int
	tileSetImage *ebiten.Image
}

func (m *Map) Draw(screen *ebiten.Image) {
	for i, l := range m.GetTileset() {
		for j, t := range l {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64((j%10)*32), float64((i%10)*32))

			sx := (t % 5) * 32
			sy := (t / 5) * 32
			screen.DrawImage(m.tileSetImage.SubImage(image.Rect(sx, sy, sx+32, sy+32)).(*ebiten.Image), op)
		}
	}
}

func (m *Map) GetTileset() [][]int {
	return m.mapData.Tilesets[m.tileset]
}

func LoadMap(filepath string) error {
	f, err := os.ReadFile(filepath)

	if err != nil {
		return err
	}

	mapData := MapData{}

	err = json.Unmarshal(f, &mapData)

	if err != nil {
		return err
	}

	r, err := os.ReadFile("./tileset.png")

	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(r))
	if err != nil {
		return err
	}

	CurrMap = &Map{tileset: 0, mapData: mapData, tileSetImage: ebiten.NewImageFromImage(img)}

	return nil
}

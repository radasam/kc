package types

import (
	"bytes"
	"encoding/json"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type ItemClass string

const (
	ItemClassSeed    ItemClass = "seed"
	ItemClassHarvest ItemClass = "harvest"
)

type Item struct {
	Id        string `db:"id"`
	Weight    int
	Image     *ebiten.Image
	ImagePath string
	Class     ItemClass
	ClassData []byte
}

func (it *Item) GetImage() *ebiten.Image {
	return it.Image
}

func (it *Item) UnmarshalJSON(b []byte) error {
	var i struct {
		Id        string `db:"id"`
		Weight    int
		ImagePath string
		Class     ItemClass
		ClassData map[string]interface{}
	}

	err := json.Unmarshal(b, &i)

	if err != nil {
		return err
	}

	r, err := os.ReadFile(i.ImagePath)

	if err != nil {
		return err
	}

	img, _, err := image.Decode(bytes.NewReader(r))

	if err != nil {
		return err
	}

	byt, err := json.Marshal(&i.ClassData)

	if err != nil {
		return err
	}

	it.Id = i.Id
	it.Weight = i.Weight
	it.Image = ebiten.NewImageFromImage(img)
	it.Class = i.Class
	it.ClassData = byt

	return nil
}

type IItem interface {
	GetImage() *ebiten.Image
}

type IUseable interface {
	Use(xPos float64, yPos float64) bool
}

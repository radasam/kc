package events

import (
	"fmt"
	"image"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"kc.com/kc/db"
	"kc.com/kc/helpers"
	"kc.com/kc/types"
)

type InfoEvent interface {
	GetImage() *ebiten.Image
	GetSize() image.Point
	GetAdded() int64
}

type InventoryEvent struct {
	plus  *ebiten.Image
	item  *types.Item
	image *ebiten.Image
	added int64
}

func (ie *InventoryEvent) GetSize() image.Point {
	plusSize := ie.plus.Bounds().Size()
	itemSize := ie.item.Image.Bounds().Size()

	return image.Point{X: plusSize.X + itemSize.X, Y: int(math.Max(float64(plusSize.Y), float64(itemSize.Y)))}
}

func (ie *InventoryEvent) CreateImage() {
	size := ie.GetSize()

	img := ebiten.NewImage(size.X, size.Y)

	itemOffsetY := size.Y - ie.item.Image.Bounds().Size().Y
	itemOps := &ebiten.DrawImageOptions{}
	itemOps.GeoM.Translate(0, float64(itemOffsetY)/2.0)

	img.DrawImage(ie.item.Image, itemOps)

	plusOffsetY := size.Y - ie.plus.Bounds().Size().Y
	plusOffsetX := ie.item.Image.Bounds().Size().X
	plusOps := &ebiten.DrawImageOptions{}
	plusOps.GeoM.Translate(float64(plusOffsetX), float64(plusOffsetY)/2.0)

	img.DrawImage(ie.plus, plusOps)

	ie.image = img
}

func (ie *InventoryEvent) GetImage() *ebiten.Image {
	return ie.image
}

func (ie *InventoryEvent) GetAlpha() float64 {
	return float64(1.0 - ((ie.added - time.Now().Unix()) / 5.0))
}

func (ie *InventoryEvent) GetAdded() int64 {
	return ie.added
}

func NewInventoryEvent(id string, count int) (*InventoryEvent, error) {

	item, err := db.ItemDb.GetEntry(id)

	if err != nil {
		return nil, err
	}
	plusOne := helpers.NumbersHelper.ToImage(fmt.Sprintf("+%d", count))

	ie := &InventoryEvent{item: item, plus: plusOne, added: time.Now().UnixMilli()}

	ie.CreateImage()

	return ie, nil
}

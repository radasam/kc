package bases

import (
	"fmt"
	"math"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	items "kc.com/kc/Items"
	"kc.com/kc/helpers"
	"kc.com/kc/managers"
	"kc.com/kc/types"
	"kc.com/kc/ui"
)

type Inventory struct {
	Show         bool
	capacity     float64
	itemMap      map[string]types.InventoryItem
	items        []string
	capacityUsed float64
	image        *ebiten.Image
	imageValid   bool
	activeItem   int
	gridW        int
	gridH        int
}

func (i *Inventory) MoveCursor(x int, y int) {
	if x == 1 && i.activeItem+1 < len(i.items) {
		i.activeItem += 1
	} else if x == -1 && i.activeItem > 0 {
		i.activeItem -= 1
	} else if y == 1 && i.activeItem+3 < len(i.items) {
		i.activeItem += 3
	} else if i.activeItem-3 > 0 {
		i.activeItem -= 3
	}
	println(i.activeItem, len(i.items))
}

func (i *Inventory) Toggle() {
	if i.Show {
		i.Show = false
		managers.KeypressManager.UnBind("inventory-up")
		managers.KeypressManager.UnBind("inventory-down")
		managers.KeypressManager.UnBind("inventory-left")
		managers.KeypressManager.UnBind("inventory-right")
	} else {
		i.Show = true
		managers.KeypressManager.BindPress("inventory-up", ebiten.KeyArrowUp, func() { i.MoveCursor(0, -1) })
		managers.KeypressManager.BindPress("inventory-down", ebiten.KeyArrowDown, func() { i.MoveCursor(0, 1) })
		managers.KeypressManager.BindPress("inventory-left", ebiten.KeyArrowLeft, func() { i.MoveCursor(-1, 0) })
		managers.KeypressManager.BindPress("inventory-right", ebiten.KeyArrowRight, func() { i.MoveCursor(1, 0) })
	}
}

func (i *Inventory) AddItem(id string, count int, weight float64) bool {
	if i.capacityUsed+(weight*float64(count)) > i.capacity {
		return false
	} else {
		i.capacityUsed += (weight * float64(count))
		if _, ok := i.itemMap[id]; ok {
			i.itemMap[id] = types.InventoryItem{Count: count + i.itemMap[id].Count, Weight: weight}
		} else {
			i.itemMap[id] = types.InventoryItem{Count: count, Weight: weight}
			i.items = append(i.items, id)
		}
		i.imageValid = false
		return true
	}
}

func (i *Inventory) CheckItem(id string) int {
	if e, ok := i.itemMap[id]; ok {
		return e.Count
	} else {
		return 0
	}
}

func (i *Inventory) RemoveItem(id string, count int) bool {
	has := i.CheckItem(id)

	if has < count {
		return false
	} else {
		entry := i.itemMap[id]
		i.capacityUsed -= entry.Weight * float64(count)
		if has == count {
			delete(i.itemMap, id)
			i.items = helpers.FilterSlice(i.items, id)
		} else {
			i.itemMap[id] = types.InventoryItem{Count: i.itemMap[id].Count - count, Weight: entry.Weight}
		}
		i.imageValid = false
		return true
	}
}

func (i *Inventory) createImageForEntry(itemId string, count int) (*ebiten.Image, error) {
	success, item := items.Create(itemId)

	if !success {
		return nil, fmt.Errorf("item doesnt exist")
	}

	countImg := helpers.NumbersHelper.ToImage(strconv.Itoa(count))

	h := math.Max(float64(item.GetImage().Bounds().Size().Y), float64(countImg.Bounds().Size().Y))
	w := item.GetImage().Bounds().Size().X + countImg.Bounds().Size().X

	newImage := ebiten.NewImage(w, int(h))

	itemOp := &ebiten.DrawImageOptions{}
	itemOp.GeoM.Translate(0, (h-float64(item.GetImage().Bounds().Size().Y))/2.0)
	newImage.DrawImage(item.GetImage(), itemOp)

	countOp := &ebiten.DrawImageOptions{}
	countOp.GeoM.Translate(float64(item.GetImage().Bounds().Size().X), (h-float64(countImg.Bounds().Size().Y))/2.0)
	newImage.DrawImage(countImg, countOp)

	return newImage, nil
}

func (i *Inventory) Draw(screen *ebiten.Image, baseX float64, baseY float64) {
	if len(i.itemMap) > 0 {
		if !i.imageValid {
			images := make([]*ebiten.Image, len(i.itemMap))
			for ind, id := range i.items {
				item := i.itemMap[id]

				img, err := i.createImageForEntry(id, item.Count)
				if err != nil {
					println(err.Error())
				}
				images[ind] = img
				if img.Bounds().Size().X > i.gridW {
					i.gridW = img.Bounds().Size().X
				}
				if img.Bounds().Size().Y > i.gridH {
					i.gridH = img.Bounds().Size().Y
				}
			}

			panel := ui.CreateGrid(images, 3)

			i.image = panel
			i.imageValid = true
		}

		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Translate(baseX, baseY-float64(i.image.Bounds().Size().Y))
		screen.DrawImage(i.image, ops)
		xGrid := i.activeItem % 3
		yGrid := math.Floor(float64(i.activeItem) / 3)
		ui.DrawCursor(screen, int(baseX)+ui.UI_SPACING+((i.gridW+ui.UI_SPACING)*xGrid), int(baseY-float64(i.image.Bounds().Size().Y))+ui.UI_SPACING+((i.gridH+ui.UI_SPACING)*int(yGrid)), i.gridW+ui.UI_SPACING, i.gridH+ui.UI_SPACING)
	}

}

func (i *Inventory) UseItem(xPos float64, yPos float64) bool {
	if len(i.items) > 0 {
		success, item := items.Create(i.items[i.activeItem])

		if !success {
			return false
		}

		if useable, ok := item.(types.IUseable); ok {
			consumed := useable.Use(xPos, yPos)

			if consumed {
				i.RemoveItem(i.items[i.activeItem], 1)
			}
			return true
		}
	}

	return false
}

func NewInventory(capacity float64) *Inventory {
	return &Inventory{
		capacity:     capacity,
		capacityUsed: 0,
		itemMap:      map[string]types.InventoryItem{},
		imageValid:   false,
		activeItem:   0,
	}
}

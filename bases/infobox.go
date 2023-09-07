package bases

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"kc.com/kc/events"
)

type Infobox struct {
	events map[int]events.InfoEvent
}

func (ib *Infobox) ShiftEvents() {
	for i := 4; i >= 0; i-- {
		if i == 0 {
			delete(ib.events, i)
		}
		ib.events[i] = ib.events[i-1]
	}
}

func (ib *Infobox) AddEvent(e events.InfoEvent) {
	ib.ShiftEvents()
	ib.events[0] = e
}

func (ib *Infobox) Draw(screen *ebiten.Image, baseX float64, baseY float64) {
	yOffset := 2.0
	now := time.Now().UnixMilli()

	for i := 0; i < 5; i++ {
		if v, ok := ib.events[i]; ok && v != nil {
			if now-v.GetAdded() > 1000 {
				delete(ib.events, i)
			} else {
				xPos := baseX - float64(v.GetSize().X)/2.0
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(xPos, baseY-yOffset-32)
				op.ColorScale.ScaleAlpha(float32(1) - (float32((now - v.GetAdded())) / float32(1000)))
				screen.DrawImage(v.GetImage(), op)
			}
		}
		yOffset += 18
	}
}

func NewInfoBox() *Infobox {
	return &Infobox{events: map[int]events.InfoEvent{}}
}

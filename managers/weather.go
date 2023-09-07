package managers

import (
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"kc.com/kc/images"
)

var WeatherManager *Weather

type Weather struct {
	sunSheet          *images.SpriteSheet
	lightTimeSeries   []int
	lightLevels       []int
	currentLightLevel int
	startTime         int64
	currentTime       int64
	overlayedImage    *ebiten.Image
}

func (w *Weather) DrawLighting(screen *ebiten.Image) {
	if w.overlayedImage == nil || w.currentLightLevel != w.lightLevels[w.lightTimeSeries[w.currentTime]] {
		if w.lightLevels[w.lightTimeSeries[w.currentTime]] < w.currentLightLevel {
			w.currentLightLevel -= 1
		} else {
			w.currentLightLevel += 1
		}

		bg := ebiten.NewImage(screen.Bounds().Dx(), screen.Bounds().Dy())

		fillColor := color.NRGBA{R: 0, G: 0, B: 0, A: uint8(w.currentLightLevel)}

		vector.DrawFilledRect(bg, 0, 0, float32(screen.Bounds().Dx()), float32(screen.Bounds().Dy()), fillColor, true)

		w.overlayedImage = bg
	}

	op := &ebiten.DrawImageOptions{}

	screen.DrawImage(w.overlayedImage, op)
}

func (w *Weather) GetWeather() int {
	return w.lightTimeSeries[w.currentTime]
}

func (w *Weather) OnUpdate() {
	now := time.Now().Unix()

	if now-w.startTime > (w.currentTime+1)*10 && w.currentTime < 11 {
		w.currentTime += 1
		w.sunSheet.SetRow(w.lightTimeSeries[w.currentTime])
		w.overlayedImage = nil
	}
}

func (w *Weather) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screen.Bounds().Dx())-30, 6)
	screen.DrawImage(w.sunSheet.GetImage(), op)
}

func NewWeatherManager() error {
	image, err := ImageManager.GetImage("./sun-sheet.png", "weather-sun")

	if err != nil {
		return err
	}

	sunSheet := images.NewSpriteSheet(image, 24, 24, 0)

	WeatherManager = &Weather{
		sunSheet:          sunSheet,
		lightTimeSeries:   []int{0, 1, 2, 3, 4, 5, 5, 4, 3, 2, 1, 0},
		lightLevels:       []int{200, 175, 150, 100, 50, 0},
		startTime:         time.Now().Unix(),
		currentTime:       0,
		currentLightLevel: 200,
	}

	return nil
}

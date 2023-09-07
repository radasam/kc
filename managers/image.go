package managers

import (
	"bytes"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var ImageManager *Image

type Image struct {
	imageMap           map[string]*ebiten.Image
	indentifierMap     map[string]int //image id > [] indentifier id
	indentifierToImage map[string]string
}

func (i *Image) addToObjectMap(filepath string, indentifier string) {
	i.indentifierToImage[indentifier] = filepath
	if _, ok := i.indentifierMap[filepath]; ok {
		i.indentifierMap[filepath] += 1
	} else {
		i.indentifierMap[filepath] = 1
	}
}

func (i *Image) RemoveImage(indentifier string) {
	if filepath, ok := i.indentifierToImage[indentifier]; ok {
		if i.indentifierMap[filepath] == 1 {
			delete(i.indentifierMap, filepath)
			delete(i.imageMap, filepath)
		} else {
			i.indentifierMap[filepath] -= 1
		}

		delete(i.indentifierToImage, indentifier)
	}
}

func (i *Image) GetImage(filepath string, indentifier string) (*ebiten.Image, error) {
	if img, ok := i.imageMap[filepath]; ok {
		i.addToObjectMap(filepath, indentifier)
		return img, nil
	} else {
		r, err := os.ReadFile(filepath)

		if err != nil {
			return nil, err
		}

		goimg, _, err := image.Decode(bytes.NewReader(r))
		if err != nil {
			return nil, err
		}

		img := ebiten.NewImageFromImage(goimg)

		i.imageMap[filepath] = img
		i.addToObjectMap(filepath, indentifier)
		return img, nil
	}
}

func NewImageManager() {
	ImageManager = &Image{imageMap: map[string]*ebiten.Image{}, indentifierMap: map[string]int{}, indentifierToImage: map[string]string{}}
}

package characters

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type DirectionalSprite struct {
	sprite  *Sprite
	lastX   float64
	lastY   float64
	currDir int
	flipped bool
	state   int
	dist    float64
}

func (s *DirectionalSprite) stateChange(x float64, y float64) {
	delta := math.Abs(x-s.lastX) + math.Abs(y-s.lastY)

	if delta == 0 {
		s.state = 0
	} else {
		s.dist += delta
		if s.dist > 10 {
			s.state += 1
			if s.state > 3 {
				s.state = 0
			}
			s.dist -= 10
		}
	}

}

func (s *DirectionalSprite) Draw(screen *ebiten.Image, x float64, y float64) {

	s.stateChange(x, y)
	if y > s.lastY {
		s.sprite.Draw(screen, x, y, 0, s.state, false)
		s.currDir = 0
		s.flipped = false
	} else if x > s.lastX {
		s.sprite.Draw(screen, x, y, 1, s.state, false)
		s.currDir = 1
		s.flipped = false
	} else if y < s.lastY {
		s.sprite.Draw(screen, x, y, 2, s.state, false)
		s.currDir = 2
		s.flipped = false
	} else if x < s.lastX {
		s.sprite.Draw(screen, x, y, 1, s.state, true)
		s.currDir = 1
		s.flipped = true
	} else {
		s.sprite.Draw(screen, x, y, s.currDir, s.state, s.flipped)
	}
	s.lastX = x
	s.lastY = y
}

func NewDirectionalSprite(filepath string, x float64, y float64, width float64, height float64) (*DirectionalSprite, error) {

	sprite, err := NewSprite(filepath, width, height)

	if err != nil {
		return nil, err
	}

	return &DirectionalSprite{sprite: sprite, lastX: x, lastY: y, currDir: 2, flipped: false, state: 0, dist: 0}, nil
}

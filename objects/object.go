package objects

import (
	"fmt"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
	"kc.com/kc/managers"
	"kc.com/kc/types"
)

type Object struct {
	id    string
	xPos  float64
	yPos  float64
	image *ebiten.Image
	w     float64
	h     float64
	state int
}

func (o *Object) GetID() string {
	return o.id
}

func (o *Object) GetPos() []float64 {
	return []float64{o.xPos, o.yPos}
}

func (o *Object) GetDim() []float64 {
	return []float64{o.w, o.h}
}

func (o *Object) GetVerts() []types.Vert {
	verts := []types.Vert{}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			verts = append(verts, types.Vert{X: o.xPos + (o.w * float64(i)) - o.w/2, Y: o.yPos + (o.h * float64(j)) - o.h/2, Id: fmt.Sprintf("%s-%d", o.id, (i*2)+(j))})
		}
	}

	return verts
}

func (o *Object) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.xPos-o.w/2, o.yPos-o.h/2)
	screen.DrawImage(ebiten.NewImageFromImage(o.image.SubImage(image.Rect(0, 0+(o.state*int(o.h)), int(o.w), int(o.h)+(o.state*int(o.h))))), op)
}

func (o *Object) Init() {

}

func (o *Object) IncrementState() {
	o.state += 1
}

func (o *Object) SetState(state int) {
	o.state = state
}

func NewObject(id string, filepath string, xPos float64, yPos float64, w float64, h float64) (*Object, error) {
	img, err := managers.ImageManager.GetImage(filepath, id)

	if err != nil {
		return nil, err
	}

	return &Object{id: id, image: img, xPos: xPos, yPos: yPos, w: w, h: h, state: 0}, nil
}

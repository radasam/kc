package managers

import (
	"image/color"
	"strconv"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"kc.com/kc/helpers"
	"kc.com/kc/types"
)

const (
	MAX_DOT_COUNT = 4
	MIN_QUAD_SIZE = 64
)

type Quad struct {
	totalDots int
	id        string
	Children  []string
	Dots      map[string]types.Vert
	xPos      float64
	yPos      float64
	Dim       float64
	Parent    string
}

func (q *Quad) Merge() {
	if q.totalDots <= MAX_DOT_COUNT {
		for _, c := range q.Children {
			cq := CollisionManager.QuadMap[c]
			for _, v := range cq.Dots {
				q.Dots[v.Id] = v
				CollisionManager.RemoveVertToQuad(v.Id)
				CollisionManager.AddVertToQuad(v.Id, q.id)
			}
			CollisionManager.RemoveFromQuadMap(c)
		}

		q.Children = []string{}
		if q.Parent != "" {
			CollisionManager.QuadMap[q.Parent].Merge()
		}
	}

}

func (q *Quad) ReduceDots() {
	q.totalDots -= 1
	if q.Parent != "" {
		CollisionManager.QuadMap[q.Parent].ReduceDots()
	}
}

func (q *Quad) AddDot() {
	q.totalDots += 1
	if q.Parent != "" {
		CollisionManager.QuadMap[q.Parent].AddDot()
	}
}

func (q *Quad) RemoveDot(v types.Vert) {
	q.ReduceDots()
	delete(q.Dots, v.Id)
	if q.Parent != "" {
		CollisionManager.QuadMap[q.Parent].Merge()
	}
}

func (q *Quad) Draw(screen *ebiten.Image) {

	if len(q.Children) == 0 {
		vector.StrokeLine(screen, float32(q.xPos+q.Dim), float32(q.yPos), float32(q.xPos+q.Dim), float32(q.yPos+q.Dim), 1, color.White, false)
		vector.StrokeLine(screen, float32(q.xPos+q.Dim), float32(q.yPos+q.Dim), float32(q.xPos), float32(q.yPos+q.Dim), 1, color.White, false)
	}
}

func (q *Quad) Includes(v types.Vert) (bool, *Quad) {
	if len(q.Children) == 0 {

		if v.X > q.xPos && v.X <= q.xPos+q.Dim {
			if v.Y > q.yPos && v.Y <= q.yPos+q.Dim {
				return true, q
			}
		}
	} else {
		for _, iq := range q.Children {
			inc, q := CollisionManager.QuadMap[iq].Includes(v)

			if inc {
				return inc, q
			}
		}
	}

	return false, nil
}

func (q *Quad) Append(v types.Vert) {
	if len(q.Children) == 0 {
		q.AddDot()
		if len(q.Dots) == MAX_DOT_COUNT && q.Dim > MIN_QUAD_SIZE {
			q.Dots[v.Id] = v
			children := SplitQuad(q.id, q.xPos, q.yPos, q.Dim, q.Dots)

			q.Dots = map[string]types.Vert{}
			q.Children = children
		} else {
			CollisionManager.AddVertToQuad(v.Id, q.id)
			q.Dots[v.Id] = v
		}
	} else {
		for _, iq := range q.Children {
			CollisionManager.RemoveVertToQuad(v.Id)
			inc, cq := CollisionManager.QuadMap[iq].Includes(v)

			if inc {
				q.totalDots += 1
				cq.Append(v)
				break
			}
		}
	}

}

func SplitQuad(id string, xPos float64, yPos float64, dim float64, dots map[string]types.Vert) []string {

	quads := []string{}

	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			newId := id + strconv.Itoa((i*2)+(j))
			newDim := dim / 2
			newXPos := xPos + (float64(i) * dim / 2)
			newYPos := yPos + (float64(j) * dim / 2)
			newDots := map[string]types.Vert{}

			for _, v := range dots {
				if v.X > newXPos && v.X <= newXPos+newDim {
					if v.Y > newYPos && v.Y <= newYPos+newDim {
						CollisionManager.RemoveVertToQuad(v.Id)
						newDots[v.Id] = v
						CollisionManager.AddVertToQuad(v.Id, newId)
					}
				}
			}

			newQuad := &Quad{
				totalDots: len(newDots),
				id:        newId,
				xPos:      newXPos,
				yPos:      newYPos,
				Dim:       newDim,
				Dots:      newDots,
				Parent:    id,
				Children:  []string{},
			}

			if len(newDots) > MAX_DOT_COUNT && newDim/2 > MIN_QUAD_SIZE {
				children := SplitQuad(newId, newXPos, newYPos, newDim, newDots)
				newQuad.Dots = map[string]types.Vert{}
				newQuad.Children = children

			}

			CollisionManager.QuadMap[newId] = newQuad
			quads = append(quads, newId)
		}
	}

	return quads
}

var CollisionManager *Collision

type Collision struct {
	QuadMap    map[string]*Quad // quadId to quad
	ObjMap     map[string]types.Collideable
	vertToQuad map[string]string
}

func (c *Collision) GetQuads(dot types.Collideable) []*Quad {
	quads := []*Quad{}

	for _, v := range dot.GetVerts() {
		quads = append(quads, c.QuadMap[c.vertToQuad[v.Id]])
	}

	return quads
}

func (c *Collision) CheckCollisions(collideable types.Collideable) (types.CollisionType, types.Collideable) {
	rad := collideable.GetRad()
	q := c.GetQuads(collideable)
	dots := map[string]string{}

	for _, quad := range q {
		qdots := quad.Dots
		if qdots != nil {
			for _, v := range qdots {
				dotId := strings.Split(v.Id, "-")
				if _, ok := dots[dotId[0]]; !ok {
					dots[dotId[0]] = ""
				}
			}
		}

	}

	var collidesWith types.Collideable

	hasSoftCollision := false
	hasHardCollision := false

	for dotId, _ := range dots {
		dot := c.ObjMap[dotId]

		if dot.GetID() != collideable.GetID() {
			if helpers.GetDist(collideable, dot) <= rad {

				if helpers.CheckIntersection(collideable, dot) {
					collideable.OnCollision(types.HardCollision, dot)
					dot.OnCollision(types.HardCollision, collideable)
					if !hasHardCollision {
						hasHardCollision = true
						collidesWith = dot
					}
				}

				collideable.OnCollision(types.SoftCollision, dot)
				dot.OnCollision(types.SoftCollision, collideable)
				if !hasHardCollision {
					hasSoftCollision = true
					collidesWith = dot
				}
			}
		}
	}

	if hasHardCollision {
		return types.HardCollision, collidesWith
	} else if hasSoftCollision {
		return types.SoftCollision, collidesWith
	} else {
		return types.NoCollision, nil
	}
}

func (c *Collision) Register(q *Quad) {
	c.QuadMap[q.id] = q
}

func (c *Collision) AppendCollidable(col types.Collideable) {
	c.ObjMap[col.GetID()] = col
	for _, v := range col.GetVerts() {
		c.AppendVert(v)
	}

}

func (c *Collision) AppendVert(v types.Vert) {
	for i := 0; i < 4; i++ {
		inc, q := c.QuadMap[strconv.Itoa(i)].Includes(v)

		if inc {
			q.Append(v)
			break
		}
	}
}

func (c *Collision) RemoveCollideable(col types.Collideable) {
	delete(c.QuadMap, col.GetID())
	for _, v := range col.GetVerts() {
		c.RemoveVert(v)
	}
}

func (c *Collision) RemoveVert(v types.Vert) {
	if qid, ok := c.vertToQuad[v.Id]; ok {
		if q, ok := c.QuadMap[qid]; ok {
			q.RemoveDot(v)
		}
	}
}

func (c *Collision) UpdateVert(v types.Vert) {
	c.RemoveVert(v)
	c.AppendVert(v)
}

func (c *Collision) UpdateCollidable(col types.Collideable) {
	for _, v := range col.GetVerts() {
		c.UpdateVert(v)
	}
}

func (c *Collision) Draw(screen *ebiten.Image) {
	for _, q := range c.QuadMap {
		q.Draw(screen)
	}
}

func (c *Collision) RemoveFromQuadMap(qid string) {
	delete(c.QuadMap, qid)
}

func (c *Collision) AddVertToQuad(vid string, qid string) {
	c.vertToQuad[vid] = qid
}

func (c *Collision) RemoveVertToQuad(vid string) {
	delete(c.vertToQuad, vid)
}

func (c *Collision) VertToQuad(vid string) string {
	return c.vertToQuad[vid]
}

func NewCollisionManager(dim float64) {

	CollisionManager = &Collision{
		QuadMap:    map[string]*Quad{},
		ObjMap:     map[string]types.Collideable{},
		vertToQuad: map[string]string{},
	}

	SplitQuad("", 0, 0, dim, map[string]types.Vert{})

	DebugManager.AddDraw("quadmap", CollisionManager)
}

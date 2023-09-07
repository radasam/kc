package helpers

import (
	"math"

	"kc.com/kc/types"
)

func GetDist(c1 types.Collideable, c2 types.Collideable) float64 {
	p1 := c1.GetPos()
	p2 := c2.GetPos()

	return math.Pow(math.Pow(p1[0]-p2[0], 2)+math.Pow(p1[1]-p2[1], 2), 0.5)
}

func GetVertices(c types.Collideable) []types.Point {
	pos := c.GetPos()
	dim := c.GetDim()

	return []types.Point{{X: pos[0] - dim[0]/2, Y: pos[1] - dim[1]/2}, {X: pos[0] + dim[0]/2, Y: pos[1] + dim[1]/2}}

}

func CheckIntersection(c1 types.Collideable, c2 types.Collideable) bool {
	v1 := GetVertices(c1)
	v2 := GetVertices(c2)

	if v1[0].X < v2[1].X && v1[1].X > v2[0].X && v1[0].Y < v2[1].Y && v1[1].Y > v2[0].Y {
		return true
	}

	return false
}

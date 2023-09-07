package types

type CollisionType int

const (
	NoCollision   CollisionType = 0
	SoftCollision CollisionType = 1
	HardCollision CollisionType = 2
)

type Collideable interface {
	GetID() string
	GetPos() []float64
	GetRad() float64
	GetDim() []float64
	IsSolid() bool
	OnCollision(CollisionType, Collideable)
	GetVerts() []Vert
}

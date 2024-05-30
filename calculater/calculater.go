package calculater

import (
	"math"
)

type Square struct {
	X1, Y1, X2, Y2 float64
}

func NewGeo(x1, y1, x2, y2 float64) *Square {
	return &Square{
		X1: x1,
		Y1: y1,
		X2: x2,
		Y2: y2,
	}
}

func (g *Square) CalculateDistance() float64 {

	return math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
}

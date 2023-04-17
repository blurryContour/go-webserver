package pkg

import "math"

type Vector2 struct {
	X, Y float64
}

func (v *Vector2) Length() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

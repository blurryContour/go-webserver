package main

import (
	"fmt"
	"math"

	"github.com/blurryContour/go-webserver/pkg"
)

type Vertex struct {
	X, Y float64
	L    float64
}

// Receivers used to add method to struct
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Length() {
	v.L = math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4, 0}

	fmt.Println(v.Abs())
	fmt.Println(v.L)

	vec := pkg.Vector2{X: 3, Y: 4}
	fmt.Println(vec.Length())
}

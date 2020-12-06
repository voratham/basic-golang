package shape

import "math"

// Circle is struct of shape
type Circle struct {
	Radius float64
}

// Area is function
func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

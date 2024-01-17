package compositions

import "math"

type Shape interface {
	Area() float64
}

type Rectange struct {
	width, height float64
}

func (r Rectange) Area() float64 {
	return r.width * r.height
}

func (r Rectange) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

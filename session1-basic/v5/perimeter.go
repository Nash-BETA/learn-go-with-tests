package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Redius float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width + r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Redius * c.Redius
}

type Shape interface {
	Area() float64
}

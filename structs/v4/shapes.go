package main

import "math"

// Rectangle has the dimensions of a rectangle
type Rectangle struct {
	Width  float64
	Height float64
}

// Area returns the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter returns the perimeter of a rectangle
func (rectangle Rectangle) Perimeter() float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Circle represents a circle...
type Circle struct {
	Radius float64
}

// Area returns the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter returns the perimeter of a circle
func (c Circle) Perimeter() float64 {
	return math.Pi * 2 * c.Radius
}

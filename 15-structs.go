package main

import (
	"fmt"
	"math"
)

func main() {
	var c1 Circle
	c1.x = 0
	c1.y = 0
	c1.r = 10

	c2 := Circle{x: 0, y: 0, r: 10}

	c3 := Circle{0, 0, 10}

	cPtr1 := &Circle{0, 0, 10}

	cPtr2 := new(Circle)
	(*cPtr2).x = 0
	(*cPtr2).y = 0
	(*cPtr2).r = 10

	fmt.Println(c1.r, c2.r, c3.r, (*cPtr1).r, (*cPtr2).r)

	fmt.Println("The area of the circle is:", c1.area())

	aCircle := Circle{r: 10, x: 0, y: 0}
	aSquare := Square{l: 10}
	aRectangle := Rectangle{w: 10, h: 20}
	shapes := []Shape{aCircle, aSquare, aRectangle}
	fmt.Println("Total area is:", totalArea(shapes))

}

// Circle : abstraction of a circle
// with the coordinates of its center
// and radius
type Circle struct {
	x float64
	y float64
	r float64
}

// Square : a square struct
type Square struct {
	l float64
}

// Rectangle : a rectangle struct
type Rectangle struct {
	h float64
	w float64
}

// Shape : a shape interface
type Shape interface {
	area() float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (s Square) area() float64 {
	return s.l * s.l
}

func (r Rectangle) area() float64 {
	return r.h * r.w
}

func totalArea(shapes []Shape) float64 {
	total := 0.0
	for _, shape := range shapes {
		total += shape.area()
	}
	return total
}

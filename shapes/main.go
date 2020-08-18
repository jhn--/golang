package main

import (
	"fmt"
	"math"
)

/*
Structs and Interfaces
*/

// Circle is a circle.
type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	// this function is a 'method'.
	// (c *Circle) is a 'receiver', which is like a parameter (it has a name and a type).
	// creating the creating the function this way allows us to call the function using the `.` operator.
	// when calling c.area(), Go naturally knows to pass a pointer to Circle object for 'methods'.
	return math.Pi * c.r * c.r
}

// Rectangle is a rectangle.
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	// same type (it is a 'method') as c.area().
	l := math.Abs(r.x1 - r.x2)
	w := math.Abs(r.y1 - r.y2)
	return l * w
}

func main() {
	// initialization of Circle object
	c := Circle{x: 10, y: 10, r: 5}
	// accessing fields of Circle object via the `.` operator.
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(c.area())

	// initialize Rectangle object
	// we can omit the field names if we know the order they are defined.
	r := Rectangle{3, 10, 7, 22} // x1, y1, x2, y2 => l = 4, w = 12
	fmt.Println(r.x1, r.y1, r.x2, r.y2)
	fmt.Println(r.area())

	r.x1 = 4
	fmt.Println(r.x1, r.y1, r.x2, r.y2)
	fmt.Println(r.area())
}

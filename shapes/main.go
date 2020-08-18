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

// Interfaces are created using the `type` keyword.
// https://golang.org/ref/spec#Interface_types
// We are able to name the Rectangle's area method the same thing as Circle's area
// This ought to be achiveable as in reali life and in programming, relationships like these are common.
// Go has a way of making these accidental similarities explicit through a type known as an `Interface`.

// Shape is an interface type.
type Shape interface {
	area() float64
}

// this is a method set.
// https://golang.org/ref/spec#Method_sets.
// the method set of any interface type is called `interface`.
func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

/*
Embedded Types
*/

// Embedded types represent is-a relationships.
// ie. an Android IS A Person.
// Embedded types are also known as anonymous field.

// Person is a person.
type Person struct {
	Name string
}

func (p *Person) talk() {
	fmt.Println("Hello, I am", p.Name)
}

// now we use embedded types

// Android is a Person.
type Android struct {
	// Person Person // we don't need this syntax.
	// we use the type (Person) and don't give it a name.
	Person
	Model string
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

	fmt.Println("Total area for shapes, first Circle, then Rectangle - ")
	fmt.Println(totalArea(&c, &r))

	// new Person call Pete
	Pete := Person{"Pete"}
	Pete.talk()

	// https://golang.org/pkg/builtin/#new
	// declare a new Android
	Andy := new(Android)
	// Assign it a Name
	Andy.Name = "Andy"
	// both lines below work, the 2nd line is better as we can obviously access the `Person` method directly.
	Andy.Person.talk()
	Andy.talk()
}

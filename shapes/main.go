package main

import (
    "fmt";
    "math"
)

/*
Structs and Interfaces
*/

type Circle struct {
    x, y, r float64
}

func (c *Circle) area() float64 {
    return math.Pi * c.r * c.r
}

func main() {
    c := Circle{x: 10, y: 10, r: 5}

    fmt.Println(c.x, c.y, c.r)
    fmt.Println(c.area())
}

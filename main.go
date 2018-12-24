package main

import (
    "fmt";
)

/*
Structs and Interfaces
*/

func main() {
    type Circle struct {
        x, y, r float64
    }

    c := Circle{x: 10, y: 10, r: 5}

    fmt.Println(c.x, c.y, c.r)
}

package main

import "fmt"

/*
functions

Closure
*/

func main() {
    add := func(x, y int) int {
        return x + y
    }

    fmt.Println("Closure")
    fmt.Println(add(1,2))
}

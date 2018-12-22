package main

import "fmt"

/*
functions

recursion
*/

func main() {
    y := factorial(7)
    fmt.Println(y)
}

func factorial(x uint) uint {
    // fmt.Println("x =", x)
    if x == 1 {
        return 1
    } else {
        return x * factorial(x - 1)
    }
}
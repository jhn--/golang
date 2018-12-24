package main

import "fmt"

/*
pointers
*/

func square(x *float64) {
    fmt.Println(*x)
    *x = *x * *x
}

func main() {
    x := 1.5
    square(&x)
    fmt.Println(x)
}

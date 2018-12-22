package main

import "fmt"

/*
functions

Returning Multiple Values
*/

func main() {
    a, b, c, d := f()
    fmt.Println(a, b, c, d)
}

func f() (a int, b int, c int, d int) {
    x := 1
    y := 1
    a = x + y
    b = x - y
    c = x * y
    d = x / y
    return a, b, c, d
}
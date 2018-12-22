package main

import "fmt"

/*
functions

Write a function which takes an integer and halves it and returns true if it was even or false if it was odd. For example half(1) should return (0, false) and half(2) should return (1, true).
*/

func gethalf(arg int) (int, bool) {
    fmt.Println(arg)
    div := arg / 2
    if arg % 2 == 0 {
        return div, true
    } else {
        return div, false
    }
}


func main() {
    a, b := gethalf(1)
    fmt.Println(a, b)
    c, d := gethalf(2)
    fmt.Println(c, d)
    e, f := gethalf(3)
    fmt.Println(e, f)
    g, h := gethalf(4)
    fmt.Println(g, h)
    i, j := gethalf(5)
    fmt.Println(i, j)
    k, l := gethalf(6)
    fmt.Println(k, l)
    m, n := gethalf(7)
    fmt.Println(m, n)
}
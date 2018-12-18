package main

import "fmt"

/*
slices
*/

func main() {
    // x := [5]float64 {98, 93, 77, 82, 83}

    /*
    x := [5]float64 {
        98,
        93,
        77,
        82,
        83,
    }
    */

    var x [10]int
    fmt.Println(x)

    y := [10]int {1,2,3,4,5,6,7,8,9}
    fmt.Println(y)

    z := make([]float64, 5, 10)
    fmt.Println(z)

    a := y[:5]
    fmt.Println(a)

    b := y[3:8]
    fmt.Println(b)

    c := y[:3]
    fmt.Println(c)

    d := y[:]
    fmt.Println(d)
}

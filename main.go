package main

import "fmt"

/*
Write another program that converts from feet into meters. (1 ft = 0.3048 m)
*/

func main() {
    fmt.Print("Enter a temperature in feet: ")
    var feet float64
    fmt.Scanf("%f", &feet)

    meter := feet * 0.3048
    fmt.Println(meter)
}

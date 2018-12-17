package main

import "fmt"

/*
Using the example program as a starting point, write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
*/

func main() {
    fmt.Print("Enter a temperature in Fahrenheit: ")
    var fahrenheit float64
    fmt.Scanf("%f", &fahrenheit)

    celsius := ((fahrenheit - 32.0) * (5.0/9.0))
    fmt.Println(celsius)
}

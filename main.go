package main

import "fmt"

// this is a comment 

func main() {
    fmt.Print("Enter a number: ")
    var input float64
    fmt.Scanf("%f", &input)

    output := input * 2
    fmt.Println(output)
}

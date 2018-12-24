package main

import "fmt"

/*
pointers
*/

func zero(x int) int {
    x = 0
    return x
}

func main() {
    x := 5
    fmt.Println("x is ",x ,"with address of", &x)
    x = zero(x) // why &x ?
    fmt.Println("x is now ",x ,"with address of ", &x)
}
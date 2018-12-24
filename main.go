package main

import "fmt"

/*
pointers
*/

func zero(xPtr *int) { // whats * for?
    *xPtr = 0
}

func main() {
    x := 5
    fmt.Println("x is ",x ,"with address of", &x)
    zero(&x) // why &x ?
    fmt.Println("x is now ",x ,"with address of ", &x)
}

// * and & seems to be syntax for pointers etc, find out more.
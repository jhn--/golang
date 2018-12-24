package main

import "fmt"

/*
pointers
*/

func zero(x int) {
    x = 0
}

func main() {
    // zero function will not modify the original x variable in the main function
    x := 5
    fmt.Println("x is ", x)
    zero(x)
    fmt.Println("after zero(), x is now", x)
}
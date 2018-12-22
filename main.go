package main

import "fmt"

/*
functions

defer
*/

func first() {
    fmt.Println("first")
}

func second() {
    fmt.Println("second")
}

func main() {
    defer first()
    second()
}
package main

import "fmt"

/*
pointers

Write a program that can swap two integers (x := 1; y := 2; swap(&x, &y) should give you x=2 and y=1).
*/

func swap(xPtr *int, yPtr *int) {
    zPtr := new(int)
    *zPtr = *xPtr

    *xPtr = *yPtr
    *yPtr = *zPtr
}

func main() {
    x := 1
    y := 2
    fmt.Println("original x = ", x)
    fmt.Println("original y = ", y)
    swap(&x, &y)
    fmt.Println("new x = ", x)
    fmt.Println("new y = ", y)
}

// create a buffer pointer(zPtr)
// assign value of xPtr to zPtr
// assign value of yPtr to xPtr
// assign value of zPtr to yPtr
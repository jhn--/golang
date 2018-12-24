package main

import "fmt"

/*
pointers

new
new takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
*/

func one(xPtr *int) {
    // assigns the integer value of 1 to pointer.
    *xPtr = 1
}

func main() {
    // creates a new pointer using new() with type int and assign it to variable xPtr
    xPtr := new(int)
    // passes the pointer to one()
    one(xPtr)
    //* deferences the pointer and gives us access to the value
    fmt.Println(*xPtr)
}

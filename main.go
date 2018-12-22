package main

import "fmt"

/*
functions

Closure
increment() adds 1 to the variable x which is defined in the main function's scope. This x variable _can_ be accessed and modified by the increment function. This is why the first time we call increment we see 1 displayed, but the second time we call it we see 2 displayed.

A function like this together with the non-local variables it references is known as a closure. In this case increment and the variable x form the closure.
*/

func main() {
    x := 0
    increment := func() int {
        x++
        return x
    }

    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment())
    fmt.Println(increment())
}

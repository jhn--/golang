package main

import "fmt"

/*
functions

Closure
One way to use closure is by writing a function which returns another function which – when called – can generate a sequence of numbers. For example here's how we might generate all the even numbers:
*/

func makeEvenGenerator() func() uint {
    // returns a function and a uint...?
    i := uint(0) // initialize new variable 
    return func() (ret uint) { 
    // returns a function, ... which returns am uint
        ret = i
        i += 2
        return
    }
}

func main() {
    nextEven := makeEvenGenerator()
    // nextEven is a name for a function now, so nextEven is now nextEven()
    // if you print nextEven, you'll get its memory address ie. 0x10910b0
    fmt.Println(nextEven()) // 0
    fmt.Println(nextEven()) // 2
    fmt.Println(nextEven()) // 4
    fmt.Println(nextEven()) // 6
}

/*
makeEvenGenerator returns a function which generates even numbers. Each time it's called it adds 2 to the local i variable which – unlike normal local variables – persists between calls.

feeels like some kind of python generator
*/
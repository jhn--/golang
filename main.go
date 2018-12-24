package main

import "fmt"

/*
functions

panic & recover
*/

func main() {
    // recover() happens due to defer()

    defer func() {
        str := recover()
        fmt.Println(str)
        fmt.Println("Recovered")
    }() // what is this trailing () parenthesis for?

    panic("PANIC")
}
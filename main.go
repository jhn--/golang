package main

import "fmt"

/*
functions

panic & recover
*/

func main() {
    // recover() will never happen
    panic("PANIC")
    str := recover()
    fmt.Println(str)
}
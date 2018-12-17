package main

import "fmt"

// this is a comment 
var x string = "Hello, World"

func main() {
    fmt.Println(x)
    f()
}

func f() {
    fmt.Println(x)
}

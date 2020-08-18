package main

import "fmt"

func main() {
    x := [6]string{"a","b","c","d","e","f"}
    fmt.Println(len(x))

    if x[2:5] == ["c","d","e"]{
        fmt.Println(true)
    }
}
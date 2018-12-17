package main

import "fmt"

// this is a comment 

func main() {
    answer := (true && false) || (false && true) || !(false && false)
    fmt.Println(answer)
}
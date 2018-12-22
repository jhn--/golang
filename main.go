package main

import "fmt"

/*
functions

Write a function with one variadic parameter that finds the greatest number in a list of numbers.
*/

func getLargest(args ...int) int {
    if len(args) == 1 {
        return args[0]
    } else {
        largest := args[0]
        for _, value := range args {
            if value > largest {
                largest = value
            }
        }
        return largest
    }
}


func main() {
    a := getLargest(3,100,5,7,8,5,3,23,6,8,5,3)
    fmt.Println(a)
}
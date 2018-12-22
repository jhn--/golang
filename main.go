package main

import "fmt"

/*
functions

Variadic Functions
*/

func main() {
    xs := []int{1,2,3,4,5,6}
    total_1 := add(xs...)
    fmt.Println("total_1 from array/slice = ",total_1)

    total_2 := add(10, 9, 8)
    fmt.Println("total_2 from a bunch of ints = ",total_2)

    total_3 := add()
    fmt.Println("total_3 from zero arguments =",total_3)
}

func add(args ...int) (total int) {
    // accepts 0 or more arguments
    for _, value := range args {
        total += value
    }
    return 
}
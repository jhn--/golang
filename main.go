package main

import "fmt"

/*
Write a program that finds the smallest number in this list:

x := []int{
  48,96,86,68,
  57,82,63,70,
  37,34,83,27,
  19,97, 9,17,
}
*/

func main() {
    x := []int{
        48,96,86,68,
        57,82,63,70,
        37,34,83,27,
        19,97, 9,17,
    }

    smallest := x[0]

    for i := 1; i < len(x); i++ {
        // fmt.Println(x[i])
        if x[i] < smallest{
            // smallest := x[i]
            // fmt.Println(x[i], "is less than", smallest)
            // fmt.Println(x[i], "will be the smallest value.")
            smallest = x[i]
        }
    }
    fmt.Println("smallest value is", smallest)
}

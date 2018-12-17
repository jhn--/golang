package main

import "fmt"

/*
Write a program that prints out all the numbers evenly divisible by 3 between 1 and 100. (3, 6, 9, etc.)
*/

func main() {
    for count := 1; (count * 3) <= 100; count++{
        fmt.Println(count * 3)
    }
}

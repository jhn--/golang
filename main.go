package main

import "fmt"

/*
functions

The Fibonacci sequence is defined as: fib(0) = 0, fib(1) = 1, fib(n) = fib(n-1) + fib(n-2). Write a recursive function which can find fib(n).
*/

func fibonacci(x uint) uint {
    if (x == 0) {
        return 0
    } else if (x == 1) {
        return 1
    } else {
        return (fibonacci(x - 1) + fibonacci(x - 2))
    }
}

func main() {
    xs := []uint{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
    for _, value := range xs {
        fmt.Println("fib for", value, " is:", fibonacci(value))
    }
}
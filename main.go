package main

import "fmt"

/*
functions

Using makeEvenGenerator as an example, write a makeOddGenerator function that generates odd numbers.
*/

func makeOddGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        if i == 0 {
            i++
            ret = i
            return
        }
        i += 2
        ret = i
        return
    }
}

func main() {
    makeOdd := makeOddGenerator()
    fmt.Println(makeOdd())
    fmt.Println(makeOdd())
    fmt.Println(makeOdd())
    fmt.Println(makeOdd())
}
package main

import "fmt"

/*
map 1
*/

func main() {
    x := make(map[string]int)
    x["one"] = 1
    x["two"] = 2
    x["three"] = 3
    x["four"] = 4
    x["five"] = 5

    fmt.Println(x)
    fmt.Println(x["one"])
    fmt.Println("length of x is", len(x))

    // delete element with key "one" from map x.
    delete(x, "one")

    fmt.Println(x)
    fmt.Println("length of x is", len(x))

    // element can be found, return value and true
    name, ok := x["five"]
    fmt.Println(name, ok)

    // element cannot be found, return 0 and false
    noname, no_ok := x["cannot_be_found"]
    fmt.Println(noname, no_ok)

    // if-else statement 
    if name, ok := x["four"]; ok{
        fmt.Println(name, ok)
    }else{
        fmt.Println("can't find.")
    }

    y := make(map[int]int)
    y[1] = 1
    fmt.Println(y)
    fmt.Println(y[1])

    z := make(map[int]string)
    z[1] = "one"
    fmt.Println(z)
    fmt.Println(z[1])

    // shorter way of creating maps
    t := map[string]string{
        "a":"apple",
        "b":"bus",
        "c":"cat",
        "d":"dog",
    }
    fmt.Println(t)
}

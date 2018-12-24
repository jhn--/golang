package main

import "fmt"

/*
pointers
*/

func main() {
    one := "1: How do you get the memory address of a variable?"
    one_ans := "By using the & operator."
    two := "2: How do you assign a value to a pointer?"
    two_ans := "using the * operation, ie *xPtr = 0, means 'store the value 0 in the memory address xPtr refers to.'"   
    three := "3: How do you create a new pointer?"
    three_ans := "by using the new() function, new takes a type as an argument ie. new(int), and returns a pointer."

    qns_ans := []string {one, one_ans, two, two_ans, three, three_ans}

    for _, v := range qns_ans {
        fmt.Println(v)
    }
}

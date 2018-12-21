package main

import "fmt"

/*
functions
*/

func main() {
    xs := []float64{98,93,77,82,83}
    fmt.Println(average(xs))
}


func average(xs []float64) float64 {
    total := 0.0
    for _, value := range xs {
        total += value
    }
    /*
    for i := 0; i < len(xs); i++ {
        total += xs[i]
    }
    */
    return total/float64(len(xs))
}
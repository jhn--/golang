package main

import "fmt"

/*
nested map
*/

func main() {
    elements := map[string]map[string]string{
        "H":map[string]string{
            "name": "hydrogen",
            "state": "gas",
        },
        "He":map[string]string{
            "name": "helium",
            "state": "gas",
        },
        "Li": map[string]string{
          "name":"Lithium",
          "state":"solid",
        },
        "Be": map[string]string{
          "name":"Beryllium",
          "state":"solid",
        },
        "B":  map[string]string{
          "name":"Boron",
          "state":"solid",
        },
        "C":  map[string]string{
          "name":"Carbon",
          "state":"solid",
        },
        "N":  map[string]string{
          "name":"Nitrogen",
          "state":"gas",
        },
        "O":  map[string]string{
          "name":"Oxygen",
          "state":"gas",
        },
        "F":  map[string]string{
          "name":"Fluorine",
          "state":"gas",
        },
        "Ne":  map[string]string{
          "name":"Neon",
          "state":"gas",
        },
    }

    if el, ok := elements["Li"]; ok {
        fmt.Println(el["name"], el["state"])
    }

    fmt.Println(elements["Ne"])
    fmt.Println()
    fmt.Println(elements)
}

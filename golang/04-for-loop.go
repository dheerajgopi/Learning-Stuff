package main

import "fmt"

func main() {
        for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j, s = i+1, j+1, s + "a" {
                fmt.Println("value of i, j, a is : ", i, j, s)
        }

        fmt.Println()

        // range returns index...its like 'for each..in'
        a := [...]string{"a", "b", "c", "d"}
        for i := range a {
                fmt.Println("Array item ", i, " is ", a[i])
        }

        fmt.Println()

        // in maps, range returns keys
        capitals := map[string] string {"India": "New Delhi", "Japan": "Tokyo", "Italy": "Rome"}
        for key := range capitals {
                fmt.Println("Capital of ", key, " is ", capitals[key])
        }
}

package main

import "fmt"

func main() {
        a, b := 4, 5
        if a < b {
                fmt.Println("a is less than b")
        } else if a > b {
                fmt.Println("a is greater than b")
        } else {
                fmt.Println("a is equal to b")
        }
}

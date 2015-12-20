package main

import "fmt"

func main() {
        a := 5
        fmt.Println("a is : ", a)
        fmt.Println("value at address ", &a, " is ", *(&a))
        fmt.Println()

        var i int
        fmt.Println("default value of integer is : ", i)
        fmt.Println("value at address ", &i, " is ", *(&i))
        fmt.Println()

        var s string
        fmt.Println("default value of string is : ", s)
        fmt.Println("value at address ", &s, " is ", *(&s))
        fmt.Println()

        var f float64
        fmt.Println("default value of float64 is :", f)
        fmt.Println("value at address ", &f, " is ", *(&f))
        fmt.Println()

        var arInt [3]int
        fmt.Println("default value of int array is :", arInt)
        fmt.Println("value at address ", &arInt, " is ", *(&arInt))
        fmt.Println()

        var c complex64
        fmt.Println("default value of c is :", c)
        fmt.Println("value at address ", &c, " is ", *(&c))
        fmt.Println()

        ptr := &a
        fmt.Println("value at address ", &ptr, " is ", *(&ptr))
        fmt.Println()
}

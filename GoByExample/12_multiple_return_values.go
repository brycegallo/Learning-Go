package main

import "fmt"

// go has built-in support for multiple return valus, which allows for example for returning both result and error values from a function

// the (int, int) in this function signature shows that the function returns 2 ints
func vals() (int, int) {
    return 3, 7
}
func main() {

    // here we use 2 different return values from the same call with multiple assignment
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // if we only want a subset of the returned values, we use the blank identifier _
    _, c := vals()
    fmt.Println(c)
}

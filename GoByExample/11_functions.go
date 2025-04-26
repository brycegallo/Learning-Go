package main

import "fmt"

// this function takes two ints and returns their sum as an int
func plus(a int, b int) int {
    // go requires explicit returns, i.e. it won't automatically return the value of the last expression
    return a + b
}

// when we have multiple consecutive parameters of the same type, we may omit the type name for the like-typed parameters up to the final parameter that declares the type
func plusPlus(a, b, c int) int {
    return a + b + c
}

func main() {
    // we call functions as in most languages
    res := plus(1, 2)
    fmt.Println("1+2 =", res)

    res = plusPlus(1, 2, 3)
    fmt.Println("1+2+3 =", res)
}

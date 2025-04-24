package main

import (
    "fmt"
    "math"
)

const s string = "constant"

func main() {
    fmt.Println(s)

    // a const statement can appear anywhere a var statement can
    const n = 500000000

    // constant expressions perform arithmetic with arbitrary precision
    const d = 3e20 / n
    fmt.Println(d)

    // a numeric constant has no type until it's given one, such as by an explicit conversion here
    fmt.Println(int64(d))

    // a number can be given a type by using it in a context that requires on, such as a variable assignment or function call. e.g. here math.Sin expects input of type float64
    fmt.Println(math.Sin(n))
}

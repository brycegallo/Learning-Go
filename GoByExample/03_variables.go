package main

import "fmt"

func main() {
    // var declares 1 or more variables
    var a = "initial"
    fmt.Println(a)

    // we can declare multiple variables at once
    var b, c int = 1, 2
    fmt.Println(b, c)

    // Go will infer the types of initialized variables
    var d = true
    fmt.Println(d)

    // variables declared without a corresponding initialization are zero-valued
    var e int // this will be initialized to zero
    fmt.Println(e)

    // := is shorthand for declaring and initializing a variable at the same type
    f := "apple"
    fmt.Println(f)
}

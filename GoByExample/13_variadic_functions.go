package main

import "fmt"

// variadic functions can be called with any number of trailing arguments
// // fmt.Println is a common variadic function

// this function will take an arbitrary number of ints as arguments
func sum(nums ...int) {
    fmt.Print(nums, " ")
    total := 0

    // within this function, the type of nums is equivalent to []int
    // we can call len(nums), iterate over it with range, etc
    for _, num := range nums {
	total += num
    }
    fmt.Println(total)
}

func main() {
    // variadic functiosn can be called in the usual way with individual arguments
    sum(1, 2)
    sum(1, 2, 3)
    
    // if we have multiple args in a slice, we can apply them to a variadic function using func(slice...)
    nums := []int{1, 2, 3, 4}
    sum(nums...)

    // an important aspect of functions in Go is their ability to form closures, which we'll cover next
}

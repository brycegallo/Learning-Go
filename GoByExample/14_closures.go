package main

import "fmt"

// Go supports anonymous functions, which can form closures
// anonymous functions are useful when we want to define a function inline without having to name it

// this function intSeq() returns another function which we define anonymously in intSeq()'s body
// the returned function "closes over" the variable i to form a closure
func intSeq() func() int {
    i := 0
    return func() int {
	i++
	return i
    }
}

func main() {
    // we call intSeq(), assigning the result (a function) to nextInt
    // this function value captures its own i value, which will be updated each time we call nextInt
    nextInt := intSeq()

    // we can see the effect of closure by calling nextInt a few times
    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    // to confirm that the state is unique to that particular function, we create and test a new one
    newInts := intSeq()
    fmt.Println(newInts())
}

package main

// a goroutine is a lightweight thread of execution

import (
    "fmt"
    "time"
)

func f(from string) {
    for i := 0; i < 3; i++ {
	fmt.Println(from, ":", i)
    }
}

func main() {

    // suppose we have a function call f(s). here's how we 'd call that in the usual way, running it synchronously
    f("direct")

    // to invoke this function in a goroutine, use go f(s). this new goroutine will execute concurrently with the calling one
    go f("goroutine")

    // we can also start a goroutine for an anonymous function call
    go func(msg string) {
	fmt.Println(msg)
    }("going")

    // our two function calls are running asynchronously in separate goroutines now. we wait for them to finish (for a more robust appraoch, we use a WaitGroup, which will be covered in another lesson later)
    time.Sleep(time.Second)
    fmt.Println("done")

    // the next lesson will cover a complement of goroutines in concurrent Go programs: channels
}

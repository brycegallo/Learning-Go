package main

// we can use channels to synchronize execution across goroutines
// here we have an example of using a blocking receiver to wait for a goroutine to finish
// when waiting for multiple goroutines to finish, we may prefer to use a WaitGroup

import (
    "fmt"
    "time"
)

// this function will run in a goroutine
// the done channel will be used to notify another goroutine that this function's work is done
func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Print("done")
    // we send a value to notify that we're done
    done <- true
}

func main() {
    // we start a worker goroutine, giving it the channel to notify on
    done := make(chan bool, 1)
    go worker(done)
    // we blcok until we receive a notification from the worker on the channel
    <-done
    // if we removed the above line from this program, it would exit before the worker even started
}



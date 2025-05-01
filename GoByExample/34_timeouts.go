package main

// timeouts are important for programs that connect to external resources or that otherwise need to bound execution time
// implementing times in Go is easy thanks to channels and select

import (
    "fmt"
    "time"
)

func main() {

    // for our example, suppose we're executing an internal call that returns its result on a channel c1 after 2 seconds. note that the channel is buffered, so the send in the goroutine is nonblocking. this is a common pattern to prevent goroutine leaks in case the channel is never read
    c1 := make(chan string, 1)
    go func() {
	time.Sleep(2 * time.Second)
	c1 <- "result 1"
    }()

    // here select is implementing a timeout. res := <-c1 awaits the result and <-time.After awaits a value to be sent after the timeout of 1 second
    // since select proceeds with the first receive that's ready, we'll take the timeout case if the operations takes more than the allowed 1 second
    select {
    case res := <-c1:
	fmt.Println(res)
    case <-time.After(1 * time.Second):
	fmt.Println("timeout 1")
    }

    // if we allow a longer timeout of 3 seconds, the receive from c2 will succeed and print the result
    c2 := make(chan string, 1)
    go func() {
	time.Sleep(2 * time.Second)
	c2 <- "result 2"
    }()
    select {
    case res := <-c2:
	fmt.Println(res)
    case <-time.After(3 * time.Second):
	fmt.Println("timeout 2")
    }
    // running this program shows the first operation timing out and the second succeeding
}

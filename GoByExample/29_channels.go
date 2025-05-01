package main

// channels are the pipes that connect concurrent goroutines. we can send values into those channels from one goroutine and receive those values into another goroutine

import "fmt"

func main() {

    // we create a new channel with make(chan val-type). channels are typed by the values they convey
    messages := make(chan string)

    // we send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above, from a new goroutine
    go func() { messages <- "ping" }()

    // the <- channel syntax receives a value from the channel. Here we'll receive the "ping" message we sent above and print it out
    msg := <-messages
    fmt.Println(msg)

    // when we run the program, the "ping" message is successfully passed from one goroutine to another via our channel

    // by default, sends and receives block until both the sender and receiver are ready. this property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization
}

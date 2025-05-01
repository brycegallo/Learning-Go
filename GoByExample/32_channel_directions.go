package main

// when using channels as function parameters, we can specify if a channel is meant to only send or receive values
// this specificity increases the type-safety of the program

import "fmt"

// this ping function only accepts a channel for sending values
// it would be a compile-time error to try and receive on this channle
func ping(pings chan<- string, msg string) {
    pings <- msg
}

// this pong function accepts one channel for receives (pings) and a second for sends (pongs)
func pong(pings <- chan string, pongs chan<- string) {
    msg := <-pings
    pongs <- msg
}

func main() {
    pings := make(chan string, 1)
    pongs := make(chan string, 1)
    ping(pings, "passed message")
    pong(pings, pongs)
    fmt.Println(<-pongs)
}

 package main

 // Go's select lets us wait on multiple channel operations
 // combining goroutines and channels with select is a powerful feature of Go

 import (
     "fmt"
     "time"
 )

 func main() {
     // for our example we'll select across two channels
     c1 := make(chan string)
     c2 := make(chan string)

     // each channel will receive a value after some amount of time, to simulate e.g. blocking RPC operations executing in concurrent goroutines
     go func() {
	 time.Sleep(1 * time.Second)
	 c1 <- "one"
     }()
     go func() {
	 time.Sleep(2 * time.Second)
	 c2 <- "two"
     }()

     // we'll use select to await both of these values simultaneously, printing each as it arrives
     for i := 0; i < 2; i++ {
	 select {
	 case msg1 := <-c1:
	     fmt.Println("received", msg1)
	 case msg2 := <-c2:
	     fmt.Println("received", msg2)
	 }
     }
     // we'll receive the values "one" and then "two" as expected

     // note that the total execution time will be around only 2 seconds since both the 1 and 2 second Sleeps execute concurrently
 }	 

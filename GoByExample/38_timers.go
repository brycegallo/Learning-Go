package main

// we often want to execute Go code at some point in the future, or repeatedly at some interval
// Go's built-in time and ticker features make both of these tasks easy
// first, we'll look at timers here

import (
    "fmt"
    "time"
)

func main() {
    // timers represent a single event in the future. we tell the timer how long we want to wait, and it provides a channel that will be notified at that time, this timer below will wait 2 seconds

    timer1 := time.NewTimer(2 * time.Second)

    // the <-timer1.C blocks on the timer's channel c until it sends a value indicating that the timer fired
    <-timer1.C
    fmt.Println("Timer 1 fired")

    // if we just wanted to wait, we could have used time.Sleep
    // one reason a timer may be useful is that we can cancel a time before it fires
    // here's an example below
    timer2 := time.NewTimer(time.Second)
    go func() {
	<-timer2.C
	fmt.Println("Timer 2 fired")
    }()
    stop2 := timer2.Stop()
    if stop2 {
	fmt.Println("Timer 2 stopped")
    }

    // giving timer2 enough time to fire, if it was ever going to, to show it is in fact stopped 
    time.Sleep(2 * time.Second)

    // the first timer will fire around 2 seconds after we start the program, but the second timer should be stopped before it has a chance to fire
}

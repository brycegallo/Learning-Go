package main

// by default, channels are unbuffered, meaning that they will only accepts sends (chan <-) if there is a corresponding receive (<- chan) ready to receive the sent value. buffered channels accept a limited number of values without a corresponding receiver for those values

import "fmt"

func main() {

    // here we make a channel of strings buffering up to 2 values
    messages := make(chan string, 2)

    // because this channel is buffered, we can send these values into the channel without a corresponding concurrent receive
    messages <- "buffered"
    messages <- "channel"

    // later we can receive these two values as usual
    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

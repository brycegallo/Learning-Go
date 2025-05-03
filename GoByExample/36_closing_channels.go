package main

// closing a channel indicates that no more values will be sent on it
// this can be useful to communicate completion to the channel's receivers

import "fmt"

// in this example we'll use a jobs channel to communicate work to be done from the main() goroutine to a worker goroutine
// when we have no more jobs for the worker we'll close the jobs channel
func main() {
    jobs := make(chan int, 5)
    done := make(chan bool)

    // here's the worker goroutine. it repeatedly receives from jobs with j, more := jobs
    // in thi special 2-value form of receive, the more value will be false if jobs has been closed, and all values in the channel have already been received
    // we use this to notify on done when we've worked all our jobs
    go func() {
	for {
	    j, more := <-jobs
	    if more {
		fmt.Println("received job", j)
	    } else {
		fmt.Println("received all jobs")
		done <- true
		return
	    }
	}
    }()

    // this sends 3 jobs to the worker over the jobs channel, then closes it
    for j := 1; j <= 3; j++ {
	jobs <- j
	fmt.Println("sent job", j)
    }
    close(jobs)
    fmt.Println("sent all jobs")

    <-done
    // we await the worker using the synchronization approach we saw in earlier lessons
    _, ok := <-jobs
    fmt.Println("received more jobs:", ok)
    // reading from a close channel succeeds immediately, returning the zero value of the underlying type
    // the optional second return value is true if the value received was delivered by a successful send operation to the channel, or false if it was a zero value generated because the channel is closed and empty
}
// the idea of closed channels leads naturally to our next example: range over channels

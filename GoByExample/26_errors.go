package main

// in Go it's idiomatic to communicate errors via an explicit, separate return value
// this contrasts with the exceptions used in languages like Java and the overloaded single result / error value sometimes used in C
// Go's approach makes it easy to see which functions return errors and to handle them using the same language constructs employed for other, non-error tasks

import (
    "errors"
    "fmt"
)

// by convention, errors are the last return value and have type error, a built-in interface
func f(arg int) (int, error) {
    if arg == 42 {
	// errors.New constructs a basic error value with the given error message
	return -1, errors.New("can't work with 42")
    }
    // a nil value in the rror position indicats that there was no error
    return arg + 3, nil
}

// a sentinel error is a predeclared variable that is used to signif a specific error condition
var ErrOutOfTea = fmt.Errorf("no more tea available")
var ErrPower = fmt.Errorf("can't boil water")

func makeTea(arg int) error {
    if arg == 2 {
	return ErrOutOfTea
    } else if arg == 4 {
	// we can wrap errors with higher-level errors to add context. The simplest way to do this is with the %w verb in fmt.Errorf. Wrapped errors create a logical chain (A wraps B, which wraps C, etc) that can be queried with functions like error.Is and errors.As
	return fmt.Errorf("making tea: %w", ErrPower)
    }
    return nil
}

func main() {
    for _, i := range []int{7, 42} {
	// it's common to use an inline error to check in the if line
	if r, e := f(i); e != nil {
	    fmt.Println("f failed:", e)
	} else {
	    fmt.Println("f worked:", r)
	}
    }

    for i := range 5 {
	if err := makeTea(i); err != nil {
	    // errors.Is checks that a given error (or any error in its chain) matches a specific error value. This is especially useful with wrapped or nest errors, allowing you to identify specific error types or sentinel errors in a chain of errors
	    if errors.Is(err, ErrOutOfTea) {
		fmt.Println("We should buy new tea!")
	    } else if errors.Is(err, ErrPower) {
		fmt.Println("Now it is dark.")
	    } else {
		fmt.Println("unknown error: %s\n", err)
	    }
	    continue
	}
	fmt.Println("Tea is ready")
    }
}


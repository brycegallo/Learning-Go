package main

import "fmt"

func main() {

    // most basic type with single condition
    i := 1
    for i <= 3 {
	fmt.Println(i)
	i = i + 1
    }

    // a C-like initial/condition/after for loop
    for j := 0; j <3; j++ {
	fmt.Println(j)
    }

    // range over an integer is another option
    for i := range 3 {
	fmt.Println("range", i)
    }

    // conditionless-for will loop until broken out of, or return from enclosing function
    for {
	fmt.Println("loop")
	break
    }

    // continue will skip to the next iteration of the loop
    for n:= range 6 {
	if n%2 == 0 {
	    continue
	}
	fmt.Println(n)
    }
}

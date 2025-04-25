package main

import "fmt"

func main() {

    // conditions in go don't need parentheses but do need brackets
    if 7%2 == 0 {
	fmt.Println("7 is even")
    } else {
	fmt.Println("7 is odd")
    }
    
    if 8%4 == 0 {
	fmt.Println("8 is divisible by 4")
    }

    if 8%2 == 0 || 7%2 == 0 {
	fmt.Println("either 8 or 7 are even")
    }

    // a statement can precede conditionals
    // any variables declared in this statement are available in the current and all subsequent branches
    if num := 9; num < 0 {
	fmt.Println(num, "is negative")
    } else if num < 10 {
	fmt.Println(num, "has 1 digit")
    } else {
	fmt.Println(num, "has multiple digits")
    }
    // there is no ternary if in Go, so a full statement is needed for even basic conditions
}

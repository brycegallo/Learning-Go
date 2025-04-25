package main

import "fmt"

func main() {

    // here we create an array a that will hold exactly 5 ints
    // the type of elements and length are both part of the earray's type
    // array's are zero-valued by default, which for ints means 0s
    var a [5]int
    fmt.Println("emp:", a)

    // we can set a value at an index using the "array[index] = value" syntax, and get a value with array[index]
    a[4] = 100
    fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    // len is a builtin which returns the length of an array
    fmt.Println("len:", len(a))

    // this syntax declares and initializes an array in one line
    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // using ... will have the compiler count the number of elements for us
    b = [...]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    // if you specify the index with :, the elements in between will be zeroed
    b = [...]int{100, 3: 400, 500}
    fmt.Println("idx:", b)

    // array types are one-dimensional, but we can compose types to build multi-dimensional data structures
    var twoD [2][3]int
    for i := 0; i < 2; i++ {
	for j := 0; j < 3; j++ {
	    twoD[i][j] = i + j
	}
    }
    fmt.Println("2d: ", twoD)

    // we can create and initialize multi-dimensional arrays at once too
    twoD = [2][3]int{
	{1, 2, 3},
	{1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
    // note that arrays appear in the form [v1 v2 v3 ...] when printed with fmt.Println
}



package main

import (
    "fmt"
    "slices"
)

func main() {
    // unlike arrays, slices are typed only by the elements they contain (not the number of elements)
    // an uninitialized slice equals to nil and has a length of 0
    var s []string
    fmt.Println("uninit:", s, s == nil, len(s) == 0)

    // to create a slice with non-zero length, we use the builtin make
    // here we make a slice of strings of length 3 (initially zero-valued)
    // by default a new slice's capacity equals its length
    // if we know the slice is going to grow ahead of time, it's possible to pass a capacity explicitly to make
    s = make([]string, 3)
    fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

    // we can set and get like with arrays
    s[0] = "a"
    s[1] = "b"
    s[2] = "c"
    fmt.Println("set:", s)
    fmt.Println("get:", s[2])

    // len returns the length of the slice
    fmt.Println("len:", len(s))

    // slices support several operations that make them richer than arrays
    // the append builtin returns a slice containing one or more new values
    // we need to accept a return value from append as we may get a new slice value
    s = append(s, "d")
    s = append(s, "e", "f")
    fmt.Println("apd:", s)

    // slices can also be copy'd
    // here we create an empty slice c of the same length as s and copy into c from s
    c := make([]string, len(s))
    copy(c, s)
    fmt.Println("cpy:", c)

    // slices support a slice operator with the syntax slice[low:high] 
    // here we get a slice of the elements s[2], s[3], and s[4]
    l := s[2:5]
    fmt.Println("sl1:", l)

    // this slices up to (but excluding) s[5]
    l = s[:5]
    fmt.Println("s2:", l)

    // this slices up to (and including) s[2]
    l = s[2:]
    fmt.Println("s3:", l)

    // we can declare and initialize a variable for slice in a single line
    t := []string{"g", "h", "i"}
    fmt.Println("dlc:", t)

    // the slice package contains a number of useful utility functions for slices
    t2 := []string{"g", "h", "i"}
    if slices.Equal(t, t2) {
	fmt.Println("t == t2")
    }

    // slices can be composed into multi-dimensional data structures. the length of the inner slices can vary,unlike with multi-dimensional arrays
    twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
	innerLen := i + 1
	twoD[i] = make([]int, innerLen)
	for j := 0; j < innerLen; j++ {
	    twoD[i][j] = i + j
	}
    }
    fmt.Println("2d: ", twoD)
    // while slices are different types than arrays, they are rendered similarly by fmt.Println
}

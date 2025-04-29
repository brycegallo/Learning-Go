package main

import "fmt"

// Go's structs are typed collections of fields
// they're useful fpr grpi[omg data together to form records

// this person struct has name and age fields
type person struct {
    name string
    age int
}

// newPerson constructs a new person struct with the given name
func newPerson(name string) *person {
    p := person{name: name}
    p.age = 42
    return &p
    // Go is a garbage collected language, we can safely return a pointer to a local variable
    // it will only be cleaned up by the garbage colelctor when there are no active references to it
}

func main() {
    // this syntax creates a new struct
    fmt.Println(person{"Bob", 20})

    // we can name the fields when initializing a struct
    fmt.Println(person{name: "Alice", age: 30})

    // ommitted fields will be zero-valued
    fmt.Println(person{name: "Fred"})

    // an & prefix yields a pointer to the struct
    fmt.Println(&person{name: "Ann", age: 40})

    // it's idiomatic to encapsulate new struct creation in constructor functions
    fmt.Println(newPerson("Jon"))

    // we access struct fields with a dot
    s := person{name: "Sean", age: 50}
    fmt.Println(s.name)

    // we can also use dots with struct pointers - the pointers are automatically dereferenced
    sp := &s
    fmt.Println(sp.age)

    // structs are mutable
    sp.age = 51
    fmt.Println(sp.age)
    
    // if a struct type is only used for a single value, we don't have to give it a name
    // the value can have an anonymous struct type, this technique is commonly used for table-driven tests
    dog := struct {
	name string
	isGood bool
    }{
	"Rex",
	true,
    }
    fmt.Println(dog)
}

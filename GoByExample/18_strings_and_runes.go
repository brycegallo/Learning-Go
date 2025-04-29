package main

import (
    "fmt"
    "unicode/utf8"
)

// a Go string is a read-only slice of bytes. the language and the standard library treats string specially - as containers of text encoded in UTF-8. in other languages, strings are made of "characters"
// in Go, the concept of a character is called a rune - it's an integer that represents a unicode point

func main() {

    // s is a string assigned a literal value representing the world "hello" in Thai
    // Go string literals are UTF-8 encoded text
    const s = "สวัสดี"

    // since strings are equivalent to []byte, this will produce the length of the raw bytes stored within
    fmt.Println("Len:", len(s))

    // indexing into a string produces the raw byte values at each index
    // this loop generates the hex values of all the bytes that constitute the code points in s
    for i := 0; i < len(s); i++ {
	fmt.Printf("%x ", s[i])
    }
    fmt.Println()

    // to count how many runes are in a string, we can use the utf8 package
    // note that the run-time of RuneCountInString depends on the size of the string, because it has to decode each UTF-8 rune sequentially
    // some Thai characters are represented by UTF-8 code points that can span multiple bytes, so the result of this count may be surprising
    fmt.Println("Rune count:", utf8.RuneCountInString(s))

    // a range loop handles strings specially and decodes each rune along with its offset in the string
    for idx, runeValue := range s {
	fmt.Printf("%#U starts at %d\n", runeValue, idx)
    }

    fmt.Println("\nUsing DecodeRuneInString")
    for i, w := 0, 0; i < len(s); i += w {
	runeValue, width := utf8.DecodeRuneInString(s[i:])
	fmt.Printf("%#U starts at %d\n", runeValue, i)
	w = width

	// this demonstrates passing a rune value to a function
	examineRune(runeValue)
    }
}

func examineRune(r rune) {
    // values enclosed in single quotes are rune literals
    // we can compare ar une value to a rune literal directly
    if r == 't' {
	fmt.Println("found tee")
    } else if r == 'ส' {
	fmt.Println("found so sua")
    }
}




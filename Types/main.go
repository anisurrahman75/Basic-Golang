package main

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

/*
The following are the basic types in Go
- bool
- numeric types
  - int8, int16, int32, int64, int (signed integers)
  - uint8, uint16, uint32, uint32, uint64, uint, uintptr (unsigned integers)
  - float32, float64
  - complex64, complex128
  - byte (alias for uint8)
  - rune (alias for int32)

-string
*/
func main() {
	/*
		-------------------------------------Bool-------------------------------------
			- and: &&
			- or: ||
			- not: !
	*/
	a := true
	b := false
	fmt.Println("a: ", a, "b: ", b)
	c := a && b
	fmt.Println("C: ", c)
	d := a || b
	fmt.Println("D: ", d)
	e := !a
	fmt.Println("E: ", e)
	/*
		-------------------------------------Complex Number-------------------------------------
		- There are two parts of a complex number. The real and imaginary part.
		- we can use complex constructor for initialization. Ex: c1:= complex(2, 3)
	*/
	c1 := complex(2, 3)
	c2 := 4 + 5i             // complex initializer syntax a + ib
	c3 := c1 + c2            // addition just like other variables
	fmt.Println("Add: ", c3) // prints "Add: (6+8i)"
	re := real(c3)           // get real part
	im := imag(c3)           // get imaginary part
	fmt.Println(re, im)      // prints 6 8

	/*
		-------------------------------------STRING-------------------------------------

		- A string is a slice of bytes in Go. Strings in Go are Unicode compliant and are UTF-8 Encoded.
		- Strings can be created by enclosing a set of characters inside double quotes " ".
	*/
	firstName := "Anisur"
	LastName := "Rahman"
	fmt.Println("firstName: ", firstName, "\nLastName: ", LastName)
	name := firstName + LastName
	// Since a string is a slice of bytes, it's possible to access each byte of a string.
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c", name[i]) // character
	}
	fmt.Println()

	// gives wrong output as 'ñ' has unicode value 'U+00F1'
	name = "Señor"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}
	fmt.Printf("\n\n")
	// A rune is a builtin type in Go and it's the alias of int32. Rune represents a Unicode code point in Go.
	runes := []rune(name)
	name = firstName + LastName
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
		//fmt.Print(name[i], " ") // Byte-Code
	}
	fmt.Printf("\n\n")
	// string length
	/*
		The RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string.
		This method takes a string as an argument and returns the number of runes in it.
		As we discussed earlier, len(s) is used to find the number of bytes in the string and it doesn't return the string length.
	*/

	fmt.Println("Runes in name:", utf8.RuneCountInString(name))
	fmt.Println("Bytes in name:", len(name))
	name = firstName + LastName
	// strings are immutable
	/*
		error: cannot assign to name[0]
		name[0] = 'p'
		fmt.Println(name)
		To work with it we need to work with runes
	*/
	runes = []rune(name)
	runes[0] = 'P'
	name = string(runes) // type conversion rune to string
	fmt.Println(name)
	/* ........................................................Type Conversion
	- Go is very strict about explicit typing. There is no automatic type promotion or conversion.
	*/
	p := 10
	q := float32(p)
	fmt.Printf("%T\n", q) // Printing type of 'q'

	// Type uintptr in Golang is an integer type that is large enough to contain the bit pattern of any pointer. In other words, uintptr is an integer representation of a memory address.
	// Below is how we declare a variable of type uintptr:
	var u uintptr = 0xc82000c290
	fmt.Println(u)
	fmt.Printf("%T\n", u)
	// https://stackoverflow.com/questions/59042646/whats-the-difference-between-uint-and-uintptr-in-golang

	//......................................................... string to int converting, make sure that you import "strconv"
	s := "123"
	fmt.Printf("%T, %v\n", s, s)
	num, err := strconv.Atoi(s)
	//num, err := strconv.ParseInt(s, 10, 0)
	if err != nil {
		fmt.Println("some error occurred while conversion")
	} else {
		fmt.Println("no error, successfully converted!")
	}
	fmt.Printf("%T, %v\n", num, num)

	// int to string conversion, import "strconv" & look that Itoa returns only one value
	h := 123
	fmt.Printf("%T, %v\n", h, h)
	t := strconv.Itoa(h)
	fmt.Printf("%T, %v\n", t, t)
	t = strconv.FormatInt(int64(h), 10)
	fmt.Printf("%T, %v\n", t, t)
}
